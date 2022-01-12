import * as cdk from "aws-cdk-lib";

export enum ReplicationDestinationStorageClass {
  DEEP_ARCHIVE = "DEEP_ARCHIVE",
  GLACIER = "GLACIER",
  INTELLIGENT_TIERING = "INTELLIGENT_TIERING",
  ONEZONE_IA = "ONEZONE_IA",
  OUTPOSTS = "OUTPOSTS",
  REDUCED_REDUNDANCY = "REDUCED_REDUNDANCY",
  STANDARD = "STANDARD",
  STANDARD_IA = "STANDARD_IA",
}

export enum ReplicationRuleStatus {
  DISABLED = "Disabled",
  ENABLED = "Enabled",
}

export interface ReplicationRulePropertyNoDestination {
  readonly status?: ReplicationRuleStatus;
  readonly deleteMarkerReplication?: cdk.IResolvable | cdk.aws_s3.CfnBucket.DeleteMarkerReplicationProperty;
  readonly filter?: cdk.IResolvable | cdk.aws_s3.CfnBucket.ReplicationRuleFilterProperty;
  readonly id?: string;
  readonly prefix?: string;
  readonly priority?: number;
  readonly sourceSelectionCriteria?: cdk.IResolvable | cdk.aws_s3.CfnBucket.SourceSelectionCriteriaProperty;
}
export interface ReplicationDestinationPropertyNoBucket {
  readonly accessControlTranslation?: cdk.IResolvable | cdk.aws_s3.CfnBucket.AccessControlTranslationProperty;
  readonly account?: string;
  readonly encryptionConfiguration?: cdk.IResolvable | cdk.aws_s3.CfnBucket.EncryptionConfigurationProperty;
  readonly metrics?: cdk.IResolvable | cdk.aws_s3.CfnBucket.MetricsProperty;
  readonly replicationTime?: cdk.IResolvable | cdk.aws_s3.CfnBucket.ReplicationTimeProperty;
  readonly storageClass?: ReplicationDestinationStorageClass;
}

export interface BucketReplicationProps {
  readonly sourceBucket: cdk.aws_s3.Bucket;
  readonly destinationBucket: cdk.aws_s3.Bucket;
  readonly replicationRuleProperties?: ReplicationRulePropertyNoDestination;
  readonly replicationDestinationProperties?: ReplicationDestinationPropertyNoBucket;
}

export class BucketReplication extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props: BucketReplicationProps) {
    super(scope, id);

    const sourceAccount = cdk.Stack.of(props.sourceBucket).account;
    const destinationAccount = cdk.Stack.of(props.destinationBucket).account;

    const cfnSourceBucket = props.sourceBucket.node.defaultChild as cdk.aws_s3.CfnBucket;

    const replicationRole = new cdk.aws_iam.Role(this, "ReplicationRole", {
      assumedBy: new cdk.aws_iam.ServicePrincipal("s3.amazonaws.com"),
    });

    replicationRole.addToPolicy(
      new cdk.aws_iam.PolicyStatement({
        effect: cdk.aws_iam.Effect.ALLOW,
        resources: [`${props.destinationBucket.bucketArn}/*`],
        actions: ["s3:ReplicateObject", "s3:ReplicateDelete", "s3:ReplicateTags"],
      }),
    );
    replicationRole.addToPolicy(
      new cdk.aws_iam.PolicyStatement({
        effect: cdk.aws_iam.Effect.ALLOW,
        resources: [props.sourceBucket.bucketArn],
        actions: ["s3:GetReplicationConfiguration", "s3:ListBucket"],
      }),
    );
    replicationRole.addToPolicy(
      new cdk.aws_iam.PolicyStatement({
        effect: cdk.aws_iam.Effect.ALLOW,
        resources: [`${props.sourceBucket.bucketArn}/*`],
        actions: ["s3:GetObjectVersion", "s3:GetObjectVersionAcl", "s3:GetObjectVersionTagging"],
      }),
    );
    if (sourceAccount !== destinationAccount) {
      props.destinationBucket.addToResourcePolicy(
        new cdk.aws_iam.PolicyStatement({
          effect: cdk.aws_iam.Effect.ALLOW,
          principals: [new cdk.aws_iam.ArnPrincipal(`arn:aws:iam::${sourceAccount}:root`)],
          resources: [`${props.destinationBucket.bucketArn}/*`],
          actions: ["s3:ReplicateObject", "s3:ReplicateDelete", "s3:ReplicateTags"],
        }),
      );
      props.destinationBucket.addToResourcePolicy(
        new cdk.aws_iam.PolicyStatement({
          effect: cdk.aws_iam.Effect.ALLOW,
          principals: [new cdk.aws_iam.ArnPrincipal(`arn:aws:iam::${sourceAccount}:root`)],
          resources: [`${props.destinationBucket.bucketArn}`],
          actions: ["s3:List*", "s3:GetBucketVersioning", "s3:PutBucketVersioning"],
        }),
      );
    }
    cfnSourceBucket.replicationConfiguration = {
      role: replicationRole.roleArn,
      rules: [
        {
          destination: {
            storageClass: ReplicationDestinationStorageClass.STANDARD,
            ...props.replicationDestinationProperties,
            bucket: props.destinationBucket.bucketArn,
          },
          status: ReplicationRuleStatus.ENABLED,
          ...props.replicationRuleProperties,
        },
      ],
    };
  }
}

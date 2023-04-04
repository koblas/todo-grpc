import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { TerraformAsset } from "cdktf";

export enum Runtime {
  GO1_X = "go1.x",
  PYTHON3_8 = "python3.8",
  AL2 = "provided.al2",
}

export enum Architecture {
  ARM = "arm64",
  X86 = "x86_64",
}

export interface StackLambdaConfig {
  /**
   * Map of allowed triggers to create Lambda permissions
   */
  allowedTriggers?: (Pick<
    aws.lambdaPermission.LambdaPermissionConfig,
    "statementId" | "action" | "principal" | "sourceAccount" | "sourceArn" | "eventSourceToken"
  > & { service?: string })[];

  /**
   * Instruction set architecture for your lambda function
   */
  architecture: Architecture;

  /**
   * The asset for this function
   *   (local_existing_package)
   */
  asset?: TerraformAsset;

  /**
   * Controls whether async event policy should be added to IAM role for Lambda Function
   */
  attachAsyncEventPolicy?: boolean;

  /**
   * Controls whether CloudWatch Logs policy should be added to IAM role for Lambda Function
   */
  attachCloudwatchLogsPolicy?: boolean;

  /**
   * Controls whether SNS/SQS dead letter notification policy should be added to
   * IAM role for Lambda Function
   */
  attachDeadLetterPolicy?: boolean;

  /**
   * Controls whether VPC/network policy should be added to IAM role for Lambda Function
   */
  attachNetworkPolicy?: boolean;

  /**
   * Controls whether policy should be added to IAM role for Lambda Function
   */
  attachPolicy?: boolean;

  /**
   * Controls whether policy_json should be added to IAM role for Lambda Function
   */
  attachPolicyJson?: boolean;

  /**
   * Controls whether policy_jsons should be added to IAM role for Lambda Function
   */
  attachPolicies?: aws.iamPolicy.IamPolicy[];

  /**
   * Controls whether policy_statements should be added to IAM role for Lambda Function
   */
  attachPolicyStatements?: boolean;

  /**
   * Controls whether X-Ray tracing policy should be added to IAM role for Lambda Function
   */
  attachTracingPolicy?: boolean;

  /**
   * Map of dynamic policy statements for assuming Lambda Function role (trust relationship)
   */
  assumeRolePolicyStatements?: aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocumentStatement[];

  /**
   * The type of authentication that the Lambda Function URL uses. Set to 'AWS_IAM' to
   * restrict access to authenticated IAM users only. Set to 'NONE' to bypass IAM
   * authentication and create a public endpoint.
   */
  authorizationType?: string;

  /**
   * The ARN of the KMS Key to use when encrypting log data.
   */
  cloudwatchLogsKmsKeyId?: string;

  /**
   * Specifies the number of days you want to retain log events in the specified log group.
   * Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.
   */
  cloudwatchLogsRetentionInDays?: number;

  /**
   * A map of tags to assign to the resource.
   */
  cloudwatchLogsTags?: Record<string, string>;

  /**
   * Amazon Resource Name (ARN) for a Code Signing Configuration
   */
  codeSigningConfigArn?: string;

  /**
   * A list of Architectures Lambda layer is compatible with. Currently x86_64 and arm64 can be specified.
   */
  compatibleArchitectures?: Architecture[];

  /**
   * A list of Runtimes this layer is compatible with. Up to 5 runtimes can be specified.
   */
  compatibleRuntimes?: Runtime[];

  /**
   * CORS settings to be used by the Lambda Function URL
   */
  cors?: aws.lambdaFunctionUrl.LambdaFunctionUrlCors;

  /**
   * Controls whether resources should be created
   */
  create?: boolean;

  /**
   * Controls whether async event configuration for Lambda Function/Alias should be created
   */
  createAsyncEventConfig?: boolean;

  /**
   * Whether to allow triggers on current version of Lambda Function (this will revoke
   * permissions from previous version because Terraform manages only current resources)
   */
  createCurrentVersionAllowedTriggers?: boolean;

  /**
   * Whether to allow async event configuration on current version of Lambda
   * Function (this will revoke permissions from previous version because
   * Terraform manages only current resources)
   */
  createCurrentVersionAsyncEventConfig?: boolean;

  /**
   * Controls whether Lambda Function resource should be created
   */
  createFunction?: boolean;

  /**
   * Controls whether the Lambda Function URL resource should be created
   */
  createFunctionUrl?: boolean;

  /**
   * Controls whether Lambda Layer resource should be created
   */
  createLayer?: boolean;

  /**
   * Controls whether Lambda package should be created
   */
  createPackage?: boolean;

  /**
   * Controls whether IAM role for Lambda Function should be created
   */
  createRole?: boolean;

  /**
   * Whether to allow triggers on unqualified alias pointing to $LATEST version
   */
  createUnqualifiedAliasAsyncEventConfig?: boolean;

  /**
   * Whether to allow async event configuration on unqualified alias pointing to $LATEST version
   */
  createUnqualifiedAliasLambdaFunctionUrl?: boolean;

  /**
   * Whether to allow triggers on unqualified alias pointing to $LATEST version
   */
  createUnqualifiedAliasAllowedTriggers?: boolean;

  /**
   * Whether to use an existing CloudWatch log group or create new
   */
  useExistingCloudwatchLogGroup?: boolean;

  /**
   * The ARN of an SNS topic or SQS queue to notify when an invocation fails.
   */
  deadLetterTargetArn?: string;

  /**
   * A map that defines environment variables for the Lambda Function.
   */
  environment?: aws.lambdaFunction.LambdaFunctionConfig["environment"];

  /**
   * Amount of ephemeral storage (/tmp) in MB your Lambda Function can use at runtime.
   * Valid value between 512 MB to 10,240 MB (10 GB).
   */
  ephemeralStorageSize?: number;

  /**
   * A unique name for your Lambda Function
   */
  functionName: string;

  /**
   * IAM policy name. It override the default value, which is the same as role_name
   */
  roleName?: string;

  /**
   * Description of IAM role to use for Lambda Function
   */
  roleDescription?: string;

  /**
   * The ARN of the policy that is used to set the permissions boundary for the IAM role used by Lambda Function
   */
  rolePermissionBoundary?: string;

  /**
   * Path of IAM role to use for Lambda Function
   */
  rolePath?: string;

  /**
   * Path of policies to that should be added to IAM role for Lambda Function
   */
  policyPath?: string;

  /**
   * Specifies to force detaching any policies the IAM role has before destroying it.
   */
  roleForceDetachPolicies?: boolean;

  /**
   * A map of tags to assign to IAM role
   */
  roleTags?: Record<string, string>;

  /**
   * Description of your Lambda Function (or Layer)
   */
  description?: string;

  /**
   * Lambda Function entrypoint in your code
   */
  handler?: string;

  /**
   * Whether to ignore changes to the function's source code hash.
   * Set to true if you manage infrastructure and code deployments separately.
   */
  ignoreSourceCodeHash?: boolean;

  /**
   * The CMD for the docker image
   */
  imageConfigCommand?: string[];

  /**
   * The ENTRYPOINT for the docker image
   */
  imageConfigEntryPoint?: string[];

  /**
   * The working directory for the docker image
   */
  imageConfigWorkingDirecdtory?: string;

  /**
   * The ECR image URI containing the function's deployment package.
   */
  imageUri?: string;

  /**
   * The ARN of KMS key to use by your Lambda Function
   */
  kmsKeyArn?: string;

  /**
   * Set this to true if using Lambda@Edge, to enable publishing,
   * limit the timeout, and allow edgelambda.amazonaws.com to invoke the function
   */
  lambdaAtEdge?: boolean;

  /**
   * IAM role ARN attached to the Lambda Function. This governs both who / what
   * can invoke your Lambda Function, as well as what resources our Lambda Function
   * has access to. See Lambda Permission Model for more details.
   */
  lambdaRoleArn?: string;

  /**
   * Name of Lambda Layer to create
   */
  layerName?: string;

  /**
   * Whether to retain the old version of a previously deployed Lambda Layer.
   */
  layerSkipDestroy?: boolean;

  /**
   * List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function.
   */
  layers?: string[];

  /**
   * License info for your Lambda Layer. Eg, MIT or full url of a license.
   */
  licenseInfo?: string;

  /**
   * Amount of memory in MB your Lambda Function can use at runtime.
   * Valid value between 128 MB to 10,240 MB (10 GB), in 64 MB increments.
   */
  memorySize?: number;

  /**
   * The Lambda deployment package type. Valid options: Zip or Image
   */
  packageType?: "Zip" | "Image";

  /**
   * Whether to publish creation/change as new Lambda Function Version.
   */
  publish?: boolean;

  /**
   * Amount of capacity to allocate. Set to 1 or greater to enable, or set to 0 to
   * disable provisioned concurrency.
   */
  provisionedConcurrentExecutions?: number;

  /**
   * Lambda Function runtime
   */
  runtime: Runtime;

  /**
   * The amount of reserved concurrent executions for this Lambda Function.
   * A value of 0 disables Lambda Function from being triggered and -1 removes
   * any concurrency limitations. Defaults to Unreserved Concurrency Limits -1.
   */
  reservedConcurrentExecutions?: number;

  /**
   * The canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, bucket-owner-read, and bucket-owner-full-control. Defaults to private.
   */
  s3Acl?: string;

  /**
   * S3 bucket to store artifacts
   */
  s3Bucket?: string;

  /**
   * The S3 bucket object with keys bucket, key, version pointing to an existing zip-file to use
   */
  s3ExistingPackage?: {
    bucket?: string;
    key?: string;
    versionId?: string;
  };

  /**
   * Specifies the desired Storage Class for the artifact uploaded to S3. Can be either STANDARD, REDUCED_REDUNDANCY, ONEZONE_IA, INTELLIGENT_TIERING, or STANDARD_IA.
   */
  s3ObjectStorageClass?: string;

  /**
   * A map of tags to assign to S3 bucket object.
   */
  s3ObjectTags?: Record<string, string>;

  /**
   * Set to true to not merge tags with s3_object_tags. Useful to avoid breaching S3 Object 10 tag limit.
   */
  s3ObjectTagsOnly?: string;

  /**
   * Directory name where artifacts should be stored in the S3 bucket. If unset, the path from artifacts_dir is used
   */
  s3Prefix?: string;

  /**
   * Specifies server-side encryption of the object in S3. Valid values are "AES256" and "aws:kms".
   */
  s3ServerSideEncryption?: string;

  /**
   * (Optional) Snap start settings for low-latency startups
   */
  snapStart?: boolean;

  /**
   * Whether to store produced artifacts on S3 or locally.
   */
  storeOnS3?: boolean;

  /**
   * A map of tags to assign to resources.
   */
  tags?: Record<string, string>;

  /**
   * The amount of time your Lambda Function has to run in seconds.
   */
  timeout?: number;

  /**
   * List of additional trusted entities for assuming Lambda Function role (trust relationship)
   */
  trustedEntities?: ({ type: string; indentifiers: string[] } | { services: string[] })[];

  // Lifecycle rules
  lifecycle?: aws.lambdaFunction.LambdaFunction["lifecycle"];
}

function assertString(value: unknown, msg: string): asserts value is string {
  if (typeof value !== "string") {
    throw new Error(msg);
  }
}

export class StackLambda extends Construct {
  public lambda?: aws.lambdaFunction.LambdaFunction;
  public lambdaLayer?: aws.lambdaLayerVersion.LambdaLayerVersion;
  public role?: aws.iamRole.IamRole;

  constructor(scope: Construct, id: string, props: StackLambdaConfig) {
    super(scope, id);

    // set the defaults
    props = {
      attachCloudwatchLogsPolicy: true,
      authorizationType: "NONE",
      create: true,
      createCurrentVersionAllowedTriggers: true,
      createCurrentVersionAsyncEventConfig: true,
      createFunction: true,
      createPackage: true,
      createRole: true,
      createUnqualifiedAliasAllowedTriggers: true,
      createUnqualifiedAliasAsyncEventConfig: true,
      createUnqualifiedAliasLambdaFunctionUrl: true,
      packageType: "Zip",
      reservedConcurrentExecutions: -1,
      roleForceDetachPolicies: true,
      s3Acl: "private",
      s3ObjectStorageClass: "ONEZONE_IA",
      timeout: 3,
      ...props,
    };

    if (!props.create) {
      return;
    }
    this.createIAM(props);

    if (!this.role?.arn && !props.lambdaRoleArn) {
      throw new Error("either ");
    }
    const role = this.role?.arn ?? props.lambdaRoleArn;
    assertString(role, "no valid role provided");

    let s3package: aws.s3Object.S3Object | undefined;
    const filename = props.asset?.path;

    if (props.createPackage && props.storeOnS3) {
      assertString(props.s3Bucket, "no bucket provided");
      assertString(filename, "no asset provided for s3 uplaod");

      const s3Key = `${props.s3Prefix}${filename.replace(/^.*\//, "")}`;

      s3package = new aws.s3Object.S3Object(this, "package", {
        bucket: props.s3Bucket,
        acl: props.s3Acl,
        key: s3Key,
        source: props.asset?.fileName,
        storageClass: props.s3ObjectStorageClass,
        serverSideEncryption: props.s3ServerSideEncryption,
        tags: props.s3ObjectTagsOnly ? props.s3ObjectTags : { ...props.tags, ...props.s3ObjectTags },
      });
    }

    const s3Key = props.s3ExistingPackage?.key ?? s3package?.key;
    const s3Bucket = props.s3ExistingPackage?.bucket ?? s3package?.bucket;
    const s3ObjectVersion = props.s3ExistingPackage?.versionId ?? s3package?.versionId;

    let cloudwatchLogGroupArn;

    if (props.createFunction && !props.createLayer) {
      this.lambda = new aws.lambdaFunction.LambdaFunction(this, "this", {
        functionName: props.functionName,
        description: props.description,
        role,
        handler: props.packageType == "Zip" ? props.handler : undefined,
        runtime: props.packageType == "Zip" ? props.runtime : undefined,
        memorySize: props.memorySize,

        reservedConcurrentExecutions: props.reservedConcurrentExecutions,
        layers: props.layers,
        timeout: props.lambdaAtEdge ? Math.min(props.timeout ?? 3, 30) : props.timeout,
        publish: props.lambdaAtEdge || props.snapStart || props.publish,
        kmsKeyArn: props.kmsKeyArn,
        imageUri: props.imageUri,
        packageType: props.packageType,
        architectures: [props.architecture],
        codeSigningConfigArn: props.codeSigningConfigArn,
        ...(props.ephemeralStorageSize ? { size: props.ephemeralStorageSize } : {}),
        filename,
        ...(props.ignoreSourceCodeHash ? {} : { sourceCodeHash: props.asset?.assetHash }),
        s3Bucket,
        s3Key,
        s3ObjectVersion,
        ...(props.imageConfigCommand?.length ||
        props.imageConfigEntryPoint?.length ||
        props.imageConfigWorkingDirecdtory
          ? {
              imageConfig: {
                command: props.imageConfigCommand,
                entryPoint: props.imageConfigEntryPoint,
                workingDirectory: props.imageConfigWorkingDirecdtory,
              },
            }
          : {}),
        environment: props.environment,
        ...(props.deadLetterTargetArn
          ? {
              deadLetterConfig: {
                targetArn: props.deadLetterTargetArn,
              },
            }
          : {}),

        //   dynamic "tracing_config" {
        //     for_each = var.tracing_mode == null ? [] : [true]
        //     content {
        //       mode = var.tracing_mode
        //     }
        //   }

        //   dynamic "vpc_config" {
        //     for_each = var.vpc_subnet_ids != null && var.vpc_security_group_ids != null ? [true] : []
        //     content {
        //       security_group_ids = var.vpc_security_group_ids
        //       subnet_ids         = var.vpc_subnet_ids
        //     }
        //   }

        //   dynamic "file_system_config" {
        //     for_each = var.file_system_arn != null && var.file_system_local_mount_path != null ? [true] : []
        //     content {
        //       local_mount_path = var.file_system_local_mount_path
        //       arn              = var.file_system_arn
        //     }
        //   }

        ...(props.snapStart ? { snapStart: { applyOn: "PublishedVersions" } } : {}),
        lifecycle: props.lifecycle ?? {},

        tags: props.tags,
      });

      if (props.useExistingCloudwatchLogGroup) {
        const value = new aws.dataAwsCloudwatchLogGroup.DataAwsCloudwatchLogGroup(this, "lambda-logs", {
          name: `/aws/lambda/${props.lambdaAtEdge ? "us-east-1." : ""}${props.functionName}`,
        });
        cloudwatchLogGroupArn = value.arn;
      } else {
        const value = new aws.cloudwatchLogGroup.CloudwatchLogGroup(this, "lambda-logs", {
          name: `/aws/lambda/${props.lambdaAtEdge ? "us-east-1." : ""}${props.functionName}`,
          retentionInDays: props.cloudwatchLogsRetentionInDays,
          kmsKeyId: props.cloudwatchLogsKmsKeyId,
          tags: { ...props.tags, ...props.cloudwatchLogsTags },
        });
        cloudwatchLogGroupArn = value.arn;
      }

      if (props.provisionedConcurrentExecutions && props.provisionedConcurrentExecutions > -1) {
        new aws.lambdaProvisionedConcurrencyConfig.LambdaProvisionedConcurrencyConfig(this, "current", {
          functionName: this.lambda.functionName,
          qualifier: this.lambda.version,
          provisionedConcurrentExecutions: props.provisionedConcurrentExecutions,
        });
      }

      if (props.createFunctionUrl) {
        new aws.lambdaFunctionUrl.LambdaFunctionUrl(this, "url", {
          functionName: this.lambda.functionName,
          qualifier: props.createUnqualifiedAliasLambdaFunctionUrl ? this.lambda.version : undefined,
          authorizationType: props.authorizationType ?? "NONE",
          cors: props.cors,
        });
      }

      if (
        (props.createCurrentVersionAllowedTriggers || props.createUnqualifiedAliasAllowedTriggers) &&
        props.allowedTriggers
      ) {
        const params = props.allowedTriggers.map((t, idx) => ({
          functionName: this.lambda?.functionName ?? "",
          statementId: t.statementId ?? String(idx),
          action: t.action ?? "lambda:InvokeFunction",
          principal: t.principal ?? `${t.service ?? ""}.amazonaws.com`,
          sourceArn: t.sourceArn,
          sourceAccount: t.sourceAccount,
          eventSourceToken: t.eventSourceToken,
        }));

        if (props.createCurrentVersionAllowedTriggers) {
          params.forEach((item) => {
            new aws.lambdaPermission.LambdaPermission(this, "current${idx}", {
              ...item,
              // Error: "We currently do not support adding policies for $LATEST.",
              // qualifier: this.lambda?.version,
            });
          });
        }
        // Currently not supported by terraform
        // if (props.createUnqualifiedAliasAllowedTriggers) {
        //   params.forEach((item) => {
        //     new aws.lambdaPermission.LambdaPermission(this, "unqual${idx}", {
        //       ...item,
        //     });
        //   });
        // }
      }

      // TODO
      // - aws_lambda_function_event_invoke_config
      // - aws_lambda_event_source_mapping
    }

    //
    if (props.createLayer) {
      assertString(props.layerName, "layerName is missing");

      this.lambdaLayer = new aws.lambdaLayerVersion.LambdaLayerVersion(this, "lambda-layers", {
        layerName: props.layerName,
        description: props.description,
        filename,
        s3Bucket,
        s3Key,
        s3ObjectVersion,

        compatibleRuntimes: props.compatibleRuntimes,
        compatibleArchitectures: props.compatibleArchitectures,

        licenseInfo: props.licenseInfo,
        skipDestroy: props.layerSkipDestroy,
        ...(props.ignoreSourceCodeHash ? {} : { sourceCodeHash: props.asset?.assetHash }),
      });
    }

    this.attachPolicies(props, cloudwatchLogGroupArn);
  }

  createIAM(props: StackLambdaConfig) {
    if (!props.createRole) {
      return;
    }

    const name = props.roleName || props.functionName || "";

    const trustedPrincipals: { type: string; identifiers: string[] }[] = [];

    const servicesScratch: Record<string, boolean> = {
      "lambda.amazonaws.com": true,
      "edgelambda.amazonaws.com": !!props.lambdaAtEdge,
    };
    if (props.trustedEntities) {
      props.trustedEntities.forEach((e) => {
        (((e as any).services as string[]) ?? []).forEach((s) => {
          servicesScratch[s] = true;
        });
        if (e.hasOwnProperty("type")) {
          trustedPrincipals.push(e as any);
        }
      });
    }
    const trustedServices = Object.entries(servicesScratch)
      .filter(([, v]) => v)
      .map(([k]) => k);

    const lambdaPolicy = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "assume_role", {
      statement: [
        {
          effect: "Allow",
          actions: ["sts:AssumeRole"],
          principals: [{ type: "Service", identifiers: trustedServices }, ...trustedPrincipals],
        },

        ...(props.assumeRolePolicyStatements ?? []),
      ],
    });

    this.role = new aws.iamRole.IamRole(this, "lambda", {
      name,
      description: props.roleDescription,
      path: props.rolePath,
      forceDetachPolicies: props.roleForceDetachPolicies,
      permissionsBoundary: props.rolePermissionBoundary,
      assumeRolePolicy: lambdaPolicy.json,
      tags: { ...props.tags, ...props.roleTags },
    });
  }

  attachPolicies(props: StackLambdaConfig, cloudwatchLogGroupArn: string | undefined) {
    if (!this.role || !props.createRole) {
      return;
    }

    const name = props.roleName || props.functionName || "";

    // #################
    // # Cloudwatch Logs
    // #################
    if (props.attachCloudwatchLogsPolicy) {
      const logsPolicyDoc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "logs_doc", {
        statement: [
          {
            effect: "Allow",
            actions: [
              !props.useExistingCloudwatchLogGroup ? "logs:CreateLogGroup" : "",
              "logs:CreateLogStream",
              "logs:PutLogEvents",
            ].filter((v) => !!v),

            resources: [`${cloudwatchLogGroupArn}:*`, `${cloudwatchLogGroupArn}:*:*`],
          },
        ],
      });

      const logsPolicy = new aws.iamPolicy.IamPolicy(this, "logs_policy", {
        name: `${name}-logs`,
        path: props.policyPath,
        policy: logsPolicyDoc.json,
        tags: props.tags,
      });

      new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, "logs", {
        role: this.role.id,
        policyArn: logsPolicy.arn,
      });
    }

    // #####################
    // # Dead Letter Config
    // #####################
    if (props.attachDeadLetterPolicy && props.deadLetterTargetArn) {
      const dlPolicyDoc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "dead_letter_doc", {
        statement: [
          {
            effect: "Allow",

            actions: ["sns:Publish", "sqs:SendMessage"],

            resources: [props.deadLetterTargetArn],
          },
        ],
      });

      const dlPolicy = new aws.iamPolicy.IamPolicy(this, "dead_letter_policy", {
        name: `${name}-dl`,
        path: props.policyPath,
        policy: dlPolicyDoc.json,
        tags: props.tags,
      });

      new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, "dead_letter", {
        role: this.role.id,
        policyArn: dlPolicy.arn,
      });
    }

    // ######
    // # VPC
    // ######
    if (props.attachNetworkPolicy) {
      const partition = new aws.dataAwsPartition.DataAwsPartition(this, "current");

      const vpcPolicyDoc = new aws.dataAwsIamPolicy.DataAwsIamPolicy(this, "vpc_data", {
        arn: `arn:${partition.partition}:iam::aws:policy/service-role/AWSLambdaENIManagementAccess`,
      });

      const vpcPolicy = new aws.iamPolicy.IamPolicy(this, "vpc_policy", {
        name: `${name}-dl`,
        path: props.policyPath,
        policy: vpcPolicyDoc.policy,
        tags: props.tags,
      });

      new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, "vpc", {
        role: this.role.id,
        policyArn: vpcPolicy.arn,
      });
    }

    // #####################################
    // # Additional policies (list of JSON)
    // #####################################

    if (props.attachPolicies) {
      props.attachPolicies.forEach((policy, idx) => {
        if (!this || !this.role) {
          throw new Error("this not defined");
        }
        // const policy = new aws.iamPolicy.IamPolicy(this, `custom_policy_${idx}`, {
        //   name: `${name}-custom-${idx}`,
        //   path: props.policyPath,
        //   policy: element,
        //   tags: props.tags,
        // });

        new aws.iamPolicyAttachment.IamPolicyAttachment(this, `custom_attach_${idx}`, {
          name: `${name}-${idx}`,
          roles: [this.role.id],
          policyArn: policy.arn,
        });
      });
    }

    // TODO -- more things from here
    //   https://github.com/terraform-aws-modules/terraform-aws-lambda/blob/master/iam.tf#L213
  }
}

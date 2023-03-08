import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";

export interface Props
  extends Pick<aws.s3Bucket.S3BucketConfig, "bucketPrefix" | "bucket" | "forceDestroy" | "website"> {
  expiresInDays?: number;
  enableCors?: boolean;
  public?: boolean;
}

export class StateS3Bucket extends Construct {
  public bucket: aws.s3Bucket.S3Bucket;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    const bucket = new aws.s3Bucket.S3Bucket(this, "bucket", {
      ...props,
    });
    this.bucket = bucket;

    new aws.s3BucketServerSideEncryptionConfiguration.S3BucketServerSideEncryptionConfigurationA(this, "encryption", {
      bucket: bucket.bucket,
      rule: [
        {
          applyServerSideEncryptionByDefault: {
            sseAlgorithm: "AES256",
          },
        },
      ],
    });

    new aws.s3BucketPublicAccessBlock.S3BucketPublicAccessBlock(this, "publicaccess", {
      bucket: bucket.bucket,
      blockPublicAcls: !props.public,
      blockPublicPolicy: !props.public,
      ignorePublicAcls: !props.public,
      restrictPublicBuckets: !props.public,
    });

    new aws.s3BucketOwnershipControls.S3BucketOwnershipControls(this, "ownership", {
      bucket: bucket.bucket,
      rule: {
        objectOwnership: "BucketOwnerEnforced",
      },
    });

    if (props.expiresInDays) {
      new aws.s3BucketLifecycleConfiguration.S3BucketLifecycleConfiguration(this, "lifecycle", {
        bucket: bucket.bucket,
        rule: [
          {
            id: "everything",
            status: "Enabled",
            expiration: {
              days: props.expiresInDays,
            },
          },
        ],
      });
    }

    if (props.enableCors) {
      new aws.s3BucketCorsConfiguration.S3BucketCorsConfiguration(this, "cors", {
        bucket: bucket.bucket,
        corsRule: [
          {
            allowedMethods: ["GET", "PUT"],
            allowedOrigins: ["*"],
            allowedHeaders: ["*"],
          },
        ],
      });
    }
  }
}

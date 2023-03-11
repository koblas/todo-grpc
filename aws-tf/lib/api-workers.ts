import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { GoHandler } from "./components/gohandler";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
}

export class WorkerFile extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      privateBucket,
      publicBucket,
      uploadBucket,
    }: {
      eventbus: aws.snsTopic.SnsTopic;
      uploadBucket: aws.s3Bucket.S3Bucket;
      publicBucket: aws.s3Bucket.S3Bucket;
      privateBucket: aws.s3Bucket.S3Bucket;
    },
  ) {
    super(scope, id);

    const handler = new GoHandler(this, "worker-file", {
      path: ["workers", "file"],
      memorySize: 512,
      environment: {
        variables: {
          BUS_ENTITY_ARN: eventbus.arn,
          PRIVATE_BUCKET: privateBucket.bucket,
          PUBLIC_BUCKET: publicBucket.bucket,
        },
      },
      eventbus,
      parameters: ["/common/*"],
      s3buckets: [uploadBucket, publicBucket],
    });

    handler.eventQueue("worker-file", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/twirp/corepb.v1.FileEventbus/FileUploaded"],
      }),
    });
  }
}

export class WorkerUser extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
    }: {
      eventbus: aws.snsTopic.SnsTopic;
    },
  ) {
    super(scope, id);

    const handler = new GoHandler(this, "worker-user", {
      path: ["workers", "user"],
      memorySize: 512,
      environment: {
        variables: {
          BUS_ENTITY_ARN: eventbus.arn,
        },
      },
      parameters: ["/common/*"],
    });

    handler.eventQueue("worker-user", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": [
          "/twirp/corepb.v1.UserEventbus/SecurityPasswordChange",
          "/twirp/corepb.v1.UserEventbus/SecurityRegisterToken",
          "/twirp/corepb.v1.UserEventbus/SecurityForgotRequest",
          "/twirp/corepb.v1.UserEventbus/SecurityInviteToken",
        ],
      }),
    });
  }
}

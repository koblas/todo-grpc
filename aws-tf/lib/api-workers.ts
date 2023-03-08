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
          PRIVATE_BUCKET: privateBucket.bucket,
          PUBLIC_BUCKET: publicBucket.bucket,
        },
      },
      s3buckets: [uploadBucket, publicBucket],
    });

    handler.eventQueue("websocket-todo", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/twirp/corepbv1.eventbus.FileEventbus/FileUploaded"],
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
    });

    handler.eventQueue("websocket-todo", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": [
          "/twirp/corepbv1.eventbus.UserEventbus/SecurityPasswordChange",
          "/twirp/corepbv1.eventbus.UserEventbus/SecurityRegisterToken",
          "/twirp/corepbv1.eventbus.UserEventbus/SecurityForgotRequest",
          "/twirp/corepbv1.eventbus.UserEventbus/SecurityInviteToken",
        ],
      }),
    });
  }
}

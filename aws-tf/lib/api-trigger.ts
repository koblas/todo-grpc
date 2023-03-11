import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { GoHandler } from "./components/gohandler";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
}

export class TriggerS3 extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { eventbus, bucket }: { eventbus: aws.snsTopic.SnsTopic; bucket: aws.s3Bucket.S3Bucket },
  ) {
    super(scope, id);

    const handler = new GoHandler(this, "trigger-s3", {
      path: ["trigger", "s3"],
      eventbus,
      parameters: ["/common/*"],
      allowedTriggers: [
        {
          action: "lambda:InvokeFunction",
          principal: "s3.amazonaws.com",
          sourceArn: bucket.arn,
        },
      ],
      s3buckets: [bucket],
      environment: {
        variables: {
          BUS_ENTITY_ARN: eventbus.arn,
        },
      },
    });

    if (!handler.role) {
      throw new Error("role not created");
    }

    new aws.s3BucketNotification.S3BucketNotification(this, `${id}-events`, {
      bucket: bucket.id,
      lambdaFunction: [
        {
          events: ["s3:ObjectCreated:*"],
          lambdaFunctionArn: handler.lambda.arn,
        },
      ],

      dependsOn: [handler.role],
    });
  }
}

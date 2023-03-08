import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { GoHandler } from "./api-utils";

export class TriggerS3 extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { eventbus, bucket }: { eventbus: cdk.aws_sns.Topic; bucket: cdk.aws_s3.IBucket },
  ) {
    super(scope, id);

    const handler = new GoHandler(this, "trigger-s3", ["trigger", "s3"], { eventbus, parameters: ["/common/*"] });

    const notificaiton = new cdk.aws_s3_notifications.LambdaDestination(handler.lambda);
    notificaiton.bind(scope, bucket);
    bucket.addEventNotification(cdk.aws_s3.EventType.OBJECT_CREATED, notificaiton);
  }
}

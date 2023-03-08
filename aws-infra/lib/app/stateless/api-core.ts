import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { GoHandler } from "./api-utils";

export class CoreTodo extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const db = new cdk.aws_dynamodb.Table(this, "db", {
      tableName: "app-todo",
      billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
      sortKey: { name: "sk", type: cdk.aws_dynamodb.AttributeType.STRING },
    });

    new GoHandler(this, "core-todo", ["core", "todo"], {
      eventbus,
      parameters: ["/common/*"],
      dynamo: db,
    });
  }
}

export class CoreUser extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const db = new cdk.aws_dynamodb.Table(this, "db", {
      tableName: "app-user",
      billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
    });

    new GoHandler(
      this,
      "core-user",
      ["core", "user"],

      { eventbus, parameters: ["/common/*"], dynamo: db },
    );
  }
}

export class CoreOauthUser extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    new GoHandler(this, "core-oauth_user", ["core", "oauth_user"], { parameters: ["/common/*", "/oauth/*"], eventbus });
  }
}

// export class CoreSendEmail extends Construct {
//   constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
//     super(scope, id);

//     const { eventbus } = props;

//     const lambda = new cdk.aws_lambda.Function(this, "lambda", {
//       functionName: "core-send-email",
//       code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
//       runtime: cdk.aws_lambda.Runtime.GO_1_X,
//       handler: "core-send-email",
//     });

//     wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });
//   }
// }

export class CoreSendEmailQueue extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const handler = new GoHandler(this, "core-send_email", ["core", "send_email"]);

    handler.listenQueue("core-send-email", { queueName: "send-email" });
  }
}

///
//
//

export class TriggerS3 extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { eventbus, bucket }: { eventbus: cdk.aws_sns.Topic; bucket: cdk.aws_s3.Bucket },
  ) {
    super(scope, id);

    const handler = new GoHandler(
      this,
      "trigger-s3",
      ["trigger", "s3"],

      { eventbus, parameters: ["/common/*"] },
    );

    const notificaiton = new cdk.aws_s3_notifications.LambdaDestination(handler.lambda);
    notificaiton.bind(scope, bucket);
    bucket.addEventNotification(cdk.aws_s3.EventType.OBJECT_CREATED, notificaiton);
  }
}

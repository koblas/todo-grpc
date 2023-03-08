import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { GoHandler } from "./components/gohandler";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
}

export class CoreTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const db = new aws.dynamodbTable.DynamodbTable(this, "db", {
      name: "app-todo",
      billingMode: "PAY_PER_REQUEST",
      // partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },

      attribute: [
        { name: "pk", type: "S" },
        { name: "sk", type: "S" },
      ],
      hashKey: "pk",
      rangeKey: "sk",
    });

    new GoHandler(this, "core-todo", {
      path: ["core", "todo"],
      eventbus,
      parameters: ["/common/*"],
      dynamo: db,
    });
  }
}

export class CoreUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const db = new aws.dynamodbTable.DynamodbTable(this, "db", {
      name: "app-user",
      billingMode: "PAY_PER_REQUEST",
      // partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },

      attribute: [
        { name: "pk", type: "S" },
        // { name: "sk", type: "S" },
      ],
      hashKey: "pk",
      // rangeKey: "sk",
    });

    new GoHandler(this, "core-user", {
      path: ["core", "user"],
      eventbus,
      parameters: ["/common/*"],
      dynamo: db,
    });
  }
}

export class CoreOauthUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    new GoHandler(this, "core-oauth-user", {
      path: ["core", "oauth-user"],
      eventbus,
      parameters: ["/common/*", "/oauth/*"],
    });
  }
}

export class CoreSendEmailQueue extends Construct {
  constructor(scope: Construct, id: string, _: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const handler = new GoHandler(this, "core-send-email", { path: ["core", "send-email"] });

    handler.listenQueue("send-email");
  }
}
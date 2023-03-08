import * as cdk from "aws-cdk-lib";
import { IWebSocketApi, IWebSocketStage, WebSocketApi, WebSocketStage } from "@aws-cdk/aws-apigatewayv2-alpha";
import { Construct } from "constructs";
import { SqsToLambda } from "@aws-solutions-constructs/aws-sqs-lambda";
import { SubscriptionFilter } from "aws-cdk-lib/aws-sns";
import { GoHandler } from "./api-utils";

export class WebsocketTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const handler = new GoHandler(this, "websocket-todo", ["websocket", "todo"], {
      eventbus,
      parameters: ["/common/*"],
    });

    handler.eventQueue("websocket-todo", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.TodoEventbus/TodoChange"],
        }),
      },
    });
  }
}

export class WebsocketUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const handler = new GoHandler(this, "websocket-user", ["websocket", "user"], {
      parameters: ["/common/*"],
      eventbus,
    });

    handler.eventQueue("websocket-user", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/UserChange"],
        }),
      },
    });
  }
}

export class WebsocketBroadcast extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      wsstage,
      wsapi,
      db,
    }: { db: cdk.aws_dynamodb.ITable; eventbus: cdk.aws_sns.Topic; wsstage: WebSocketStage; wsapi: IWebSocketApi },
  ) {
    super(scope, id);

    // const db = cdk.aws_dynamodb.Table.fromTableName(this, "conns", "ws-connection");

    const handler = new GoHandler(this, "websocket-broadcast", ["websocket", "broadcast"], {
      environment: {
        CONN_DB: db.tableName,
        WS_ENDPOINT: wsstage.callbackUrl,
      },
      parameters: ["/common/*"],
      dynamo: db,
    });

    wsstage.grantManagementApiAccess(handler.lambda);

    handler.eventQueue("websocket-broadcast", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.BroadcastEventbus/Send"],
        }),
      },
    });
  }
}

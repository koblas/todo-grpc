import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { GoHandler } from "./components/gohandler";
import { WebsocketConfig } from "./gw-websocket";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
  eventbus: aws.snsTopic.SnsTopic;
}

export class WebsocketFile extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const handler = new GoHandler(this, "websocket-file", {
      path: ["websocket", "file"],
      eventbus,
      parameters: ["/common/*"],
      environment: {
        variables: {
          BUS_ENTITY_ARN: eventbus.arn,
        },
      },
    });

    handler.eventQueue("websocket-file", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/core.v1.FileEventbusService/FileCompleted"],
      }),
    });
  }
}

export class WebsocketTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const handler = new GoHandler(this, "websocket-todo", {
      path: ["websocket", "todo"],
      eventbus,
      parameters: ["/common/*"],
      environment: {
        variables: {
          BUS_ENTITY_ARN: eventbus.arn,
        },
      },
    });

    handler.eventQueue("websocket-todo", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/core.v1.TodoEventbusService/TodoChange"],
      }),
    });
  }
}

export class WebsocketUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const handler = new GoHandler(this, "websocket-user", {
      path: ["websocket", "user"],
      eventbus,
      parameters: ["/common/*"],
      environment: {
        variables: {
          BUS_ENTITY_ARN: eventbus.arn,
        },
      },
    });

    handler.eventQueue("websocket-user", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/core.v1.UserEventbusService/UserChange"],
      }),
    });
  }
}

export class WebsocketBroadcast extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      wsconf,
    }: {
      wsconf: WebsocketConfig;
      eventbus: aws.snsTopic.SnsTopic;
    },
  ) {
    super(scope, id);

    const region = new aws.dataAwsRegion.DataAwsRegion(this, "rcurrent");
    const account = new aws.dataAwsCallerIdentity.DataAwsCallerIdentity(this, "acurrent");

    const mgtdoc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "gwdoc", {
      statement: [
        {
          effect: "Allow",
          actions: ["execute-api:ManageConnections"],
          resources: [
            // `arn:aws:execute-api:${region.name}:${account.accountId}:${wsapi.id}/${wsstage.name}/*/@connections`,
            `arn:aws:execute-api:${region.name}:${account.accountId}:${wsconf.wsapi.id}/${wsconf.wsstage.name}/*/@connections/*`,
          ],
        },
      ],
    });

    const handler = new GoHandler(this, "websocket-broadcast", {
      path: ["websocket", "broadcast"],
      environment: {
        variables: {
          CONN_DB: wsconf.wsdb.name,
          WS_ENDPOINT: wsconf.wsCallbackUrl,
          BUS_ENTITY_ARN: eventbus.arn,
        },
      },
      attachPolicies: [
        new aws.iamPolicy.IamPolicy(this, "gwpolicy", {
          policy: mgtdoc.json,
        }),
      ],
      parameters: ["/common/*"],
      dynamo: wsconf.wsdb,
    });

    handler.eventQueue("websocket-broadcast", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/core.v1.BroadcastEventbusService/Send"],
      }),
    });
  }
}

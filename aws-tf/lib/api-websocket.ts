import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { GoHandler } from "./components/gohandler";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
  eventbus: aws.snsTopic.SnsTopic;
}

export class WebsocketTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: aws.snsTopic.SnsTopic }) {
    super(scope, id);

    const handler = new GoHandler(this, "websocket-todo", {
      path: ["websocket", "todo"],
      eventbus,
      parameters: ["/common/*"],
    });

    handler.eventQueue("websocket-todo", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/twirp/corepbv1.eventbus.TodoEventbus/TodoChange"],
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
    });

    handler.eventQueue("websocket-todo", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/twirp/corepbv1.eventbus.UserEventbus/UserChange"],
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
      wsstage,
      wsapi,
      db,
    }: {
      db: aws.dynamodbTable.DynamodbTable;
      eventbus: aws.snsTopic.SnsTopic;
      wsstage: aws.apigatewayv2Stage.Apigatewayv2Stage;
      wsapi: aws.apigatewayv2Api.Apigatewayv2Api;
    },
  ) {
    super(scope, id);

    const region = new aws.dataAwsRegion.DataAwsRegion(this, "rcurrent");
    const account = new aws.dataAwsCallerIdentity.DataAwsCallerIdentity(this, "acurrent");
    const callbackUrl = `https://${wsapi.id}.execute-api.${region.name}.com/${wsstage.name}`;

    const mgtdoc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "gwdoc", {
      statement: [
        {
          effect: "Allow",
          actions: ["execute-api:ManageConnections"],
          resources: [
            `arn:aws:execute-api:${region.name}:${account.accountId}:${wsapi.id}/${wsstage.name}/*/@connections/*`,
          ],
        },
      ],
    });

    const handler = new GoHandler(this, "websocket-broadcast", {
      path: ["websocket", "broadcast"],
      environment: {
        variables: {
          CONN_DB: db.name,
          WS_ENDPOINT: callbackUrl,
        },
      },
      attachPolicies: [
        new aws.iamPolicy.IamPolicy(this, "gwpolicy", {
          policy: mgtdoc.json,
        }),
      ],
      parameters: ["/common/*"],
      dynamo: db,
    });

    // TODO TODO TODO
    // const arn = aws_cdk_lib_1.Stack.of(this.api).formatArn({
    //   service: "execute-api",
    //   resource: this.api.apiId,
    // });
    // return aws_iam_1.Grant.addToPrincipal({
    //   grantee: identity,
    //   actions: ["execute-api:ManageConnections"],
    //   resourceArns: [`${arn}/${this.stageName}/*/@connections/*`],
    // });

    // const policy = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "policy", {
    //   statement: [
    //     {
    //       effect: "Allow",
    //       principals: [
    //         {
    //           type: "Service",
    //           identifiers: [handler.lambda.arn],
    //         },
    //       ],
    //       actions: ["execute-api:ManageConnections"],
    //       resources: [wsapi.arn],
    //     },
    //   ],
    // });

    // wsstage.grantManagementApiAccess(handler.lambda);

    handler.eventQueue("websocket-broadcast", eventbus, {
      filterPolicy: JSON.stringify({
        "twirp.path": ["/twirp/corepbv1.eventbus.BroadcastEventbus/Send"],
      }),
    });
  }
}

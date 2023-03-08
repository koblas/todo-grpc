import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { DomainName, WebSocketApi, WebSocketStage } from "@aws-cdk/aws-apigatewayv2-alpha";
import { WebSocketLambdaIntegration } from "@aws-cdk/aws-apigatewayv2-integrations-alpha";
import { GoHandler } from "./api-utils";
import { MockPing } from "../utils/mockping";
import { CertificateStack } from "../utils/certificate";

export class WebsocketHandler extends Construct {
  public wsapi: WebSocketApi;
  public wsstage: WebSocketStage;
  public db: cdk.aws_dynamodb.Table;

  constructor(scope: Construct, id: string, props: { hostedZone: cdk.aws_route53.IHostedZone; hostname: string }) {
    super(scope, id);

    const domainName = `${props.hostname}.${props.hostedZone.zoneName}`;

    const { certificate } = new CertificateStack(this, "cert", {
      hostedZone: props.hostedZone,
      domainName,
    });

    const dn = new DomainName(this, "wsDn", {
      domainName,
      certificate,
    });

    this.db = new cdk.aws_dynamodb.Table(this, "db", {
      tableName: "ws-connection",
      billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
      sortKey: { name: "sk", type: cdk.aws_dynamodb.AttributeType.STRING },
      timeToLiveAttribute: "delete_at",
    });

    const handler = new GoHandler(this, "publicapi-websocket", ["publicapi", "websocket"], {
      environment: {
        CONN_DB: this.db.tableName,
      },
      parameters: ["/common/*"],
      dynamo: this.db,
    });

    this.wsapi = new WebSocketApi(this, "wsapi", {
      connectRouteOptions: { integration: new WebSocketLambdaIntegration("connect", handler.lambda) },
      disconnectRouteOptions: { integration: new WebSocketLambdaIntegration("disconnect", handler.lambda) },
      defaultRouteOptions: { integration: new WebSocketLambdaIntegration("default", handler.lambda) },
    });

    this.wsstage = new WebSocketStage(this, "DefaultStage", {
      webSocketApi: this.wsapi,
      stageName: "$default",
      autoDeploy: true,
      domainMapping: { domainName: dn },
    });

    // Setup two websocket responses
    new MockPing(this, "mockPing", { routeKey: "ping", sockapi: this.wsapi });
    new MockPing(this, "mockCursor", {
      routeKey: "cursor",
      sockapi: this.wsapi,
      responseJson: JSON.stringify({ statusCode: 200 }),
    });

    new cdk.aws_route53.ARecord(this, "Alias", {
      zone: props.hostedZone,
      recordName: props.hostname,
      target: cdk.aws_route53.RecordTarget.fromAlias(
        new cdk.aws_route53_targets.ApiGatewayv2DomainProperties(dn.regionalDomainName, dn.regionalHostedZoneId),
      ),
    });
  }
}

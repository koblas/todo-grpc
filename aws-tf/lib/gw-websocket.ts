import * as aws from "@cdktf/provider-aws";
import { Fn } from "cdktf";
import { Construct } from "constructs";
import { CertificateDomain } from "./components/certificate";
import { GoHandler } from "./components/gohandler";

interface Props {
  zone: { zoneId: string };
  wsHostname: string;
  domainName: string;
}
export interface WebsocketConfig {
  wsapi: aws.apigatewayv2Api.Apigatewayv2Api;
  wsstage: aws.apigatewayv2Stage.Apigatewayv2Stage;
  wsdb: aws.dynamodbTable.DynamodbTable;
  wsdomainName: string;
  wsCallbackUrl: string;
}

export class WebsocketHandler extends Construct {
  public wsconf: WebsocketConfig;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    this.wsconf = {} as any;
    const region = new aws.dataAwsRegion.DataAwsRegion(this, "current");

    const policyAssume = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "assume", {
      statement: [
        {
          effect: "Allow",
          actions: ["sts:AssumeRole"],
          principals: [{ identifiers: ["apigateway.amazonaws.com"], type: "Service" }],
        },
      ],
    });

    const logsrole = new aws.iamRole.IamRole(this, "wslogrole", {
      assumeRolePolicy: policyAssume.json,
      managedPolicyArns: [
        new aws.dataAwsIamPolicy.DataAwsIamPolicy(this, "wspo", {
          name: "AmazonAPIGatewayPushToCloudWatchLogs",
        }).arn,
      ],
    });

    this.wsconf.wsapi = new aws.apigatewayv2Api.Apigatewayv2Api(this, "wsapi", {
      name: "websocket-api",
      protocolType: "WEBSOCKET",
      routeSelectionExpression: "$request.body.action",
      credentialsArn: logsrole.arn,
    });

    this.wsconf.wsdomainName = `${props.wsHostname}.${props.domainName}`;

    const cert = new CertificateDomain(this, "ws", {
      domainName: this.wsconf.wsdomainName,
      zoneId: props.zone.zoneId,
    });

    const dn = new aws.apigatewayv2DomainName.Apigatewayv2DomainName(this, "wsdn", {
      domainName: this.wsconf.wsdomainName,
      domainNameConfiguration: {
        certificateArn: cert.cert.arn,
        endpointType: "REGIONAL",
        securityPolicy: "TLS_1_2",
      },
    });

    new aws.route53Record.Route53Record(this, "wsalias", {
      name: props.wsHostname,
      type: "A",
      zoneId: props.zone.zoneId,
      alias: {
        name: dn.domainNameConfiguration.targetDomainName,
        zoneId: dn.domainNameConfiguration.hostedZoneId,
        evaluateTargetHealth: true,
      },
    });

    // const dn = new route53Record.Route53Record(this, "wsDn", {
    //   name: domainName,
    //   domainName,
    //   certificate,
    // });

    this.wsconf.wsdb = new aws.dynamodbTable.DynamodbTable(this, "db", {
      name: "ws-connection",
      billingMode: "PAY_PER_REQUEST",
      // partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
      // sortKey: { name: "sk", type: cdk.aws_dynamodb.AttributeType.STRING },
      attribute: [
        { name: "pk", type: "S" },
        { name: "sk", type: "S" },
      ],
      hashKey: "pk",
      rangeKey: "sk",

      ttl: {
        enabled: true,
        attributeName: "delete_at",
      },
      // billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      // timeToLiveAttribute: "delete_at",
    });

    const handler = new GoHandler(this, "publicapi-websocket", {
      // path: ["publicapi", "websocket"],
      apiTrigger: this.wsconf.wsapi,
      environment: {
        variables: {
          CONN_DB: this.wsconf.wsdb.name,
        },
      },
      parameters: ["/common/*"],
      dynamo: this.wsconf.wsdb,
    });

    const policyInvoke = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "invoke", {
      statement: [
        {
          effect: "Allow",
          actions: ["lambda:InvokeFunction"],
          resources: [handler.lambda.arn],
        },
      ],
    });

    const gwpolicy = new aws.iamPolicy.IamPolicy(this, "wspolicy", {
      name: "WsMessageApiGatewayPolicy",
      path: "/",
      policy: policyInvoke.json,
    });

    const role = new aws.iamRole.IamRole(this, "wsrole", {
      name: "WsMessengerApiGatewayRole",
      assumeRolePolicy: policyAssume.json,
      managedPolicyArns: [gwpolicy.arn],
    });

    const logs = new aws.cloudwatchLogGroup.CloudwatchLogGroup(this, "logs", {
      name: "/apigw/wsapi",
      retentionInDays: 3,
    });

    this.wsconf.wsstage = new aws.apigatewayv2Stage.Apigatewayv2Stage(this, "stage", {
      apiId: this.wsconf.wsapi.id,
      description: "Default Route",
      name: "$default",
      autoDeploy: true,
      // accessLogSettings: {
      //   destinationArn: logs.arn,
      //   format: JSON.stringify({
      //     requestId: "$context.requestId",
      //     extendedRequestId: "$context.extendedRequestId",
      //     ip: "$context.identity.sourceIp",
      //     caller: "$context.identity.caller",
      //     user: "$context.identity.user",
      //     requestTime: "$context.requestTime",
      //     httpMethod: "$context.httpMethod",
      //     resourcePath: "$context.resourcePath",
      //     status: "$context.status",
      //     protocol: "$context.protocol",
      //     responseLength: "$context.responseLength",
      //   }),
      // },
      defaultRouteSettings: {
        throttlingBurstLimit: 1000,
        throttlingRateLimit: 100,
        dataTraceEnabled: false,
        loggingLevel: "OFF",
        detailedMetricsEnabled: false,
      },
      lifecycle: {
        ignoreChanges: ["deployment_id"],
      },
      dependsOn: [logs, logsrole],

      // stageVariables: "$default",
      // domainMapping: { domainName: dn },
    });

    new aws.apigatewayv2ApiMapping.Apigatewayv2ApiMapping(this, "wsmap", {
      apiId: this.wsconf.wsapi.id,
      stage: this.wsconf.wsstage.id,
      domainName: dn.id,
    });

    const integration = new aws.apigatewayv2Integration.Apigatewayv2Integration(this, "integration", {
      apiId: this.wsconf.wsapi.id,
      integrationType: "AWS_PROXY",
      integrationUri: handler.lambda.invokeArn,
      credentialsArn: role.arn,
      contentHandlingStrategy: "CONVERT_TO_TEXT",
      passthroughBehavior: "WHEN_NO_MATCH",
    });

    new aws.apigatewayv2IntegrationResponse.Apigatewayv2IntegrationResponse(this, "response", {
      apiId: this.wsconf.wsapi.id,
      integrationId: integration.id,
      integrationResponseKey: "/200/",
    });

    ["default", "connect", "disconnect"].forEach((key) => {
      const dRoute = new aws.apigatewayv2Route.Apigatewayv2Route(this, `ws${key}`, {
        apiId: this.wsconf.wsapi.id,
        routeKey: `$${key}`,
        target: `integrations/${integration.id}`,
      });
      new aws.apigatewayv2RouteResponse.Apigatewayv2RouteResponse(this, `wsr${key}`, {
        apiId: this.wsconf.wsapi.id,
        routeId: dRoute.id,
        routeResponseKey: "$default",
      });
    });

    // Setup two websocket responses
    new MockPing(this, "ping", { routeKey: "ping", sockapi: this.wsconf.wsapi });
    new MockPing(this, "cursor", {
      routeKey: "cursor",
      sockapi: this.wsconf.wsapi,
      responseJson: JSON.stringify({ statusCode: 200 }),
    });

    this.wsconf.wsCallbackUrl = `https://${this.wsconf.wsapi.id}.execute-api.${region.name}.amazonaws.com/${this.wsconf.wsstage.name}`;
  }
}

export class MockPing extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      sockapi,
      routeKey,
      responseJson,
    }: { routeKey: string; sockapi: aws.apigatewayv2Api.Apigatewayv2Api; responseJson?: string },
  ) {
    super(scope, id);

    const integration = new aws.apigatewayv2Integration.Apigatewayv2Integration(this, "integration", {
      apiId: sockapi.id,
      integrationType: "MOCK",
      templateSelectionExpression: "200",
      requestTemplates: {
        "200": '{"statusCode":200}',
      },
      passthroughBehavior: "WHEN_NO_MATCH",
    });
    const route = new aws.apigatewayv2Route.Apigatewayv2Route(this, "route", {
      apiId: sockapi.id,
      routeKey,
      operationName: "pingRoute",
      routeResponseSelectionExpression: "$default",
      target: Fn.join("/", ["integrations", integration.id]),
    });

    if (responseJson) {
      new aws.apigatewayv2IntegrationResponse.Apigatewayv2IntegrationResponse(this, "response", {
        apiId: sockapi.id,
        integrationId: integration.id,
        integrationResponseKey: "/200/",
        responseTemplates: {
          "200": responseJson,
        },
      });
      new aws.apigatewayv2RouteResponse.Apigatewayv2RouteResponse(this, "routeResponse", {
        apiId: sockapi.id,
        routeId: route.id,
        routeResponseKey: "$default",
      });
    }
  }
}

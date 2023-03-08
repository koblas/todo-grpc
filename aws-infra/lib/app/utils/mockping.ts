import * as cdk from "aws-cdk-lib";
import { WebSocketApi } from "@aws-cdk/aws-apigatewayv2-alpha";
import { Construct } from "constructs";

export class MockPing extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { sockapi, routeKey, responseJson }: { routeKey: string; sockapi: WebSocketApi; responseJson?: string },
  ) {
    super(scope, id);

    const intgration = new cdk.aws_apigatewayv2.CfnIntegration(this, "integration", {
      apiId: sockapi.apiId,
      integrationType: "MOCK",
      requestTemplates: {
        "200": '{"statusCode":200}',
      },
      templateSelectionExpression: "200",
      passthroughBehavior: "WHEN_NO_MATCH",
    });
    const route = new cdk.aws_apigatewayv2.CfnRoute(this, "route", {
      apiId: sockapi.apiId,
      routeKey,
      routeResponseSelectionExpression: "$default",
      operationName: "pingRoute",
      target: new cdk.StringConcat().join("integrations/", intgration.ref),
    });
    if (responseJson) {
      new cdk.aws_apigatewayv2.CfnIntegrationResponse(this, "response", {
        apiId: sockapi.apiId,
        integrationId: intgration.ref,
        integrationResponseKey: "/200/",
        responseTemplates: {
          "200": responseJson,
        },
      });
      new cdk.aws_apigatewayv2.CfnRouteResponse(this, "routeResponse", {
        apiId: sockapi.apiId,
        routeId: route.ref,
        routeResponseKey: "$default",
      });
    }
  }
}

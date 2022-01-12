import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";
import { buildLambda } from "./lambda-util";

export async function publicTodo(corestack: pulumi.StackReference) {
  const { lambda } = await buildLambda("public-todo", {
    code: new pulumi.asset.AssetArchive({
      app: new pulumi.asset.FileAsset("../build/publicapi-todo"),
    }),
  });

  const appgw = aws.apigatewayv2.Api.get("appgw-todo", corestack.getOutputValue("apigwId"));

  const integration = new aws.apigatewayv2.Integration("publicapi-todo", {
    description: "Public API Integration",
    apiId: appgw.id,
    integrationType: "AWS_PROXY",
    connectionType: "INTERNET",
    integrationMethod: "POST",
    integrationUri: lambda.invokeArn,
    payloadFormatVersion: "2.0",
    requestParameters: {
      "overwrite:path": "/twirp/api.todo.TodoService/$request.path.proxy",
    },
    //     contentHandlingStrategy: "CONVERT_TO_TEXT",
    //     passthroughBehavior: "WHEN_NO_MATCH",
  });

  const route = new aws.apigatewayv2.Route("public-todo", {
    apiId: appgw.id,
    routeKey: "POST /v1/todo/{proxy+}",
    target: integration.id.apply((id) => `integrations/${id}`),
  });
  const apigwExecuteArn = await corestack.getOutputValue("apigwExecuteArn");

  // Give API Gateway permissions to invoke the Lambda
  new aws.lambda.Permission("lambdaPermission-todo", {
    action: "lambda:InvokeFunction",
    principal: "apigateway.amazonaws.com",
    function: lambda,
    sourceArn: route.routeKey.apply((key) => {
      return `${apigwExecuteArn}*/*${key.split(" ")[1]}`;
    }),
  });
}

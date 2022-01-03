import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";
import { buildLambda } from "./lambda-util";

export async function publicAuth(corestack: pulumi.StackReference) {
  const redisHost = await corestack.getOutputValue("redisHost");
  const redisPort = await corestack.getOutputValue("redisPort");

  const redisAddr = redisHost && redisPort ? `${redisHost}:${redisPort}` : "";

  const { lambda } = await buildLambda("public-auth", {
    code: new pulumi.asset.AssetArchive({
      app: new pulumi.asset.FileAsset("../build/publicapi-auth"),
    }),
    environment: {
      variables: {
        REDIS_ADDR: redisAddr,
      },
    },
  });

  const appgw = aws.apigatewayv2.Api.get("appgw", corestack.getOutputValue("apigwId"));

  const publicAuthIntegration = new aws.apigatewayv2.Integration("publicapi-auth", {
    description: "Public API Integration",
    apiId: appgw.id,
    integrationType: "AWS_PROXY",
    connectionType: "INTERNET",
    integrationMethod: "POST",
    integrationUri: lambda.invokeArn,
    payloadFormatVersion: "2.0",
    //     contentHandlingStrategy: "CONVERT_TO_TEXT",
    passthroughBehavior: "WHEN_NO_MATCH",
    requestParameters: {
      "overwrite:path": "/twirp/api.auth.AuthenticationService/$request.path.proxy",
    },
  });

  const route = new aws.apigatewayv2.Route("public-auth", {
    apiId: appgw.id,
    routeKey: "POST /v1/auth/{proxy+}",
    target: publicAuthIntegration.id.apply((id) => `integrations/${id}`),
  });
  const apigwExecuteArn = await corestack.getOutputValue("apigwExecuteArn");

  // Give API Gateway permissions to invoke the Lambda
  new aws.lambda.Permission("lambdaPermission", {
    action: "lambda:InvokeFunction",
    principal: "apigateway.amazonaws.com",
    function: lambda,
    sourceArn: route.routeKey.apply((key) => {
      return `${apigwExecuteArn}*/*${key.split(" ")[1]}`;
    }),
  });
}

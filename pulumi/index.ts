import * as awsx from "@pulumi/awsx";
import * as aws from "@pulumi/aws";
import { RedirectProps } from "react-router";
// import * as pulumi from "@pulumi/pulumi";

async function main() {
  // Setup AppGatway logging
  const apigwRole = new aws.iam.Role("cloudwatch", {
    assumeRolePolicy: {
      Version: "2012-10-17",
      Statement: [
        {
          Sid: "",
          Effect: "Allow",
          Principal: {
            Service: "apigateway.amazonaws.com",
          },
          Action: "sts:AssumeRole",
        },
      ],
    },
  });

  const apigwPolicy = new aws.iam.RolePolicy("appgwPolicy", {
    role: apigwRole,
    policy: {
      Version: "2012-10-17",
      Statement: [
        {
          Effect: "Allow",
          Action: [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:DescribeLogGroups",
            "logs:DescribeLogStreams",
            "logs:PutLogEvents",
            "logs:GetLogEvents",
            "logs:FilterLogEvents",
          ],
          Resource: "*",
        },
      ],
    },
  });

  // Onetime setup of appgateway to cloudwatch logging
  const appAccount = new aws.apigateway.Account("cloudwatch", { cloudwatchRoleArn: apigwRole.arn });

  // Define a new GET endpoint from an existing Lambda Function.
  const apigw = new aws.apigatewayv2.Api("appgw", {
    protocolType: "HTTP",
    corsConfiguration: {
      // allowCredentials: true,
      allowOrigins: ["*"],
      allowMethods: ["OPTIONS", "POST"],
      // allowHeaders: ["*"],
      allowHeaders: [
        "Content-Type",
        "Authorization",
        "Accept",
        "Accept-Language",
        "Content-Language",
        "User-Agent",
        "Origin",
      ],
      maxAge: 300,
    },
  });

  const route = new aws.apigatewayv2.Route("cors-route", {
    apiId: apigw.id,
    routeKey: "OPTIONS /{proxy+}",
  });

  const gwlogs = new aws.cloudwatch.LogGroup("appgw-default", {
    retentionInDays: 3,
  });

  const apigwStage = new aws.apigatewayv2.Stage("default", {
    apiId: apigw.id,
    autoDeploy: true,
    name: "$default",
    accessLogSettings: {
      destinationArn: gwlogs.arn,
      format: JSON.stringify({
        requestId: "$context.requestId",
        extendedRequestId: "$context.extendedRequestId",
        ip: "$context.identity.sourceIp",
        caller: "$context.identity.caller",
        user: "$context.identity.user",
        requestTime: "$context.requestTime",
        httpMethod: "$context.httpMethod",
        resourcePath: "$context.resourcePath",
        status: "$context.status",
        protocol: "$context.protocol",
        responseLength: "$context.responseLength",
      }),
    },
  });

  const lambdaS3 = new aws.s3.Bucket("lambda-storage", {
    acl: "private",
    versioning: { enabled: true },
  });

  // TODO - enable this when we're setup in VPC
  let redis: aws.elasticache.Cluster | null = null;
  if (process.env.ISN_NOT_SET) {
    redis = new aws.elasticache.Cluster("redis", {
      engine: "redis",
      nodeType: "cache.t4g.micro",
      numCacheNodes: 1,
    });
  }

  new aws.ssm.Parameter("common/jwt_secret", {
    name: "/common/jwt_secret",
    type: "SecureString",
    value: "oVIPN3MZRIC73IvLPgYYBfuNMkAS3Rzu",
  });

  const entityTopic = new aws.sns.Topic("app-topic");

  new aws.ssm.Parameter("/common/bus_entity_arn", {
    name: "/common/bus_entity_arn",
    type: "String",
    value: entityTopic.arn,
  });
  new aws.ssm.Parameter("/common/url_base", {
    name: "/common/url_base",
    type: "String",
    value: "http://localhost:1234",
  });

  return {
    apigwStageArn: apigwStage.arn,
    apigwExecuteArn: apigw.executionArn,
    entityTopicArn: entityTopic.arn,
    apigwArn: apigw.arn,
    apigwId: apigw.id,
    redisId: redis?.id ?? "",
    redisHost: redis?.cacheNodes[0].address ?? "",
    redisPort: redis?.cacheNodes[0].port ?? "",
    lambdaS3Id: lambdaS3.id,
  };
}

module.exports = main();

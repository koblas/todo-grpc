import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { GoHandler } from "./components/gohandler";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
}

function createApi(
  scope: Construct,
  {
    lambda,
    apigw,
    apiPath,
    targetPath,
  }: {
    lambda: aws.lambdaFunction.LambdaFunction;
    apigw: aws.apigatewayv2Api.Apigatewayv2Api;
    apiPath: string;
    targetPath: string;
  },
) {
  const integration = new aws.apigatewayv2Integration.Apigatewayv2Integration(scope, "integration", {
    apiId: apigw.id,
    integrationType: "AWS_PROXY",
    connectionType: "INTERNET",
    // contentHandlingStrategy: "CONVERT_TO_TEXT",
    integrationMethod: "ANY",
    integrationUri: lambda.arn,
    passthroughBehavior: "WHEN_NO_MATCH",
    payloadFormatVersion: "2.0",

    requestParameters: {
      "overwrite:path": `/${targetPath}/$request.path.proxy`,
    },
  });

  new aws.apigatewayv2Route.Apigatewayv2Route(scope, "route", {
    apiId: apigw.id,
    routeKey: `POST /api/v1/${apiPath}/{proxy+}`,
    target: `integrations/${integration.id}`,
  });
}

export class PublicAuth extends Construct {
  constructor(scope: Construct, id: string, { apigw }: { apigw: aws.apigatewayv2Api.Apigatewayv2Api }) {
    super(scope, id);

    const { lambda } = new GoHandler(this, "publicapi-auth", {
      path: ["publicapi", "auth"],
      apiTrigger: apigw,
      parameters: ["/common/*"],
    });

    createApi(this, {
      lambda,
      apigw,
      apiPath: "auth",
      targetPath: "api.auth.v1.AuthenticationService",
    });
  }
}

export class PublicFile extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { bucket, apigw }: { bucket: aws.s3Bucket.S3Bucket; apigw: aws.apigatewayv2Api.Apigatewayv2Api },
  ) {
    super(scope, id);

    const { lambda } = new GoHandler(this, "publicapi-file", {
      path: ["publicapi", "file"],
      environment: {
        variables: {
          UPLOAD_BUCKET: bucket.bucket,
        },
      },
      apiTrigger: apigw,
      parameters: ["/common/*"],
      s3buckets: [bucket],
    });

    createApi(this, {
      lambda,
      apigw,
      apiPath: "file",
      targetPath: "api.file.v1.FileService",
    });
  }
}

export class PublicTodo extends Construct {
  constructor(scope: Construct, id: string, { apigw }: { apigw: aws.apigatewayv2Api.Apigatewayv2Api }) {
    super(scope, id);

    const { lambda } = new GoHandler(this, "publicapi-todo", {
      path: ["publicapi", "todo"],
      apiTrigger: apigw,
      parameters: ["/common/*"],
    });

    createApi(this, {
      lambda,
      apigw,
      apiPath: "todo",
      targetPath: "api.todo.v1.TodoService",
    });
  }
}

export class PublicUser extends Construct {
  constructor(scope: Construct, id: string, { apigw }: { apigw: aws.apigatewayv2Api.Apigatewayv2Api }) {
    super(scope, id);

    const { lambda } = new GoHandler(this, "publicapi-user", {
      path: ["publicapi", "user"],
      apiTrigger: apigw,
      parameters: ["/common/*"],
    });

    createApi(this, {
      lambda,
      apigw,
      apiPath: "user",
      targetPath: "api.user.v1.UserService",
    });
  }
}

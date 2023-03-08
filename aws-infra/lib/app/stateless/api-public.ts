import * as cdk from "aws-cdk-lib";
import {
  HttpApi,
  PayloadFormatVersion,
  HttpMethod,
  ParameterMapping,
  MappingValue,
} from "@aws-cdk/aws-apigatewayv2-alpha";
import { HttpLambdaIntegration } from "@aws-cdk/aws-apigatewayv2-integrations-alpha";
import { Construct } from "constructs";
import { GoHandler } from "./api-utils";

export class PublicAuth extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const handler = new GoHandler(this, "publicapi-auth", ["publicapi", "auth"], { parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", handler.lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipbv1.auth.AuthenticationService/$request.path.proxy"),
      ),
    });

    apigw.addRoutes({
      path: "/v1/auth/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
    apigw.addRoutes({
      path: "/api/v1/auth/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
  }
}

export class PublicFile extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { apigw, uploadBucket }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi; uploadBucket: cdk.aws_s3.IBucket },
  ) {
    super(scope, id);

    const handler = new GoHandler(this, "publicapi-file", ["publicapi", "file"], {
      environment: {
        // UPLOAD_BUCKET: uploadBucket.bucketName,
      },
      parameters: ["/common/*"],
      s3buckets: [uploadBucket],
    });

    const integration = new HttpLambdaIntegration("integration", handler.lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipbv1.file.FileService/$request.path.proxy"),
      ),
    });

    apigw.addRoutes({
      path: "/v1/file/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
    apigw.addRoutes({
      path: "/api/v1/file/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
  }
}

export class PublicUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const handler = new GoHandler(this, "publicapi-user", ["publicapi", "user"], {
      eventbus,
      parameters: ["/common/*"],
    });

    const integration = new HttpLambdaIntegration("integration", handler.lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipbv1.user.UserService/$request.path.proxy"),
      ),
    });

    apigw.addRoutes({
      path: "/v1/user/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
    apigw.addRoutes({
      path: "/api/v1/user/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
  }
}

export class PublicTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const handler = new GoHandler(this, "publicapi-todo", ["publicapi", "todo"], {
      eventbus,
      parameters: ["/common/*"],
    });

    const integration = new HttpLambdaIntegration("integration", handler.lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipbv1.todo.TodoService/$request.path.proxy"),
      ),
    });

    apigw.addRoutes({
      path: "/v1/todo/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
    apigw.addRoutes({
      path: "/api/v1/todo/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
  }
}

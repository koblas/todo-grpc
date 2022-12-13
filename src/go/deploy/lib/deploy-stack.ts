import * as path from "path";
import * as cdk from "aws-cdk-lib";
import {
  HttpApi,
  CorsHttpMethod,
  PayloadFormatVersion,
  HttpMethod,
  ParameterMapping,
  MappingValue,
  DomainName,
  WebSocketApi,
  WebSocketStage,
} from "@aws-cdk/aws-apigatewayv2-alpha";
import { HttpLambdaIntegration, WebSocketLambdaIntegration } from "@aws-cdk/aws-apigatewayv2-integrations-alpha";
import { LambdaToDynamoDB } from "@aws-solutions-constructs/aws-lambda-dynamodb";
import { LambdaToSns } from "@aws-solutions-constructs/aws-lambda-sns";
import { Construct } from "constructs";
import { SqsToLambda } from "@aws-solutions-constructs/aws-sqs-lambda";
import { SubscriptionFilter } from "aws-cdk-lib/aws-sns";
import { goFunction } from "./utils";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class DeployStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const config = {
      zoneName: "iqvine.com",
      apihostname: "api",
      wshostname: "wsapi",
    };

    const { hostedZone, certificate } = this.getRoute53HostedZone(config.zoneName, [
      `${config.apihostname}.${config.zoneName}`,
      `${config.wshostname}.${config.zoneName}`,
    ]);

    // The code that defines your stack goes here
    const eventbus = new cdk.aws_sns.Topic(this, "eventbus", {});

    const apiDn = new DomainName(this, "apiDn", {
      domainName: `${config.apihostname}.${config.zoneName}`,
      certificate,
    });
    const wsDn = new DomainName(this, "wsDn", {
      domainName: `${config.wshostname}.${config.zoneName}`,
      certificate,
    });

    //
    const apigw = new ApiGw(this, "apigw", {
      dn: apiDn,
      hostedZone: hostedZone,
      hostname: config.apihostname,
    });
    const apiws = new WebsocketHandler(this, "websock", {
      dn: wsDn,
      hostedZone: hostedZone,
      hostname: config.wshostname,
    });

    new cdk.CfnOutput(this, `apigw-export-apiId`, {
      exportName: `apigw-apiId`,
      value: apigw.apigw.apiId,
    });
    new cdk.CfnOutput(this, `apigw-export-endpoint`, {
      exportName: `apigw-endpoint`,
      value: apigw.apigw.apiEndpoint,
    });

    new cdk.CfnOutput(this, `wsapi-export-apiId`, {
      exportName: `wsapi-apiId`,
      value: apiws.wsapi.apiId,
    });
    new cdk.CfnOutput(this, `wsapi-export-endpoint`, {
      exportName: `wsapi-endpoint`,
      value: apiws.wsapi.apiEndpoint,
    });

    //
    new cdk.aws_ssm.StringParameter(this, "bus_entity_arn", {
      tier: cdk.aws_ssm.ParameterTier.STANDARD,
      description: "SNS Topic for events",
      parameterName: "/common/bus_entity_arn",
      stringValue: eventbus.topicArn,
    });
    new cdk.aws_ssm.StringParameter(this, "url_base", {
      tier: cdk.aws_ssm.ParameterTier.STANDARD,
      description: "Base URL for the service",
      parameterName: "/common/url_base",
      stringValue: "http://localhost:1234",
    });
    // new cdk.aws_ssm.StringParameter(this, "jwt_secret", {
    //   tier: cdk.aws_ssm.ParameterTier.STANDARD,
    //   description: "Secret for the JWT Key",
    //   parameterName: "/common/jwt_secret",
    //   stringValue: "xyzzy",
    // });

    // apigw.addRoutes({
    //   methods: [HttpMethod.OPTIONS],
    //   path: "/{proxy+}",
    // });

    new CoreUser(this, "core-user", { eventbus });
    new CoreTodo(this, "core-todo", { eventbus });
    new CoreOauthUser(this, "core-oauth-user", { eventbus });
    // new CoreSendEmail(this, "core-send-email", { eventbus });
    new PublicUser(this, "public-user", { eventbus, apigw: apigw.apigw });
    new PublicAuth(this, "public-auth", { eventbus, apigw: apigw.apigw });
    new PublicTodo(this, "public-todo", { eventbus, apigw: apigw.apigw });
    new CreateWorkers(this, "workers", { eventbus });
    new CoreSendEmailQueue(this, "send-email-queue", { eventbus });
    new WebsocketTodo(this, "websocket-todo", { eventbus, wsstage: apiws.wsstage, wsapi: apiws.wsapi });

    // await sqsWorkers(corestack);
  }

  private getRoute53HostedZone(baseDns: string, domains: string[]) {
    const hostedZone = cdk.aws_route53.HostedZone.fromLookup(this, "HostedZone", {
      domainName: baseDns,
    });

    const [initialDns, ...rest] = domains;

    const certificate = new cdk.aws_certificatemanager.DnsValidatedCertificate(this, "Certificate", {
      hostedZone,
      domainName: initialDns,
      subjectAlternativeNames: [baseDns, ...rest],
      region: this.region,
      validation: cdk.aws_certificatemanager.CertificateValidation.fromDns(hostedZone),
    });

    return { hostedZone, certificate };
  }
}

export class ApiGw extends Construct {
  apigw: HttpApi;

  constructor(
    scope: Construct,
    id: string,
    props: { dn: DomainName; hostedZone: cdk.aws_route53.IHostedZone; hostname: string },
  ) {
    super(scope, id);

    const apigw = new HttpApi(this, "apigw", {
      corsPreflight: {
        allowOrigins: ["*"],
        allowMethods: [CorsHttpMethod.GET, CorsHttpMethod.HEAD, CorsHttpMethod.OPTIONS, CorsHttpMethod.POST],
        allowHeaders: [
          // "Content-Type",
          "Authorization",
          "Content-Type",
          "X-Api-Key",
          // "Accept",
          // "Accept-Language",
          // "Content-Language",
          // "User-Agent",
          // "Origin",
        ],
        maxAge: cdk.Duration.days(10),
      },
      defaultDomainMapping: {
        domainName: props.dn,
      },
    });

    new cdk.aws_apigatewayv2.CfnRoute(this, "OptionsResource", {
      apiId: apigw.apiId,
      routeKey: "OPTIONS /{proxy+}",
    });

    new cdk.aws_route53.ARecord(this, "Alias", {
      zone: props.hostedZone,
      recordName: props.hostname,
      target: cdk.aws_route53.RecordTarget.fromAlias(
        new cdk.aws_route53_targets.ApiGatewayv2DomainProperties(
          props.dn.regionalDomainName,
          props.dn.regionalHostedZoneId,
        ),
      ),
    });

    this.apigw = apigw;
  }
}

export class WebsocketHandler extends Construct {
  public wsapi: WebSocketApi;
  public wsstage: WebSocketStage;

  constructor(
    scope: Construct,
    id: string,
    props: { dn: DomainName; hostedZone: cdk.aws_route53.IHostedZone; hostname: string },
  ) {
    super(scope, id);

    const db = new cdk.aws_dynamodb.Table(this, "db", {
      tableName: "ws-connection",
      billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
      sortKey: { name: "sk", type: cdk.aws_dynamodb.AttributeType.STRING },
      timeToLiveAttribute: "delete_at",
    });

    const lambda = goFunction(this, "publicapi-websocket", ["publicapi", "websocket"], {
      environment: {
        CONN_DB: db.tableName,
      },
    });

    wireLambda(this, lambda, { parameters: ["/common/*"], dynamo: db });

    this.wsapi = new WebSocketApi(this, "wsapi", {
      connectRouteOptions: { integration: new WebSocketLambdaIntegration("connect", lambda) },
      disconnectRouteOptions: { integration: new WebSocketLambdaIntegration("disconnect", lambda) },
      defaultRouteOptions: { integration: new WebSocketLambdaIntegration("default", lambda) },
    });

    this.wsstage = new WebSocketStage(this, "DefaultStage", {
      webSocketApi: this.wsapi,
      stageName: "$default",
      autoDeploy: true,
      domainMapping: { domainName: props.dn },
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
        new cdk.aws_route53_targets.ApiGatewayv2DomainProperties(
          props.dn.regionalDomainName,
          props.dn.regionalHostedZoneId,
        ),
      ),
    });
  }
}

///
// Usage:  `new MockPing(this, "mockPing", { routeKey: "ping", sockapi: YOUR_WEB_SOCKET});`
//
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

export class CoreTodo extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const db = new cdk.aws_dynamodb.Table(this, "db", {
      tableName: "app-todo",
      billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
      sortKey: { name: "sk", type: cdk.aws_dynamodb.AttributeType.STRING },
    });

    const lambda = goFunction(this, "core-todo", ["core", "todo"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"], dynamo: db });
  }
}

export class CoreUser extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const db = new cdk.aws_dynamodb.Table(this, "db", {
      tableName: "app-user",
      billingMode: cdk.aws_dynamodb.BillingMode.PAY_PER_REQUEST,
      partitionKey: { name: "pk", type: cdk.aws_dynamodb.AttributeType.STRING },
    });

    const lambda = goFunction(this, "core-user", ["core", "user"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"], dynamo: db });
  }
}

export class CoreOauthUser extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const lambda = goFunction(this, "core-oauth_user", ["core", "oauth_user"]);

    wireLambda(this, lambda, { parameters: ["/common/*", "/oauth/*"], eventbus });
  }
}

// export class CoreSendEmail extends Construct {
//   constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
//     super(scope, id);

//     const { eventbus } = props;

//     const lambda = new cdk.aws_lambda.Function(this, "lambda", {
//       functionName: "core-send-email",
//       code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
//       runtime: cdk.aws_lambda.Runtime.GO_1_X,
//       handler: "core-send-email",
//     });

//     wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });
//   }
// }

export class CoreSendEmailQueue extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const lambda = goFunction(this, "core-send_email", ["core", "send_email"]);

    new QueueLambda(this, "core-send-email", {
      eventbus,
      queueProps: { queueName: "send-email" },
      existingLambdaObj: lambda,
    });
  }
}

export class PublicAuth extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const lambda = goFunction(this, "publicapi-auth", ["publicapi", "auth"]);

    wireLambda(this, lambda, { parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/api.auth.AuthenticationService/$request.path.proxy"),
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

export class PublicUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const lambda = goFunction(this, "publicapi-user", ["publicapi", "user"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/api.auth.UserService/$request.path.proxy"),
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

export class PublicTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const lambda = goFunction(this, "publicapi-todo", ["publicapi", "todo"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/api.todo.TodoService/$request.path.proxy"),
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

//
export class WebsocketTodo extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { eventbus, wsstage, wsapi }: { eventbus: cdk.aws_sns.Topic; wsstage: WebSocketStage; wsapi: WebSocketApi },
  ) {
    super(scope, id);

    const db = cdk.aws_dynamodb.Table.fromTableName(this, "conns", "ws-connection");

    const lambda = goFunction(this, "websocket-todo", ["websocket", "todo"], {
      environment: {
        CONN_DB: db.tableName,
        WS_ENDPOINT: wsstage.callbackUrl,
      },
    });

    wsstage.grantManagementApiAccess(lambda);

    const queue = new SqsToLambda(this, "sqs", {
      existingLambdaObj: lambda,
      deadLetterQueueProps: {
        queueName: `${id}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${id}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        // SNS cannot deliver to encrypted SQS queues
        encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
      },
    });

    wireLambda(this, lambda, { parameters: ["/common/*"], dynamo: db });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(queue.sqsQueue, {
        rawMessageDelivery: true,
        filterPolicy: {
          "twirp.path": SubscriptionFilter.stringFilter({
            allowlist: ["/twirp/core.eventbus.TodoEventbus/TodoChange"],
          }),
        },
      }),
    );

    //
  }
}

///
//
//
export class CreateWorkers extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    new QueueWorker(this, "password-changed", {
      eventbus,
      env: { SQS_HANDLER: "userSecurity/password_changed" },
      filterPolicy: {
        stream: SubscriptionFilter.stringFilter({
          allowlist: ["event:user_security"],
        }),
        action: SubscriptionFilter.stringFilter({
          allowlist: ["USER_PASSWORD_CHANGE"],
        }),
      },
    });

    new QueueWorker(this, "register-token", {
      eventbus,
      env: { SQS_HANDLER: "userSecurity/register" },
      filterPolicy: {
        stream: SubscriptionFilter.stringFilter({
          allowlist: ["event:user_security"],
        }),
        action: SubscriptionFilter.stringFilter({
          allowlist: ["USER_REGISTER_TOKEN"],
        }),
      },
    });

    new QueueWorker(this, "forgot-request", {
      eventbus,
      env: { SQS_HANDLER: "userSecurity/forgot" },
      filterPolicy: {
        stream: SubscriptionFilter.stringFilter({
          allowlist: ["event:user_security"],
        }),
        action: SubscriptionFilter.stringFilter({
          allowlist: ["USER_FORGOT_REQUEST"],
        }),
      },
    });

    new QueueWorker(this, "user-invite", {
      eventbus,
      env: { SQS_HANDLER: "userSecurity/invite" },
      filterPolicy: {
        stream: SubscriptionFilter.stringFilter({
          allowlist: ["event:user_security"],
        }),
        action: SubscriptionFilter.stringFilter({
          allowlist: ["USER_INVITE_TOKEN"],
        }),
      },
    });
  }
}

export class QueueWorker extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      env,
      filterPolicy,
    }: {
      eventbus: cdk.aws_sns.Topic;
      env: Record<string, string>;
      filterPolicy: NonNullable<cdk.aws_sns_subscriptions.SubscriptionProps["filterPolicy"]>;
    },
  ) {
    super(scope, id);

    const lambda = goFunction(this, `core-workers-${id}`, ["core", "workers"], {
      environment: env,
    });

    const worker = new QueueLambda(this, id, {
      eventbus,
      queueProps: {
        // SNS cannot deliver to encrypted SQS queues
        encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
      },
      existingLambdaObj: lambda,
    });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(worker.queue, {
        rawMessageDelivery: true,
        filterPolicy,
      }),
    );
  }
}

export class QueueLambda extends Construct {
  queue: cdk.aws_sqs.Queue;

  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      lambdaFunctionProps,
      existingLambdaObj,
      queueProps,
    }: {
      eventbus: cdk.aws_sns.Topic;
      queueProps?: cdk.aws_sqs.QueueProps;
      lambdaFunctionProps?: cdk.aws_lambda.FunctionProps;
      existingLambdaObj?: cdk.aws_lambda.Function;
    },
  ) {
    super(scope, id);

    // Connect the queue
    const inst = new SqsToLambda(this, "sqs", {
      ...(lambdaFunctionProps
        ? {
            lambdaFunctionProps: {
              functionName: `worker-${id}`,
              logRetention: cdk.Duration.days(3).toDays(),
              ...lambdaFunctionProps,
            },
          }
        : {}),
      ...(existingLambdaObj ? { existingLambdaObj } : {}),
      deadLetterQueueProps: {
        queueName: `${id}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${id}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        ...queueProps,
      },
    });

    this.queue = inst.sqsQueue;

    // Make sure we can write to SNS
    wireLambda(this, inst.lambdaFunction, { eventbus, parameters: ["/common/*"] });
  }
}

//
//
//
function wireLambda(
  scope: Construct,
  lambda: cdk.aws_lambda.Function,
  {
    eventbus,
    parameters,
    dynamo,
  }: {
    dynamo?: cdk.aws_dynamodb.ITable;
    eventbus?: cdk.aws_sns.Topic;
    parameters?: string[];
  },
) {
  if (parameters?.length) {
    cdk.aws_iam.Grant.addToPrincipal({
      grantee: lambda,
      actions: ["ssm:DescribeParameters", "ssm:GetParameters", "ssm:GetParameter", "ssm:GetParameterHistory"],
      resourceArns: parameters.map((p) =>
        cdk.Stack.of(scope).formatArn({
          service: "ssm",
          resource: `parameter${p}`,
        }),
      ),
    });
  }

  // Invoke and SendMessage are good downstream calls
  cdk.aws_iam.Grant.addToPrincipal({
    grantee: lambda,
    actions: ["lambda:InvokeFunction"],
    resourceArns: [
      cdk.Stack.of(scope).formatArn({
        service: "lambda",
        resource: "function:*",
      }),
    ],
  });

  cdk.aws_iam.Grant.addToPrincipal({
    grantee: lambda,
    actions: ["sqs:GetQueueUrl", "sqs:SendMessage"],
    resourceArns: [
      cdk.Stack.of(scope).formatArn({
        service: "sqs",
        resource: "*",
      }),
    ],
  });

  if (eventbus) {
    new LambdaToSns(scope, "sns-perms", {
      existingLambdaObj: lambda,
      existingTopicObj: eventbus,
    });
  }

  if (dynamo) {
    new LambdaToDynamoDB(scope, "dynamo-perms", {
      existingTableObj: dynamo as cdk.aws_dynamodb.Table,
      tablePermissions: "ReadWrite",
      existingLambdaObj: lambda,
    });
  }
}

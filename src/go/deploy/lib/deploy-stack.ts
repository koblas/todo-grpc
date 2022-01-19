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
} from "@aws-cdk/aws-apigatewayv2-alpha";
import { HttpLambdaIntegration } from "@aws-cdk/aws-apigatewayv2-integrations-alpha";
import { LambdaToDynamoDB } from "@aws-solutions-constructs/aws-lambda-dynamodb";
import { LambdaToSns } from "@aws-solutions-constructs/aws-lambda-sns";
import { Construct } from "constructs";
import { SqsToLambda } from "@aws-solutions-constructs/aws-sqs-lambda";
import { SubscriptionFilter } from "aws-cdk-lib/aws-sns";
import { Grant } from "aws-cdk-lib/aws-iam";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class DeployStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const config = {
      zoneName: "iqvine.com",
      subdomain: "api",
    };

    const { hostedZone, certificate } = this.getRoute53HostedZone(config.zoneName, [
      `${config.subdomain}.${config.zoneName}`,
    ]);

    // The code that defines your stack goes here
    const eventbus = new cdk.aws_sns.Topic(this, "eventbus", {});

    const dn = new DomainName(this, "DN", {
      domainName: `${config.subdomain}.${config.zoneName}`,
      certificate,
    });

    //
    const apigw = new HttpApi(this, "apigw", {
      corsPreflight: {
        allowOrigins: ["*"],
        allowMethods: [CorsHttpMethod.GET, CorsHttpMethod.HEAD, CorsHttpMethod.OPTIONS, CorsHttpMethod.POST],
        allowHeaders: [
          // "Content-Type",
          "X-Api-Key",
          "Authorization",
          "Content-Type",
          // "Accept",
          // "Accept-Language",
          // "Content-Language",
          // "User-Agent",
          // "Origin",
        ],
        maxAge: cdk.Duration.days(10),
      },
      defaultDomainMapping: {
        domainName: dn,
      },
    });

    new cdk.aws_apigatewayv2.CfnRoute(this, "OptionsResource", {
      apiId: apigw.apiId,
      routeKey: "OPTIONS /{proxy+}",
    });

    new cdk.aws_route53.ARecord(this, "Alias", {
      zone: hostedZone,
      recordName: config.subdomain,
      target: cdk.aws_route53.RecordTarget.fromAlias(
        new cdk.aws_route53_targets.ApiGatewayv2DomainProperties(dn.regionalDomainName, dn.regionalHostedZoneId),
      ),
    });

    //
    new cdk.aws_ssm.StringParameter(this, "entity", {
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
    new CoreSendEmail(this, "core-send-email", { eventbus });
    new PublicAuth(this, "public-auth", { eventbus, apigw: apigw });
    new PublicTodo(this, "public-todo", { eventbus, apigw: apigw });
    new CreateWorkers(this, "workers", { eventbus });

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
      region: "us-east-1",
      validation: cdk.aws_certificatemanager.CertificateValidation.fromDns(hostedZone),
    });

    return { hostedZone, certificate };
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

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "core-todo",
    });

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

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      functionName: "core-user",
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "core-user",
    });

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"], dynamo: db });
  }
}

export class CoreOauthUser extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      functionName: "core-oauth-user",
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "core-oauth-user",
    });

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*", "/oauth/*"] });
  }
}

export class CoreSendEmail extends Construct {
  constructor(scope: Construct, id: string, props: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const { eventbus } = props;

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      functionName: "core-send-email",
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "core-send-email",
    });

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });
  }
}

export class PublicAuth extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      functionName: "public-auth",
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "publicapi-auth",
      logRetention: cdk.Duration.days(3).toDays(),
      timeout: cdk.Duration.seconds(10),
    });

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/api.auth.AuthenticationService/$request.path.proxy"),
      ),
    });

    apigw.addRoutes({
      path: "/v1/auth/{proxy+}",
      methods: [HttpMethod.POST],
      integration: integration,
    });
  }
}

export class PublicTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus, apigw }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi }) {
    super(scope, id);

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      functionName: "public-todo",
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "publicapi-todo",
    });

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

    const deadletter = new cdk.aws_sqs.Queue(this, "queue-dql", {
      queueName: `${id}-dlq`,
      retentionPeriod: cdk.Duration.days(7),
    });
    const queue = new cdk.aws_sqs.Queue(this, "queue", {
      queueName: id,
      retentionPeriod: cdk.Duration.days(7),
      visibilityTimeout: cdk.Duration.minutes(5),
      deadLetterQueue: {
        queue: deadletter,
        maxReceiveCount: 10,
      },
    });

    // Doesn't work to subscribe multiple times...
    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(queue, {
        rawMessageDelivery: true,
        filterPolicy,
      }),
    );

    const lambda = new cdk.aws_lambda.Function(this, "lambda", {
      functionName: `worker-${id}`,
      code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, "..", "..", "build")),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: "core-workers",
      environment: env,
    });

    // Connect the queue
    new SqsToLambda(this, "sqs", {
      existingLambdaObj: lambda,
      existingQueueObj: queue,
    });

    // Make sure we can write to SNS
    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });
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
    dynamo?: cdk.aws_dynamodb.Table;
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

  if (eventbus) {
    new LambdaToSns(scope, "sns-perms", {
      existingLambdaObj: lambda,
      existingTopicObj: eventbus,
    });
  }

  if (dynamo) {
    new LambdaToDynamoDB(scope, "dynamo-perms", {
      existingTableObj: dynamo,
      tablePermissions: "ReadWrite",
      existingLambdaObj: lambda,
    });
  }
}

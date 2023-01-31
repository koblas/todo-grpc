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
import { Construct } from "constructs";
import { goFunction, wireLambda, QueueLambda } from "./utils";
import { WebsocketBroadcast, WebsocketTodo, WebsocketUser } from "./websocket";
import { CreateWorkersFile, CreateWorkersUser } from "./workers";
import { CertificateStack } from "./certificate";

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
    const coreFile = new FileStorage(this, "core-file", { eventbus, hostedZone, hostname: "files" });
    new CoreOauthUser(this, "core-oauth-user", { eventbus });
    // new CoreSendEmail(this, "core-send-email", { eventbus });
    new PublicAuth(this, "public-auth", { eventbus, apigw: apigw.apigw });
    new PublicFile(this, "public-file", { eventbus, apigw: apigw.apigw, publicBucket: coreFile.uploadBucket });
    new PublicTodo(this, "public-todo", { eventbus, apigw: apigw.apigw });
    new PublicUser(this, "public-user", { eventbus, apigw: apigw.apigw });
    new CreateWorkersUser(this, "workers_user", { eventbus });
    new CreateWorkersFile(this, "workers_file", { eventbus, coreFile });
    new CoreSendEmailQueue(this, "send-email-queue", { eventbus });
    new WebsocketTodo(this, "websocket-todo", { eventbus });
    new WebsocketUser(this, "websocket-user", { eventbus });
    new WebsocketBroadcast(this, "websocket-broadcast", { eventbus, wsstage: apiws.wsstage, wsapi: apiws.wsapi });

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
      region: this.region, // must be in US-East-1 for Cloudfront
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

export class FileStorage extends Construct {
  // Where uploaded files are staged
  uploadBucket: cdk.aws_s3.Bucket;
  // Non-public files
  privateBucket: cdk.aws_s3.Bucket;
  // Where public (e.g. web accessible) files are staged
  // this is fronted by CloudFront
  publicBucket: cdk.aws_s3.Bucket;

  constructor(
    scope: Construct,
    id: string,
    props: {
      eventbus: cdk.aws_sns.Topic;
      hostname: string;
      hostedZone: cdk.aws_route53.IHostedZone;
    },
  ) {
    super(scope, id);

    this.publicBucket = new cdk.aws_s3.Bucket(this, "public_bucket", {
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
    });

    this.privateBucket = new cdk.aws_s3.Bucket(this, "private_bucket", {
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
    });

    this.uploadBucket = new cdk.aws_s3.Bucket(this, "upload_bucket", {
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
      lifecycleRules: [
        {
          expiration: cdk.Duration.days(1),
        },
      ],
      cors: [
        {
          allowedMethods: [cdk.aws_s3.HttpMethods.GET, cdk.aws_s3.HttpMethods.PUT],
          allowedOrigins: ["*"],
          allowedHeaders: ["*"],
        },
      ],
    });

    const domainName = `${props.hostname}.${props.hostedZone.zoneName}`;

    const certificate = new CertificateStack(scope, "certificate", {
      hostedZone: props.hostedZone,
      domainName: `files.${props.hostedZone.zoneName}`,
      region: "us-east-1",
    });

    const accessIdentity = new cdk.aws_cloudfront.OriginAccessIdentity(this, "OriginAccessIdentity", {
      comment: `${this.uploadBucket.bucketName}-access-identity`,
    });

    const distribution = new cdk.aws_cloudfront.Distribution(this, "filesource", {
      domainNames: [domainName],
      certificate: certificate.certificate,
      priceClass: cdk.aws_cloudfront.PriceClass.PRICE_CLASS_100,
      defaultRootObject: "",

      defaultBehavior: {
        origin: new cdk.aws_cloudfront_origins.S3Origin(this.publicBucket, {
          // originAccessIdentity: accessIdentity,
        }),
        // originRequestPolicy: cdk.aws_cloudfront.OriginRequestPolicy.CORS_S3_ORIGIN,
        allowedMethods: cdk.aws_cloudfront.AllowedMethods.ALLOW_GET_HEAD_OPTIONS,
        viewerProtocolPolicy: cdk.aws_cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        // responseHeadersPolicy: responseHeaderPolicy,
        functionAssociations: [],

        // responseHeadersPolicy: cdk.aws_cloudfront.ResponseHeadersPolicy.CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS,
      },

      // additionalBehaviors: additionalBehaviors,

      // We need to redirect all unknown routes back to index.html for angular routing to work
      errorResponses: [],

      minimumProtocolVersion: cdk.aws_cloudfront.SecurityPolicyProtocol.TLS_V1_2_2021,
      sslSupportMethod: cdk.aws_cloudfront.SSLMethod.SNI,
    });

    new cdk.aws_route53.ARecord(this, "Alias", {
      zone: props.hostedZone,
      recordName: domainName,
      target: cdk.aws_route53.RecordTarget.fromAlias(new cdk.aws_route53_targets.CloudFrontTarget(distribution)),
    });

    const { eventbus } = props;

    // const lambda = goFunction(this, "core-file", ["core", "file"], {
    //   environment: {
    //     BUCKET: this.uploadBucket.bucketName,
    //     BUCKET_DOMAIN: domainName,
    //     BUCKET_ALIAS: `${props.hostname}.${props.hostedZone.zoneName}`,
    //   },
    // });

    // wireLambda(this, lambda, { eventbus, parameters: ["/common/*"], s3buckets: [this.uploadBucket] });

    new TriggerS3(this, "trigger-s3", { eventbus, bucket: this.uploadBucket });
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
        MappingValue.custom("/twirp/apipb.auth.AuthenticationService/$request.path.proxy"),
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
    { eventbus, apigw, publicBucket }: { eventbus: cdk.aws_sns.Topic; apigw: HttpApi; publicBucket: cdk.aws_s3.Bucket },
  ) {
    super(scope, id);

    const lambda = goFunction(this, "publicapi-file", ["publicapi", "file"], {
      environment: {
        UPLOAD_BUCKET: publicBucket.bucketName,
      },
    });

    wireLambda(this, lambda, { parameters: ["/common/*"], s3buckets: [publicBucket] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipb.file.FileService/$request.path.proxy"),
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

    const lambda = goFunction(this, "publicapi-user", ["publicapi", "user"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipb.user.UserService/$request.path.proxy"),
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

    const lambda = goFunction(this, "publicapi-todo", ["publicapi", "todo"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });

    const integration = new HttpLambdaIntegration("integration", lambda, {
      payloadFormatVersion: PayloadFormatVersion.VERSION_2_0,
      parameterMapping: new ParameterMapping().overwritePath(
        MappingValue.custom("/twirp/apipb.todo.TodoService/$request.path.proxy"),
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

///
//
//

export class TriggerS3 extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { eventbus, bucket }: { eventbus: cdk.aws_sns.Topic; bucket: cdk.aws_s3.Bucket },
  ) {
    super(scope, id);

    const lambda = goFunction(this, "trigger-s3", ["trigger", "s3"]);

    wireLambda(this, lambda, { eventbus, parameters: ["/common/*"] });

    const notificaiton = new cdk.aws_s3_notifications.LambdaDestination(lambda);
    notificaiton.bind(scope, bucket);
    bucket.addEventNotification(cdk.aws_s3.EventType.OBJECT_CREATED, notificaiton);
  }
}

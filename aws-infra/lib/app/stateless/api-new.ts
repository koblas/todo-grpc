import { HttpApi, WebSocketApi, WebSocketStage } from "@aws-cdk/aws-apigatewayv2-alpha";
import * as cdk from "aws-cdk-lib";
import {
  OriginRequestCookieBehavior,
  OriginRequestHeaderBehavior,
  OriginRequestPolicy,
  OriginRequestQueryStringBehavior,
} from "aws-cdk-lib/aws-cloudfront";
import { Construct } from "constructs";
import { BackendStack } from "./api-backend";
import { WebsocketHandler } from "./apigw-websocket";

export interface Props {
  distribution: cdk.aws_cloudfront.Distribution;
  apigw: HttpApi;
  hostedZone: cdk.aws_route53.IHostedZone;
  // uploadBucket: cdk.aws_s3.Bucket;
  // publicBucket: cdk.aws_s3.Bucket;
  // privateBucket: cdk.aws_s3.Bucket;
  // uploadBucketArn: string;
  // publicBucketArn: string;
  // privateBucketArn: string;
  // uploadBucketName: string;
  // publicBucketName: string;
  // privateBucketName: string;
  // appHostname: string;
  // filesHostname: string;
}

export class ApiStack extends Construct {
  public wsapi: WebSocketApi;
  public wsstage: WebSocketStage;
  public wsdb: cdk.aws_dynamodb.Table;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    const { distribution } = props;

    const apiPolicy = new OriginRequestPolicy(this, "api-policy", {
      originRequestPolicyName: "apiPolicy",
      queryStringBehavior: OriginRequestQueryStringBehavior.all(),
      cookieBehavior: OriginRequestCookieBehavior.none(),
      headerBehavior: OriginRequestHeaderBehavior.none(),
    });

    const apiDomain = cdk.Fn.select(1, cdk.Fn.split("://", props.apigw.apiEndpoint));
    distribution.addBehavior("/api/*", new cdk.aws_cloudfront_origins.HttpOrigin(apiDomain), {
      originRequestPolicy: apiPolicy,
      allowedMethods: cdk.aws_cloudfront.AllowedMethods.ALLOW_ALL,
      cachePolicy: new cdk.aws_cloudfront.CachePolicy(this, "apiCachePolicy", {
        headerBehavior: cdk.aws_cloudfront.CacheHeaderBehavior.allowList("Authorization"),
      }),
      viewerProtocolPolicy: cdk.aws_cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
      responseHeadersPolicy: cdk.aws_cloudfront.ResponseHeadersPolicy.CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS,
    });

    const rewriteFunction = new cdk.aws_cloudfront.Function(this, "Function", {
      code: cdk.aws_cloudfront.FunctionCode.fromInline(`
          var re = /^\\/(api|wsapi)(\\/|$)/gm

          function handler(event) {
            var request = event.request;

            request.uri = request.uri.replace(re, "/");

            return request;
        }
      `),
    });

    //
    // Websocket
    //
    const ws = new WebsocketHandler(this, "Websocket", {
      hostedZone: props.hostedZone,
      hostname: "wsapi",
    });
    this.wsstage = ws.wsstage;
    this.wsapi = ws.wsapi;
    this.wsdb = ws.db;

    const wsPolicy = new OriginRequestPolicy(this, "ws-policy", {
      originRequestPolicyName: "webSocketPolicy",
      queryStringBehavior: OriginRequestQueryStringBehavior.all(),
      cookieBehavior: OriginRequestCookieBehavior.none(),
      headerBehavior: OriginRequestHeaderBehavior.allowList(
        "Sec-WebSocket-Key",
        "Sec-WebSocket-Version",
        "Sec-WebSocket-Protocol",
        "Sec-WebSocket-Accept",
      ),
    });

    const wsDomain = cdk.Fn.select(1, cdk.Fn.split("://", ws.wsapi.apiEndpoint));

    distribution.addBehavior("/wsapi", new cdk.aws_cloudfront_origins.HttpOrigin(wsDomain), {
      originRequestPolicy: wsPolicy,
      allowedMethods: cdk.aws_cloudfront.AllowedMethods.ALLOW_ALL,
      cachePolicy: cdk.aws_cloudfront.CachePolicy.CACHING_DISABLED,
      viewerProtocolPolicy: cdk.aws_cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
      responseHeadersPolicy: cdk.aws_cloudfront.ResponseHeadersPolicy.CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS,
      functionAssociations: [
        {
          function: rewriteFunction,
          eventType: cdk.aws_cloudfront.FunctionEventType.VIEWER_REQUEST,
        },
      ],
    });
  }
}

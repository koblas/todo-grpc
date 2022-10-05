import * as cdk from "aws-cdk-lib";
import { Aws } from "aws-cdk-lib";
import {
  CloudFrontAllowedMethods,
  OriginRequestCookieBehavior,
  OriginRequestHeaderBehavior,
  OriginRequestPolicy,
  OriginRequestQueryStringBehavior,
} from "aws-cdk-lib/aws-cloudfront";
import { HostingStack } from "./hostingStack";

export class Site extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const apiEndpoint = cdk.Fn.importValue("apigw-endpoint");
    const wsEndpoint = cdk.Fn.importValue("wsapi-endpoint");

    const apiPolicy = new OriginRequestPolicy(this, "api-policy", {
      originRequestPolicyName: "apiPolicy",
      queryStringBehavior: OriginRequestQueryStringBehavior.all(),
      cookieBehavior: OriginRequestCookieBehavior.none(),
      headerBehavior: OriginRequestHeaderBehavior.none(),
    });
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

    const apiDomain = cdk.Fn.select(1, cdk.Fn.split("://", apiEndpoint));
    const wsDomain = cdk.Fn.select(1, cdk.Fn.split("://", wsEndpoint));

    // const rewriteFunction = new cdk.aws_cloudfront.Function(this, "Function", {
    //   code: cdk.aws_cloudfront.FunctionCode.fromInline(`
    //       function handler(event) {
    //         var request = event.request;

    //         if (request.uri.startsWith('/api')) {
    //             request.uri = request.uri.substring(4)
    //         }
    //         if (request.uri.startsWith('/wsapi')) {
    //             request.uri = request.uri.substring(6)
    //         }

    //         return request;
    //     }
    //   `),
    // });

    new HostingStack(this, "app-spa", props ?? {}).createSiteFromHostedZone({
      indexDoc: "index.html",
      websiteFolder: "../dist",
      zoneName: "iqvine.com",
      subdomain: "app",

      replications: ["us-west-2"],
      additional: {
        "/api/*": {
          origin: new cdk.aws_cloudfront_origins.HttpOrigin(apiDomain),
          originRequestPolicy: apiPolicy,
          allowedMethods: cdk.aws_cloudfront.AllowedMethods.ALLOW_ALL,
          cachePolicy: new cdk.aws_cloudfront.CachePolicy(this, "apiCachePolicy", {
            headerBehavior: cdk.aws_cloudfront.CacheHeaderBehavior.allowList("Authorization"),
          }),
          viewerProtocolPolicy: cdk.aws_cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
          responseHeadersPolicy: cdk.aws_cloudfront.ResponseHeadersPolicy.CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS,
          // functionAssociations: [
          //   {
          //     function: rewriteFunction,
          //     eventType: cdk.aws_cloudfront.FunctionEventType.VIEWER_REQUEST,
          //   },
          // ],
        },
        "/wsapi": {
          origin: new cdk.aws_cloudfront_origins.HttpOrigin(wsDomain),
          originRequestPolicy: wsPolicy,
          allowedMethods: cdk.aws_cloudfront.AllowedMethods.ALLOW_ALL,
          cachePolicy: cdk.aws_cloudfront.CachePolicy.CACHING_DISABLED,
          viewerProtocolPolicy: cdk.aws_cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
          responseHeadersPolicy: cdk.aws_cloudfront.ResponseHeadersPolicy.CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS,
          // functionAssociations: [
          //   {
          //     function: rewriteFunction,
          //     eventType: cdk.aws_cloudfront.FunctionEventType.VIEWER_REQUEST,
          //   },
          // ],
        },
      },
    });
  }
}

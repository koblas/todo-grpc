import * as path from "path";
import * as fs from "fs";
import * as mime from "mime-types";
import * as aws from "@cdktf/provider-aws";
import { DataAwsIamPolicyDocument } from "@cdktf/provider-aws/lib/data-aws-iam-policy-document";
import { Construct } from "constructs";
import { CertificateDomain } from "./components/certificate";
import { StateS3Bucket } from "./components/s3bucket";
import { Fn, TerraformAsset } from "cdktf";
import { WebsocketHandler } from "./gw-websocket";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
  zone: { zoneId: string };
  domainName: string;
  wsHostname: string;
  appHostname: string;
  filesDomainName: string;
}

const API_ORIGIN = "apigw";
const SPA_ORIGIN = "static_web";
const WS_ORIGIN = "websocket";

export class Frontend extends Construct {
  public appDomainName: string;
  public distribution: aws.cloudfrontDistribution.CloudfrontDistribution;
  public wsdb: aws.dynamodbTable.DynamodbTable;
  public wsapi: aws.apigatewayv2Api.Apigatewayv2Api;
  public wsstage: aws.apigatewayv2Stage.Apigatewayv2Stage;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    //
    const { bucket } = new StateS3Bucket(this, "public", {
      bucketPrefix: "webapp",
      forceDestroy: true,
    });

    this.appDomainName = `${props.appHostname}.${props.domainName}`;

    const { cert } = new CertificateDomain(this, "fileCert", {
      domainName: this.appDomainName,
      zoneId: props.zone.zoneId,
      region: "us-east-1",
    });

    const responseHeaderPolicy = new aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy(
      this,
      "policy",
      {
        name: "app-response-policy",
        comment: "Security headers response header policy",
        securityHeadersConfig: {
          contentSecurityPolicy: {
            override: true,
            contentSecurityPolicy: [
              "default-src 'self'",
              "script-src 'self' 'unsafe-inline' 'unsafe-eval'",
              "style-src 'self' 'unsafe-inline'",
              "img-src * 'unsafe-inline' data:",
              // `connect-src 'self' wss: ws: ${domainNames.map((d) => `https://*.files.${zoneName}`).join(" ")}`,
              `connect-src 'self' wss: ws: https://${props.filesDomainName} https://*.s3.us-west-2.amazonaws.com`,
            ].join(";"),
          },
          strictTransportSecurity: {
            override: true,
            accessControlMaxAgeSec: 2 * 365 * 24 * 60 * 60,
            includeSubdomains: true,
            preload: true,
          },
          contentTypeOptions: {
            override: true,
          },
          referrerPolicy: {
            override: true,
            referrerPolicy: "strict-origin-when-cross-origin",
          },
          xssProtection: {
            override: true,
            protection: true,
            modeBlock: true,
          },
          frameOptions: {
            override: true,
            frameOption: "DENY",
          },
        },
      },
    );

    const oac = new aws.cloudfrontOriginAccessControl.CloudfrontOriginAccessControl(this, "oac", {
      name: "s3-access",
      originAccessControlOriginType: "s3",
      signingBehavior: "always",
      signingProtocol: "sigv4",
    });

    const ws = new WebsocketHandler(this, "ws", props);
    this.wsapi = ws.wsapi;
    this.wsstage = ws.wsstage;
    this.wsdb = ws.wsdb;

    // const secret = new randprovider.uuid.Uuid(this, "secret", {});

    this.distribution = new aws.cloudfrontDistribution.CloudfrontDistribution(this, "spa", {
      enabled: true,
      isIpv6Enabled: true,
      priceClass: "PriceClass_100",
      origin: [
        {
          originId: SPA_ORIGIN,
          domainName: bucket.bucketRegionalDomainName,
          originAccessControlId: oac.id,
        },
        {
          domainName: Fn.element(Fn.split("://", props.apigw.apiEndpoint), 1),
          originId: API_ORIGIN,
          customOriginConfig: {
            httpPort: 80,
            httpsPort: 443,
            originProtocolPolicy: "https-only",
            originSslProtocols: ["TLSv1.2"],
          },
        },
        {
          domainName: Fn.element(Fn.split("://", this.wsapi.apiEndpoint), 1),
          originId: WS_ORIGIN,
          customOriginConfig: {
            httpPort: 80,
            httpsPort: 443,
            originProtocolPolicy: "https-only",
            originSslProtocols: ["TLSv1.2"],
          },
        },
      ],
      defaultCacheBehavior: {
        minTtl: 0,
        defaultTtl: 1 * 60 * 60, // 1 hour
        maxTtl: 1 * 24 * 60 * 60, // 1 day
        cachedMethods: ["GET", "HEAD"],
        allowedMethods: ["HEAD", "GET", "OPTIONS"],
        viewerProtocolPolicy: "redirect-to-https",
        targetOriginId: SPA_ORIGIN,
        responseHeadersPolicyId: responseHeaderPolicy.id,
        forwardedValues: {
          cookies: {
            forward: "none",
          },
          queryString: true,
          headers: [],
        },
        compress: true,
      },
      aliases: [this.appDomainName],
      viewerCertificate: {
        acmCertificateArn: cert.arn,
        sslSupportMethod: "sni-only",
        minimumProtocolVersion: "TLSv1.2_2021",
      },
      restrictions: {
        geoRestriction: {
          restrictionType: "none",
        },
      },
      orderedCacheBehavior: [this.createApiBehavior(props)],
      defaultRootObject: "index.html",
      customErrorResponse: [
        {
          errorCachingMinTtl: 60 * 60, // 1 hour
          errorCode: 404,
          responseCode: 404,
          responsePagePath: "/index.html",
        },
      ],
    });

    new aws.route53Record.Route53Record(this, "appAlias", {
      name: this.appDomainName,
      type: "A",
      zoneId: props.zone.zoneId,
      alias: {
        name: this.distribution.domainName,
        zoneId: this.distribution.hostedZoneId,
        evaluateTargetHealth: true,
      },
    });

    // Add Assets
    bucketDeployment(this, "spa", {
      source: path.join(__dirname, "../../ts-client-react/dist"),
      bucket,
      distribution: this.distribution,
      invalidations: [],
    });

    const s3policy = new DataAwsIamPolicyDocument(this, "s3policy", {
      statement: [
        {
          actions: ["s3:GetObject", "s3:ListBucket"],
          resources: [`${bucket.arn}/*`, bucket.arn],
          principals: [
            {
              type: "Service",
              identifiers: ["cloudfront.amazonaws.com"],
            },
          ],
          condition: [
            {
              test: "StringEquals",
              variable: "AWS:SourceArn",
              values: [this.distribution.arn],
            },
          ],
        },
      ],
    });

    new aws.s3BucketPolicy.S3BucketPolicy(this, "origin", {
      bucket: bucket.id,
      policy: s3policy.json,
    });

    // TODO Wire up Websocket
  }

  createWsBehavior(_: Props): aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehavior {
    const policy = new aws.cloudfrontOriginRequestPolicy.CloudfrontOriginRequestPolicy(this, "ws-policy", {
      name: "webSocketPolicy",
      queryStringsConfig: { queryStringBehavior: "all" },
      cookiesConfig: { cookieBehavior: "none" },
      headersConfig: {
        headerBehavior: "whitelist",
        headers: {
          items: ["Sec-WebSocket-Key", "Sec-WebSocket-Version", "Sec-WebSocket-Protocol", "Sec-WebSocket-Accept"],
        },
      },
    });
    const rhead = new aws.dataAwsCloudfrontResponseHeadersPolicy.DataAwsCloudfrontResponseHeadersPolicy(
      this,
      "ws-response",
      {
        name: "Managed-CORS-and-SecurityHeadersPolicy",
      },
    );
    const cache = new aws.cloudfrontCachePolicy.CloudfrontCachePolicy(this, "ws-cache", {
      name: "api-cache",
      defaultTtl: 0,
      maxTtl: 0,
      minTtl: 0,
      parametersInCacheKeyAndForwardedToOrigin: {
        cookiesConfig: {
          cookieBehavior: "none",
        },
        headersConfig: {
          headerBehavior: "none",
        },
        queryStringsConfig: {
          queryStringBehavior: "none",
        },
      },
    });

    return {
      pathPattern: "/wsapi",
      targetOriginId: WS_ORIGIN,
      originRequestPolicyId: policy.id,
      allowedMethods: ["GET", "HEAD", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"],
      cachedMethods: ["GET", "HEAD", "OPTIONS"],
      viewerProtocolPolicy: "redirect-to-https",
      compress: true,
      minTtl: 0,
      maxTtl: 0,
      defaultTtl: 0,
      responseHeadersPolicyId: rhead.id,
      cachePolicyId: cache.id,
    };
  }

  createApiBehavior(_: Props): aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehavior {
    const policy = new aws.dataAwsCloudfrontOriginRequestPolicy.DataAwsCloudfrontOriginRequestPolicy(
      this,
      "api-request",
      {
        name: "Managed-AllViewerExceptHostHeader",
      },
    );

    const cache = new aws.cloudfrontCachePolicy.CloudfrontCachePolicy(this, "api-cache", {
      name: "api-cache",
      defaultTtl: 0,
      maxTtl: 0,
      minTtl: 0,
      parametersInCacheKeyAndForwardedToOrigin: {
        cookiesConfig: {
          cookieBehavior: "none",
        },
        headersConfig: {
          headerBehavior: "none",
        },
        queryStringsConfig: {
          queryStringBehavior: "none",
        },
      },
    });
    const rhead = new aws.dataAwsCloudfrontResponseHeadersPolicy.DataAwsCloudfrontResponseHeadersPolicy(
      this,
      "api-response",
      {
        name: "Managed-CORS-and-SecurityHeadersPolicy",
      },
    );

    return {
      pathPattern: "/api/*",
      targetOriginId: API_ORIGIN,
      originRequestPolicyId: policy.id,
      allowedMethods: ["GET", "HEAD", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"],
      cachedMethods: ["GET", "HEAD", "OPTIONS"],
      viewerProtocolPolicy: "redirect-to-https",
      compress: true,
      minTtl: 0,
      maxTtl: 0,
      defaultTtl: 0,
      responseHeadersPolicyId: rhead.id,
      cachePolicyId: cache.id,
    };
  }

  createWebsocket(_: Props) {}
}

function bucketDeployment(
  scope: Construct,
  id: string,
  params: {
    source: string;
    bucket: aws.s3Bucket.S3Bucket;
    distribution: aws.cloudfrontDistribution.CloudfrontDistribution;
    invalidations: string[];
  },
) {
  for (const item of getFiles(params.source)) {
    const nm = item.replace(params.source, "").replace(/[^A-Za-z_-]/, "_");

    const asset = new TerraformAsset(scope, `${id}-${nm}-assets`, {
      path: item,
      // type: AssetType.ARCHIVE, // if left empty it infers directory and file
    });

    new aws.s3Object.S3Object(scope, `${id}-${nm}-objects`, {
      bucket: params.bucket.bucket,
      key: asset.fileName,
      source: asset.path, // returns a posix path
      contentType: mime.lookup(asset.fileName) || "application/octet-stream",
    });
  }
}

function* getFiles(item: string): IterableIterator<string> {
  const s = fs.statSync(item);

  if (!s.isDirectory()) {
    yield item;
  }

  const dirents = fs.readdirSync(item, { withFileTypes: true });

  for (const dirent of dirents) {
    const res = path.resolve(item, dirent.name);
    if (dirent.isDirectory()) {
      yield* getFiles(res);
    } else {
      yield res;
    }
  }
}

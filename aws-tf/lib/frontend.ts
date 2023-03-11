import * as path from "path";
import * as crypto from "crypto";
import * as fs from "fs";
import * as mime from "mime-types";
import * as aws from "@cdktf/provider-aws";
import * as nullp from "@cdktf/provider-null";
import { Construct } from "constructs";
import { CertificateDomain } from "./components/certificate";
import { StateS3Bucket } from "./components/s3bucket";
import { Fn, TerraformAsset } from "cdktf";
import { WebsocketConfig, WebsocketHandler } from "./gw-websocket";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
  zone: { zoneId: string };
  domainName: string;
  wsHostname: string;
  appHostname: string;
  filesDomainName: string;
  logsBucket: aws.s3Bucket.S3Bucket;
}

const API_ORIGIN = "apigw";
const SPA_ORIGIN = "static_web";
const WS_ORIGIN = "websocket";

export class Frontend extends Construct {
  public appDomainName: string;
  public distribution: aws.cloudfrontDistribution.CloudfrontDistribution;
  public wsconf: WebsocketConfig;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    const account = new aws.dataAwsCallerIdentity.DataAwsCallerIdentity(this, "ident", {});

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
        // customHeadersConfig: {
        //   items: [
        //     {
        //       header: "Cache-control",
        //       override: true,
        //       value: "max-age=60",
        //     },
        //   ],
        // },
      },
    );

    const oac = new aws.cloudfrontOriginAccessControl.CloudfrontOriginAccessControl(this, "oac", {
      name: "s3-spa-access",
      originAccessControlOriginType: "s3",
      signingBehavior: "always",
      signingProtocol: "sigv4",
    });

    const ws = new WebsocketHandler(this, "ws", props);
    this.wsconf = ws.wsconf;

    //
    const logDoc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "logdoc", {
      statement: [
        {
          effect: "Allow",
          principals: [
            {
              type: "AWS",
              identifiers: [`arn:aws:iam::${account.accountId}:root`],
            },
          ],
          actions: ["s3:GetBucketAcl", "s3:PutBucketAcl"],
          resources: [props.logsBucket.arn],
        },
      ],
    });
    const logPolicy = new aws.s3BucketPolicy.S3BucketPolicy(this, "logpolicy", {
      bucket: props.logsBucket.id,
      policy: logDoc.json,
    });

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
          domainName: Fn.element(Fn.split("://", this.wsconf.wsapi.apiEndpoint), 1),
          // originPath: `/${this.wsstage.name}`,
          // domainName: ws.wsdomainName,
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
      orderedCacheBehavior: [this.createApiBehavior(props), this.createWsBehavior(this.wsconf.wsstage, props)],
      defaultRootObject: "index.html",
      customErrorResponse: [
        {
          errorCode: 404,
          responseCode: 200,
          responsePagePath: "/index.html",
          errorCachingMinTtl: 1 * 60 * 60, // 5 minutes (same as default)
        },
      ],
      loggingConfig: {
        bucket: props.logsBucket.bucketDomainName,
        prefix: "cloudfront/app/",
      },

      dependsOn: [logPolicy],
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
    const { hash, s3objects } = bucketDeployment(this, "spa", {
      source: path.join(__dirname, "../../ts-client-react/dist"),
      bucket,
      distribution: this.distribution,
      invalidations: [],
    });

    new nullp.resource.Resource(this, "invalidate", {
      dependsOn: s3objects,

      triggers: {
        hash,
      },

      provisioners: [
        {
          type: "local-exec",
          command: `aws cloudfront create-invalidation --distribution-id ${this.distribution.id} --paths /index.html`,
        },
      ],
    });

    const s3policy = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "s3policy", {
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

  createWsBehavior(
    wsstage: aws.apigatewayv2Stage.Apigatewayv2Stage,
    _: Props,
  ): aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehavior {
    const policy = new aws.cloudfrontOriginRequestPolicy.CloudfrontOriginRequestPolicy(this, "ws-policy", {
      name: "webSocketPolicy",
      queryStringsConfig: { queryStringBehavior: "none" },
      cookiesConfig: { cookieBehavior: "none" },
      headersConfig: {
        headerBehavior: "whitelist",
        headers: {
          items: [
            "Sec-WebSocket-Key",
            "Sec-WebSocket-Version",
            "Sec-WebSocket-Protocol",
            "Sec-WebSocket-Accept",
            "Sec-WebSocket-Extensions",
          ],
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
    const cpolicy = new aws.dataAwsCloudfrontCachePolicy.DataAwsCloudfrontCachePolicy(this, "ws-cache", {
      name: "Managed-CachingDisabled",
    });
    // const cache = new aws.cloudfrontCachePolicy.CloudfrontCachePolicy(this, "ws-cache", {
    //   name: "ws-cache",
    //   defaultTtl: 0,
    //   maxTtl: 0,
    //   minTtl: 0,
    //   parametersInCacheKeyAndForwardedToOrigin: {
    //     cookiesConfig: {
    //       cookieBehavior: "none",
    //     },
    //     headersConfig: {
    //       headerBehavior: "none",
    //     },
    //     queryStringsConfig: {
    //       queryStringBehavior: "none",
    //     },
    //   },
    // });

    const viewerFunc = new aws.cloudfrontFunction.CloudfrontFunction(this, "viewer", {
      name: "remove-wsapi",
      runtime: "cloudfront-js-1.0",
      code: `function handler(event) {
        var request = event.request;
        request.uri = "/${wsstage.name}";
        return request;
      }`,
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
      cachePolicyId: cpolicy.id,
      functionAssociation: [
        {
          eventType: "viewer-request",
          functionArn: viewerFunc.arn,
        },
      ],
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
  const s3objects: aws.s3Object.S3Object[] = [];
  const assets = [];

  for (const path of getFiles(params.source)) {
    const nm = path.replace(params.source, "").replace(/[^A-Za-z_-]/, "_");

    assets.push(
      new TerraformAsset(scope, `${id}-${nm}-assets`, {
        path,
        // type: AssetType.ARCHIVE, // if left empty it infers directory and file
      }),
    );
  }

  const hash = crypto.createHash("md5");

  const ONE_HOUR = 1 * 60 * 60;
  const ONE_YEAR = 365 * 24 * 60 * 60;

  assets.map((asset) => {
    const s3o = new aws.s3Object.S3Object(scope, `${id}-${asset.assetHash}-objects`, {
      bucket: params.bucket.bucket,
      sourceHash: asset.assetHash,
      key: asset.fileName,
      source: asset.path, // returns a posix path
      contentType: mime.lookup(asset.fileName) || "application/octet-stream",
      cacheControl: `public, max-age=${asset.fileName === "index.html" ? ONE_HOUR : ONE_YEAR}`,
    });

    hash.update(asset.assetHash);
    s3objects.push(s3o);
  });

  return { hash: hash.digest("hex"), s3objects };
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

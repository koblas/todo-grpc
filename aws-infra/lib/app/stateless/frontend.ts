import * as path from "path";
import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { ApiStack } from "./api-new";
import { HttpApi, WebSocketApi, WebSocketStage } from "@aws-cdk/aws-apigatewayv2-alpha";

export type Props = {
  apigw: HttpApi;
  hostedZone: cdk.aws_route53.IHostedZone;
  domainName: string;
  appHostname: string;
};

/*
 ** Useful resource -
 **  https://idanlupinsky.com/blog/static-site-deployment-using-aws-cloudfront-and-the-cdk/
 */

export interface StackProps extends cdk.StackProps {
  readonly indexDoc: string;
  readonly errorDoc?: string;
  readonly websiteFolder: string;
  readonly certificateARN?: string;
  readonly cfBehaviors?: cdk.aws_cloudfront.Behavior[];
  readonly cfAliases?: string[];
  readonly exportWebsiteUrlOutput?: boolean;
  readonly exportWebsiteUrlName?: string;
  readonly blockPublicAccess?: cdk.aws_s3.BlockPublicAccess;
  readonly sslMethod?: cdk.aws_cloudfront.SSLMethod;
  readonly securityPolicy?: cdk.aws_cloudfront.SecurityPolicyProtocol;
  readonly role?: cdk.aws_iam.Role;
}

export interface HostedProps {
  readonly apigw: HttpApi;
  readonly indexDoc: string;
  readonly errorDoc?: string;
  readonly cfBehaviors?: cdk.aws_cloudfront.Behavior[];
  readonly websiteFolder: string;
  readonly zoneName: string;
  readonly subdomain?: string;
  readonly role?: cdk.aws_iam.Role;
  readonly replications?: string[];

  readonly additional?: Record<string, cdk.aws_cloudfront.BehaviorOptions>;

  readonly ipFilter?: boolean;
  readonly ipList?: string[];
  readonly hostedZone: cdk.aws_route53.IHostedZone;
}

export class FrontendStack extends Construct {
  public distribution: cdk.aws_cloudfront.Distribution;
  public wsdb: cdk.aws_dynamodb.Table;
  public wsapi: WebSocketApi;
  public wsstage: WebSocketStage;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    this.createSiteFromHostedZone({
      hostedZone: props.hostedZone,
      indexDoc: "index.html",
      websiteFolder: path.join(__dirname, "..", "..", "..", "..", "ts-client-react", "dist"),
      zoneName: props.domainName,
      subdomain: props.appHostname,

      replications: ["us-west-2"],
      apigw: props.apigw,
    });
  }

  private getRoute53HostedZone(hostedZone: cdk.aws_route53.IHostedZone, baseDns: string, domains: string[]) {
    const [initialDns, ...rest] = domains;

    const certificate = new cdk.aws_certificatemanager.DnsValidatedCertificate(this, "Certificate", {
      hostedZone,
      domainName: initialDns,
      subjectAlternativeNames: [baseDns, ...rest],
      region: "us-east-1",
      validation: cdk.aws_certificatemanager.CertificateValidation.fromDns(hostedZone),
    });

    return { certificate };
  }

  getS3Bucket(config: HostedProps, isForCloudFront: boolean = true) {
    const bucket = new cdk.aws_s3.Bucket(this, "WebsiteBucket", {
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
    });

    if (config.ipFilter === true && isForCloudFront === false) {
      if (typeof config.ipList === "undefined") {
        throw new Error("When IP Filter is true then the IP List is required");
      }

      const bucketPolicy = new cdk.aws_iam.PolicyStatement();
      bucketPolicy.addAnyPrincipal();
      bucketPolicy.addActions("s3:GetObject");
      bucketPolicy.addResources(`${bucket.bucketArn}/*`);
      bucketPolicy.addCondition("IpAddress", {
        "aws:SourceIp": config.ipList,
      });

      bucket.addToResourcePolicy(bucketPolicy);
    }

    //The below "reinforces" the IAM Role's attached policy, it's not required but it allows for customers using permission boundaries to write into the bucket.
    if (config.role) {
      bucket.addToResourcePolicy(
        new cdk.aws_iam.PolicyStatement({
          actions: ["s3:GetObject*", "s3:GetBucket*", "s3:List*", "s3:DeleteObject*", "s3:PutObject*", "s3:Abort*"],
          effect: cdk.aws_iam.Effect.ALLOW,
          resources: [bucket.arnForObjects("*"), bucket.bucketArn],
          conditions: {
            StringEquals: {
              "aws:PrincipalArn": config.role.roleArn,
            },
          },
          principals: [new cdk.aws_iam.AnyPrincipal()],
        }),
      );
    }

    return bucket;
  }

  /**
   * Helper method to provide configuration for cloudfront
   */
  private getCFConfig(
    websiteBucket: cdk.aws_s3.Bucket,
    config: HostedProps,
    cert: cdk.aws_certificatemanager.DnsValidatedCertificate,
    accessIdentity: cdk.aws_cloudfront.OriginAccessIdentity,
    zoneName: string,
    domainNames: string[],
    additionalBehaviors: Record<string, cdk.aws_cloudfront.BehaviorOptions> | undefined,
  ): cdk.aws_cloudfront.DistributionProps {
    // certificate: cdk.aws_cloudfront.ViewerCertificate.fromAcmCertificate(cert, {
    //   sslMethod: cdk.aws_cloudfront.SSLMethod.SNI,
    //   securityPolicy: cdk.aws_cloudfront.SecurityPolicyProtocol.TLS_V1_2_2021,
    //   aliases: domainNames,
    // }),
    const responseHeaderPolicy = new cdk.aws_cloudfront.ResponseHeadersPolicy(
      this,
      "SecurityHeadersResponseHeaderPolicy",
      {
        comment: "Security headers response header policy",
        securityHeadersBehavior: {
          contentSecurityPolicy: {
            override: true,
            contentSecurityPolicy: [
              "default-src 'self'",
              "script-src 'self' 'unsafe-inline' 'unsafe-eval'",
              "style-src 'self' 'unsafe-inline'",
              "img-src * 'unsafe-inline' data:",
              // `connect-src 'self' wss: ws: ${domainNames.map((d) => `https://*.files.${zoneName}`).join(" ")}`,
              `connect-src 'self' wss: ws: https://files.${zoneName} https://*.s3.us-west-2.amazonaws.com`,
            ].join(";"),
          },
          strictTransportSecurity: {
            override: true,
            accessControlMaxAge: cdk.Duration.days(2 * 365),
            includeSubdomains: true,
            preload: true,
          },
          contentTypeOptions: {
            override: true,
          },
          referrerPolicy: {
            override: true,
            referrerPolicy: cdk.aws_cloudfront.HeadersReferrerPolicy.STRICT_ORIGIN_WHEN_CROSS_ORIGIN,
          },
          xssProtection: {
            override: true,
            protection: true,
            modeBlock: true,
          },
          frameOptions: {
            override: true,
            frameOption: cdk.aws_cloudfront.HeadersFrameOption.DENY,
          },
        },
      },
    );

    const rewriteFunction = new cdk.aws_cloudfront.Function(this, "Function", {
      code: cdk.aws_cloudfront.FunctionCode.fromInline(`
          function handler(event) {
            var request = event.request;
        
            if (request.uri.endsWith('/')) {
                request.uri += 'index.html';
            }
        
            return request;
        }
      `),
    });

    const cfConfig: cdk.aws_cloudfront.DistributionProps = {
      domainNames: domainNames,
      certificate: cert,
      priceClass: cdk.aws_cloudfront.PriceClass.PRICE_CLASS_100,
      defaultRootObject: "",

      defaultBehavior: {
        origin: new cdk.aws_cloudfront_origins.S3Origin(websiteBucket, {
          originAccessIdentity: accessIdentity,
        }),
        // originRequestPolicy: cdk.aws_cloudfront.OriginRequestPolicy.CORS_S3_ORIGIN,
        allowedMethods: cdk.aws_cloudfront.AllowedMethods.ALLOW_GET_HEAD_OPTIONS,
        viewerProtocolPolicy: cdk.aws_cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        responseHeadersPolicy: responseHeaderPolicy,
        functionAssociations: [
          {
            function: rewriteFunction,
            eventType: cdk.aws_cloudfront.FunctionEventType.VIEWER_REQUEST,
          },
        ],

        // responseHeadersPolicy: cdk.aws_cloudfront.ResponseHeadersPolicy.CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS,
      },

      additionalBehaviors: additionalBehaviors,

      // We need to redirect all unknown routes back to index.html for angular routing to work
      errorResponses: [
        {
          httpStatus: 403,
          responsePagePath: config.errorDoc ? `/${config.errorDoc}` : `/${config.indexDoc}`,
          responseHttpStatus: 200,
        },
        {
          httpStatus: 404,
          responsePagePath: config.errorDoc ? `/${config.errorDoc}` : `/${config.indexDoc}`,
          responseHttpStatus: 200,
        },
      ],

      minimumProtocolVersion: cdk.aws_cloudfront.SecurityPolicyProtocol.TLS_V1_2_2021,
      sslSupportMethod: cdk.aws_cloudfront.SSLMethod.SNI,
    };

    // if (typeof config.certificateARN !== "undefined" && typeof config.cfAliases !== "undefined") {
    //   cfConfig.aliasConfiguration = {
    //     acmCertRef: config.certificateARN,
    //     names: config.cfAliases,
    //   };
    // }
    // if (typeof config.sslMethod !== "undefined") {
    //   cfConfig.aliasConfiguration.sslMethod = config.sslMethod;
    // }

    // if (typeof config.securityPolicy !== "undefined") {
    //   cfConfig.aliasConfiguration.securityPolicy = config.securityPolicy;
    // }

    // if (typeof config.zoneName !== "undefined" && typeof cert !== "undefined") {
    //   cfConfig.viewerCertificate = cdk.aws_cloudfront.ViewerCertificate.fromAcmCertificate(cert, {
    //     aliases: [config.subdomain ? `${config.subdomain}.${config.zoneName}` : config.zoneName],
    //   });
    // }

    return cfConfig;
  }

  public createSiteFromHostedZone(props: HostedProps) {
    // const websiteBucket = this.getS3Bucket(config, true);
    const websiteBucket = this.getS3Bucket(props, false);

    // const zone = cdk.aws_route53.HostedZone.fromLookup(this, "HostedZone", { domainName: config.zoneName });
    const domainName = props.subdomain ? `${props.subdomain}.${props.zoneName}` : props.zoneName;
    const { certificate } = this.getRoute53HostedZone(props.hostedZone, props.zoneName, [domainName]);
    // const cert = new cdk.aws_certificatemanager.DnsValidatedCertificate(this, "Certificate", {
    //   hostedZone: zone,
    //   domainName,
    //   region: "us-east-1",
    // });

    const accessIdentity = new cdk.aws_cloudfront.OriginAccessIdentity(this, "OriginAccessIdentity", {
      comment: `${websiteBucket.bucketName}-access-identity`,
    });

    websiteBucket.addToResourcePolicy(
      new cdk.aws_iam.PolicyStatement({
        actions: ["s3:GetObject"],
        resources: [websiteBucket.arnForObjects("*")],
        principals: [
          new cdk.aws_iam.CanonicalUserPrincipal(accessIdentity.cloudFrontOriginAccessIdentityS3CanonicalUserId),
        ],
      }),
    );

    this.distribution = new cdk.aws_cloudfront.Distribution(
      this,
      "website",
      this.getCFConfig(
        websiteBucket,
        props,
        certificate,
        accessIdentity,
        props.zoneName,
        [domainName],
        props.additional,
      ),
    );

    new cdk.aws_s3_deployment.BucketDeployment(this, "BucketDeployment", {
      sources: [cdk.aws_s3_deployment.Source.asset(props.websiteFolder)],
      destinationBucket: websiteBucket,
      // Invalidate the cache for / and index.html when we deploy so that cloudfront serves latest site
      distribution: this.distribution,
      role: props.role,
      distributionPaths: ["/", `/${props.indexDoc}`],
    });

    new cdk.aws_route53.ARecord(this, "Alias", {
      zone: props.hostedZone,
      recordName: domainName,
      target: cdk.aws_route53.RecordTarget.fromAlias(new cdk.aws_route53_targets.CloudFrontTarget(this.distribution)),
    });

    if (!props.subdomain) {
      new cdk.aws_route53_patterns.HttpsRedirect(this, "Redirect", {
        zone: props.hostedZone,
        recordNames: [`www.${props.zoneName}`],
        targetDomain: props.zoneName,
      });
    }

    const apiws = new ApiStack(this, "ws", {
      hostedZone: props.hostedZone,
      distribution: this.distribution,
      apigw: props.apigw,
    });

    this.wsapi = apiws.wsapi;
    this.wsstage = apiws.wsstage;
    this.wsdb = apiws.wsdb;

    return { websiteBucket };
  }
}

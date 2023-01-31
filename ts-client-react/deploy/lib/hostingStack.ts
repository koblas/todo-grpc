import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";

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

export interface HostedZoneProps {
  readonly indexDoc: string;
  readonly errorDoc?: string;
  readonly cfBehaviors?: cdk.aws_cloudfront.Behavior[];
  readonly websiteFolder: string;
  readonly zoneName: string;
  readonly subdomain?: string;
  readonly role?: cdk.aws_iam.Role;
  readonly replications?: string[];
}

export interface GlobalProps extends cdk.StackProps {
  readonly encryptBucket?: boolean;
  readonly ipFilter?: boolean;
  readonly ipList?: string[];
  readonly role?: cdk.aws_iam.Role;
}

export class HostingStack extends Construct {
  globalProps: GlobalProps;

  constructor(scope: Construct, id: string, props: GlobalProps) {
    super(scope, id);

    const { encryptBucket = false, ipFilter = false } = props;

    this.globalProps = { encryptBucket, ipFilter };
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

  getS3Bucket(config: StackProps, isForCloudFront: boolean = true) {
    /*
    const bucket = new cdk.aws_s3.Bucket(this, "WebsiteBucket", {
      websiteIndexDocument: config.indexDoc,
      websiteErrorDocument: config.errorDoc,
      publicReadAccess: true,
      ...(this.globalProps.encryptBucket ? { encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED } : {}),
      ...(this.globalProps.ipFilter || isForCloudFront || !!config.blockPublicAccess
        ? {
            publicReadAccess: false,
            blockPublicAccess: config.blockPublicAccess,
          }
        : {}),
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });
    */
    const bucket = new cdk.aws_s3.Bucket(this, "WebsiteBucket", {
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
    });

    if (this.globalProps.ipFilter === true && isForCloudFront === false) {
      if (typeof this.globalProps.ipList === "undefined") {
        throw new Error("When IP Filter is true then the IP List is required");
      }

      const bucketPolicy = new cdk.aws_iam.PolicyStatement();
      bucketPolicy.addAnyPrincipal();
      bucketPolicy.addActions("s3:GetObject");
      bucketPolicy.addResources(`${bucket.bucketArn}/*`);
      bucketPolicy.addCondition("IpAddress", {
        "aws:SourceIp": this.globalProps.ipList,
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
    config: HostedZoneProps,
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

  public createSiteFromHostedZone(
    config: HostedZoneProps & { readonly additional?: Record<string, cdk.aws_cloudfront.BehaviorOptions> },
  ) {
    // const websiteBucket = this.getS3Bucket(config, true);
    const websiteBucket = this.getS3Bucket(config, false);

    // const zone = cdk.aws_route53.HostedZone.fromLookup(this, "HostedZone", { domainName: config.zoneName });
    const domainName = config.subdomain ? `${config.subdomain}.${config.zoneName}` : config.zoneName;
    const { hostedZone, certificate } = this.getRoute53HostedZone(config.zoneName, [domainName]);
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

    const distribution = new cdk.aws_cloudfront.Distribution(
      this,
      "cloudfrontDistribution",
      this.getCFConfig(
        websiteBucket,
        config,
        certificate,
        accessIdentity,
        config.zoneName,
        [domainName],
        config.additional,
      ),
    );

    new cdk.aws_s3_deployment.BucketDeployment(this, "BucketDeployment", {
      sources: [cdk.aws_s3_deployment.Source.asset(config.websiteFolder)],
      destinationBucket: websiteBucket,
      // Invalidate the cache for / and index.html when we deploy so that cloudfront serves latest site
      distribution,
      role: config.role,
      distributionPaths: ["/", `/${config.indexDoc}`],
    });

    new cdk.aws_route53.ARecord(this, "Alias", {
      zone: hostedZone,
      recordName: domainName,
      target: cdk.aws_route53.RecordTarget.fromAlias(new cdk.aws_route53_targets.CloudFrontTarget(distribution)),
    });

    if (!config.subdomain) {
      new cdk.aws_route53_patterns.HttpsRedirect(this, "Redirect", {
        zone: hostedZone,
        recordNames: [`www.${config.zoneName}`],
        targetDomain: config.zoneName,
      });
    }

    return { websiteBucket, distribution };
  }
}

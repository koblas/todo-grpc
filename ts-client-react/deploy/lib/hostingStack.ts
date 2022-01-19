import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";

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
    const bucket = new cdk.aws_s3.Bucket(this, "WebsiteBucket", {
      // bucketName:
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
    config: any,
    accessIdentity: cdk.aws_cloudfront.OriginAccessIdentity,
    cert: cdk.aws_certificatemanager.DnsValidatedCertificate,
    domainNames: string[],
  ) {
    const cfConfig: cdk.aws_cloudfront.CloudFrontWebDistributionProps = {
      priceClass: cdk.aws_cloudfront.PriceClass.PRICE_CLASS_100,

      originConfigs: [
        {
          s3OriginSource: {
            s3BucketSource: websiteBucket,
            originAccessIdentity: accessIdentity,
          },
          behaviors: config.cfBehaviors ? config.cfBehaviors : [{ isDefaultBehavior: true, compress: true }],
        },
      ],

      // We need to redirect all unknown routes back to index.html for angular routing to work
      errorConfigurations: [
        {
          errorCode: 403,
          responsePagePath: config.errorDoc ? `/${config.errorDoc}` : `/${config.indexDoc}`,
          responseCode: 200,
        },
        {
          errorCode: 404,
          responsePagePath: config.errorDoc ? `/${config.errorDoc}` : `/${config.indexDoc}`,
          responseCode: 200,
        },
      ],

      viewerCertificate: cdk.aws_cloudfront.ViewerCertificate.fromAcmCertificate(cert, {
        sslMethod: cdk.aws_cloudfront.SSLMethod.SNI,
        securityPolicy: cdk.aws_cloudfront.SecurityPolicyProtocol.TLS_V1_2_2021,
        aliases: domainNames,
      }),
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

  public createSiteFromHostedZone(config: HostedZoneProps) {
    const websiteBucket = this.getS3Bucket(config, true);

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

    const distribution = new cdk.aws_cloudfront.CloudFrontWebDistribution(
      this,
      "cloudfrontDistribution",
      this.getCFConfig(websiteBucket, config, accessIdentity, certificate, [domainName]),
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

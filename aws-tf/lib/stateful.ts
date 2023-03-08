import * as aws from "@cdktf/provider-aws";
import { Construct } from "constructs";
import { CertificateDomain } from "./components/certificate";
import { StateS3Bucket } from "./components/s3bucket";

export interface Props {
  apiHostname: string;
  filesHostname: string;
  domainName: string;
  publicBucketName: string;
  privateBucketName: string;
  uploadBucketName: string;
}

export class Stateful extends Construct {
  public publicBucket: aws.s3Bucket.S3Bucket;
  public privateBucket: aws.s3Bucket.S3Bucket;
  public uploadBucket: aws.s3Bucket.S3Bucket;
  public apigw: aws.apigatewayv2Api.Apigatewayv2Api;
  public zone: aws.dataAwsRoute53Zone.DataAwsRoute53Zone;
  public apiDomainName: string;
  public fileDomainName: string;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    // For certificates
    this.zone = new aws.dataAwsRoute53Zone.DataAwsRoute53Zone(this, "zone", {
      name: props.domainName,
    });

    // Buckets
    this.publicBucket = new StateS3Bucket(this, "public", {
      bucketPrefix: props.publicBucketName,
    }).bucket;
    this.privateBucket = new StateS3Bucket(this, "private", {
      bucketPrefix: props.privateBucketName,
    }).bucket;
    this.uploadBucket = new StateS3Bucket(this, "upload", {
      bucketPrefix: props.uploadBucketName,
      expiresInDays: 1,
      enableCors: true,
    }).bucket;

    // Connect up File storage CDN
    this.fileDomainName = `${props.filesHostname}.${props.domainName}`;

    const { cert } = new CertificateDomain(this, "fileCert", {
      domainName: this.fileDomainName,
      zoneId: this.zone.zoneId,
      region: "us-east-1",
    });

    const fileOriginId = "fileOrigin";

    const distribution = new aws.cloudfrontDistribution.CloudfrontDistribution(this, "files", {
      enabled: true,
      isIpv6Enabled: true,
      priceClass: "PriceClass_100",
      origin: [
        {
          domainName: this.publicBucket.bucketRegionalDomainName,
          originId: fileOriginId,
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
        defaultTtl: 24 * 60 * 60, // 1 day
        maxTtl: 7 * 24 * 60 * 60, // 1 week
        cachedMethods: ["GET", "HEAD"],
        allowedMethods: ["GET", "HEAD"],
        viewerProtocolPolicy: "redirect-to-https",
        targetOriginId: fileOriginId,
        forwardedValues: {
          cookies: {
            forward: "none",
          },
          queryString: false,
          headers: [],
        },
      },
      aliases: [this.fileDomainName],
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
    });

    new aws.route53Record.Route53Record(this, "fileAlias", {
      name: this.fileDomainName,
      type: "A",
      zoneId: this.zone.zoneId,
      alias: {
        name: distribution.domainName,
        zoneId: distribution.hostedZoneId,
        evaluateTargetHealth: true,
      },
    });

    // Create the base API GW
    this.apiDomainName = `${props.apiHostname}.${props.domainName}`;

    this.apigw = new aws.apigatewayv2Api.Apigatewayv2Api(this, "api", {
      name: "api",
      protocolType: "HTTP",
      corsConfiguration: {
        allowOrigins: ["*"],
        allowMethods: ["GET", "HEAD", "OPTIONS", "POST"],
        allowHeaders: ["Authorization", "Content-Type", "X-Api-Key"],
      },
    });

    const apiCert = new CertificateDomain(this, "apicert", {
      domainName: this.apiDomainName,
      zoneId: this.zone.zoneId,
    });

    const dn = new aws.apigatewayv2DomainName.Apigatewayv2DomainName(this, "apidn", {
      domainName: this.apiDomainName,
      domainNameConfiguration: {
        certificateArn: apiCert.cert.arn,
        endpointType: "REGIONAL",
        securityPolicy: "TLS_1_2",
      },
    });

    const stage = new aws.apigatewayv2Stage.Apigatewayv2Stage(this, "stage", {
      apiId: this.apigw.id,
      name: "$default",
      autoDeploy: true,

      lifecycle: {
        ignoreChanges: ["deployment_id"],
      },
    });
    new aws.apigatewayv2ApiMapping.Apigatewayv2ApiMapping(this, "mapping", {
      apiId: this.apigw.id,
      stage: stage.id,
      domainName: dn.id,
    });

    new aws.route53Record.Route53Record(this, "apialias", {
      name: props.apiHostname,
      type: "A",
      zoneId: this.zone.zoneId,
      alias: {
        name: dn.domainNameConfiguration.targetDomainName,
        zoneId: dn.domainNameConfiguration.hostedZoneId,
        evaluateTargetHealth: true,
      },
    });
  }
}

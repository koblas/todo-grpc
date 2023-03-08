import * as cdk from "aws-cdk-lib";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import * as s3 from "aws-cdk-lib/aws-s3";

import { Construct } from "constructs";
import { CertificateStack } from "../utils/certificate";
import { HttpApi, CorsHttpMethod, DomainName } from "@aws-cdk/aws-apigatewayv2-alpha";

export interface StatefulStackProps extends cdk.StackProps {
  appHostname: string;
  filesHostname: string;
  domainName: string;
  publicBucketName: string;
  privateBucketName: string;
  uploadBucketName: string;
}

export class StatefulStack extends Construct {
  public readonly publicBucket: s3.Bucket;
  public readonly privateBucket: s3.Bucket;
  public readonly uploadBucket: s3.Bucket;
  public readonly table: dynamodb.Table;
  public readonly hostedZone: cdk.aws_route53.IHostedZone;
  public readonly apigw: HttpApi;

  constructor(scope: Construct, id: string, props: StatefulStackProps) {
    super(scope, id);

    this.hostedZone = cdk.aws_route53.HostedZone.fromLookup(this, "HostedZone", {
      domainName: props.domainName,
    });

    this.publicBucket = new cdk.aws_s3.Bucket(this, "public_bucket", {
      bucketName: props.publicBucketName,
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
    });

    this.privateBucket = new cdk.aws_s3.Bucket(this, "private_bucket", {
      bucketName: props.privateBucketName,
      publicReadAccess: false,
      blockPublicAccess: cdk.aws_s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      accessControl: cdk.aws_s3.BucketAccessControl.PRIVATE,
      objectOwnership: cdk.aws_s3.ObjectOwnership.BUCKET_OWNER_ENFORCED,
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
    });

    this.uploadBucket = new cdk.aws_s3.Bucket(this, "upload_bucket", {
      bucketName: props.uploadBucketName,
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

    const publicHostname = `${props.appHostname}.${props.domainName}`;
    const fileHostname = `${props.filesHostname}.${props.domainName}`;

    const { certificate } = new CertificateStack(this, "certificate", {
      hostedZone: this.hostedZone,
      domainName: fileHostname,
      // alternativeNames: [fileHostname, publicHostname],
      region: "us-east-1",
    });

    const distribution = new cdk.aws_cloudfront.Distribution(this, "filesource", {
      domainNames: [fileHostname],
      certificate,
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
      zone: this.hostedZone,
      recordName: fileHostname,
      target: cdk.aws_route53.RecordTarget.fromAlias(new cdk.aws_route53_targets.CloudFrontTarget(distribution)),
    });

    // // create the s3 bucket for invoices
    // this.bucket = new s3.Bucket(this, 'Bucket', {
    //   bucketName: props.bucketName, // this is passed through per env from config
    // });

    // // create the dynamodb table
    // this.table = new dynamodb.Table(this, 'Table', {
    //   billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
    //   encryption: dynamodb.TableEncryption.AWS_MANAGED,
    //   pointInTimeRecovery: false,
    //   contributorInsightsEnabled: true,
    //   removalPolicy: RemovalPolicy.DESTROY,
    //   partitionKey: {
    //     name: 'id',
    //     type: dynamodb.AttributeType.STRING,
    //   },
    // });
    const { certificate: certificateApi } = new CertificateStack(this, "apiCert", {
      hostedZone: this.hostedZone,
      domainName: publicHostname,
      // alternativeNames: [publicHostname],
    });

    const apiDn = new DomainName(this, "apiDn", {
      domainName: publicHostname,
      certificate: certificateApi,
    });

    const gw = new ApiGw(this, "apigw", {
      dn: apiDn,
      hostedZone: this.hostedZone,
      hostname: `api-${publicHostname}`,
    });

    this.apigw = gw.apigw;

    // Export
    new cdk.CfnOutput(this, "bucket-public", {
      exportName: "bucket-public-arn",
      value: this.publicBucket.bucketArn,
    });
    new cdk.CfnOutput(this, "bucket-private", {
      exportName: "bucket-private-arn",
      value: this.privateBucket.bucketArn,
    });
    new cdk.CfnOutput(this, "bucket-upload", {
      exportName: "bucket-upload-arn",
      value: this.uploadBucket.bucketArn,
    });
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

    // new cdk.aws_route53.ARecord(this, "Alias", {
    //   zone: props.hostedZone,
    //   recordName: props.hostname,
    //   target: cdk.aws_route53.RecordTarget.fromAlias(
    //     new cdk.aws_route53_targets.ApiGatewayv2DomainProperties(
    //       props.dn.regionalDomainName,
    //       props.dn.regionalHostedZoneId,
    //     ),
    //   ),
    // });

    this.apigw = apigw;
  }
}

import * as cdk from "aws-cdk-lib";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import * as s3 from "aws-cdk-lib/aws-s3";

import { Construct } from "constructs";
import { CertificateStack } from "../utils/certificate";

export interface PublicFilesStackProps extends cdk.StackProps {
  hostname: string;
  domainName: string;
  bucket: cdk.aws_s3.Bucket;
  hostedZone: cdk.aws_route53.IHostedZone;
}

export class _PublicFile extends cdk.Stack {
  constructor(scope: Construct, id: string, props: PublicFilesStackProps) {
    super(scope, id, props);

    const hostname = `${props.hostname}.${props.domainName}`;

    const certificate = new CertificateStack(this, "certificate", {
      hostedZone: props.hostedZone,
      domainName: hostname,
      region: "us-east-1",
    });

    const distribution = new cdk.aws_cloudfront.Distribution(this, "filesource", {
      domainNames: [hostname],
      certificate: certificate.certificate,
      priceClass: cdk.aws_cloudfront.PriceClass.PRICE_CLASS_100,
      defaultRootObject: "",

      defaultBehavior: {
        origin: new cdk.aws_cloudfront_origins.S3Origin(props.bucket, {
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
      recordName: hostname,
      target: cdk.aws_route53.RecordTarget.fromAlias(new cdk.aws_route53_targets.CloudFrontTarget(distribution)),
    });
  }
}

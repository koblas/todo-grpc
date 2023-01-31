import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";

export interface CertificateProps extends cdk.StackProps {
  hostedZone: cdk.aws_route53.IHostedZone;
  domainName: string;
  alternativeNames?: [string];
  region?: string;
}

export class CertificateStack extends Construct {
  certificate: cdk.aws_certificatemanager.Certificate;

  constructor(scope: Construct, id: string, props: CertificateProps) {
    super(scope, id);

    this.certificate = new cdk.aws_certificatemanager.DnsValidatedCertificate(this, `${id}_Certificate`, {
      hostedZone: props.hostedZone,
      domainName: props.domainName,
      subjectAlternativeNames: props.alternativeNames,
      validation: cdk.aws_certificatemanager.CertificateValidation.fromDns(props.hostedZone),
      ...(props.region ? { region: props.region } : {}), // must be in US-East-1 for Cloudfront
    });

    // const apiDn = new DomainName(this, "apiDn", {
    //   domainName: `${config.apihostname}.${config.zoneName}`,
    //   certificate,
    // });
  }
}

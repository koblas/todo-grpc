import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";

export interface Props {
  domainName: string;
  zoneId: string;
  region?: string;
}

export class CertificateDomain extends Construct {
  public cert: aws.acmCertificate.AcmCertificate;
  public record: aws.route53Record.Route53Record;

  static providers: Record<string, aws.provider.AwsProvider> = {};

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    let regionProvider: aws.provider.AwsProvider | undefined;

    if (props.region) {
      regionProvider = CertificateDomain.providers?.[props.region];
      if (!regionProvider) {
        regionProvider = new aws.provider.AwsProvider(this, "eastProvider", {
          region: props.region,
          alias: "route53",
        });
        if (CertificateDomain.providers === undefined) {
          CertificateDomain.providers = {};
        }
        CertificateDomain.providers[props.region] = regionProvider;
      }
    }

    const cert = new aws.acmCertificate.AcmCertificate(this, "cert", {
      domainName: props.domainName,
      validationMethod: "DNS",
      provider: regionProvider,
    });

    this.record = new aws.route53Record.Route53Record(this, "record", {
      name: cert.domainValidationOptions.get(0).resourceRecordName,
      type: cert.domainValidationOptions.get(0).resourceRecordType,
      records: [cert.domainValidationOptions.get(0).resourceRecordValue],
      zoneId: props.zoneId,
      ttl: 60,
      allowOverwrite: true,
    });

    new aws.acmCertificateValidation.AcmCertificateValidation(this, "validation", {
      certificateArn: cert.arn,
      validationRecordFqdns: [this.record.fqdn],
      provider: regionProvider,
    });

    this.cert = cert;
  }
}

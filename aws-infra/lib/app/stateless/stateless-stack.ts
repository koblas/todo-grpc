import * as cdk from "aws-cdk-lib";
import { FrontendStack } from "./frontend";
import { ApiStack } from "./api-new";

import { Construct } from "constructs";
import { HttpApi } from "@aws-cdk/aws-apigatewayv2-alpha";

export interface StatelessStackProps extends cdk.StackProps {
  domainName: string;
  uploadBucket: cdk.aws_s3.Bucket;
  publicBucket: cdk.aws_s3.Bucket;
  privateBucket: cdk.aws_s3.Bucket;
  hostedZone: cdk.aws_route53.IHostedZone;
  apigw: HttpApi;
  appHostname: string;
  filesHostname: string;
}

export class StatelessStack extends Construct {
  public readonly frontend: cdk.Stack;
  public readonly api: cdk.Stack;

  constructor(scope: Construct, id: string, props: StatelessStackProps) {
    super(scope, id);

    // const frontend = new FrontendStack(this, "SPAFrontend", {
    //   hostedZone: props.hostedZone,
    //   domainName: props.domainName,
    //   appHostname: props.appHostname,
    // });

    this.frontend = frontend;

    // this.api = new ApiStack(this, "API", {
    //   distribution: frontend.distribution,
    //   apigw: props.apigw,
    //   hostedZone: props.hostedZone,
    //   uploadBucket: props.uploadBucket,
    //   publicBucket: props.publicBucket,
    //   privateBucket: props.privateBucket,
    //   appHostname: props.appHostname,
    //   filesHostname: props.filesHostname,
    // });
  }
}

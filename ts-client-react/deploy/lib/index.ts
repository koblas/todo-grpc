import * as cdk from "aws-cdk-lib";
import { HostingStack } from "./hostingStack";

export class HostingStackBase extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new HostingStack(scope, "HostingStack", {
      env: {
        account: process.env.CDK_DEFAULT_ACCOUNT,
        region: process.env.CDK_DEFAULT_REGION,
      },
      bucketName: "some-bucket-name",
      url: "some-url.com",
    });
  }
}

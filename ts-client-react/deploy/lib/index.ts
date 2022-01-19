import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { HostingStack } from "./hostingStack";

export class HostingStackBase extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new HostingStack(scope, "HostingStack", {
      // bucketName: "some-bucket-name",
      // url: "some-url.com",
    });
  }
}

import * as cdk from "aws-cdk-lib";
import { HostingStack } from "./hostingStack";

export class Site extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new HostingStack(this, "app-spa", props ?? {}).createSiteFromHostedZone({
      indexDoc: "index.html",
      websiteFolder: "../dist",
      zoneName: "iqvine.com",
      subdomain: "app",

      replications: ["us-west-2"],
    });
  }
}

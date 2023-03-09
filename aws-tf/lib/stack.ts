import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
import * as aws from "@cdktf/provider-aws";
import * as randprovider from "@cdktf/provider-random";
import { Stateful } from "./stateful";
import { Frontend } from "./frontend";
import { Backend } from "./backend";

export class BaseStack extends TerraformStack {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    const props = {
      wsHostname: "wsapi",
      apiHostname: "api",
      appHostname: "app",
      filesHostname: "files",
      domainName: "iqvine.com",
      publicBucketName: "iqvine-public",
      privateBucketName: "iqvine-private",
      uploadBucketName: "iqvine-upload",
      logsBucketName: "logs",
    };

    new aws.provider.AwsProvider(this, "AWS", {
      region: "us-west-2",
    });
    new randprovider.provider.RandomProvider(this, "rand");

    const stateful = new Stateful(this, "stateful", {
      domainName: props.domainName,
      filesHostname: props.filesHostname,
      publicBucketPrefxi: props.publicBucketName,
      privateBucketPrefix: props.privateBucketName,
      uploadBucketPrefix: props.uploadBucketName,
      logsBucketPrefix: props.logsBucketName,
      apiHostname: props.apiHostname,
    });

    const frontend = new Frontend(this, "frontend", {
      wsHostname: props.wsHostname,
      appHostname: props.appHostname,
      zone: stateful.zone,
      domainName: props.domainName,
      apigw: stateful.apigw,
      filesDomainName: stateful.fileDomainName,
      logsBucket: stateful.logsBucket,
    });

    new Backend(this, "backend-parts", {
      apigw: stateful.apigw,
      apiDomainName: stateful.apiDomainName,
      wsconf: frontend.wsconf,
      uploadBucket: stateful.uploadBucket,
      publicBucket: stateful.publicBucket,
      privateBucket: stateful.privateBucket,
    });
  }
}

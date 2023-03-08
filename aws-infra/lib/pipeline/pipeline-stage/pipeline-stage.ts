import * as cdk from "aws-cdk-lib";

import { Construct } from "constructs";
import { EnvironmentConfig } from "../pipeline-types/pipeline-types";
import { StatefulStack } from "../../app/stateful/stateful-stack";
import { ApiStack } from "../../app/stateless/api-new";
import { FrontendStack } from "../../app/stateless/frontend";
import { BackendStack } from "../../app/stateless/api-backend";

export class PipelineStage extends cdk.Stage {
  constructor(scope: Construct, id: string, props: EnvironmentConfig) {
    super(scope, id, props);

    new MonoStack(this, "mono", props);
  }
}

export class MonoStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: EnvironmentConfig) {
    super(scope, id, props);

    const stateful = new StatefulStack(this, "StatefulStack", {
      publicBucketName: props.shared.publicBucketName,
      privateBucketName: props.shared.privateBucketName,
      uploadBucketName: props.shared.uploadBucketName,
      domainName: props.stateful.domainName,
      appHostname: props.shared.appHostname,
      filesHostname: props.shared.filesHostname,
    });

    const frontend = new FrontendStack(this, "frontend", {
      appHostname: props.shared.appHostname,
      hostedZone: stateful.hostedZone,
      domainName: props.stateful.domainName,
      apigw: stateful.apigw,
    });

    // new ApiStack(this, "APIStack", {
    //   distribution: frontend.distribution,
    //   apigw: stateful.apigw,
    //   hostedZone: stateful.hostedZone,

    //   uploadBucket: stateful.uploadBucket,
    //   publicBucket: stateful.publicBucket,
    //   privateBucket: stateful.privateBucket,

    //   appHostname: props.shared.appHostname,
    //   filesHostname: props.shared.filesHostname,
    // });

    // const { uploadBucket, publicBucket, privateBucket } = props;
    // const uploadBucket = cdk.aws_s3.Bucket.fromBucketArn(this, "upload", props.uploadBucketArn);
    // const privateBucket = cdk.aws_s3.Bucket.fromBucketArn(this, "private", props.privateBucketArn);
    // const publicBucket = cdk.aws_s3.Bucket.fromBucketArn(this, "public", props.publicBucketArn);

    // const uploadBucket = cdk.aws_s3.Bucket.fromBucketArn(
    //   this,
    //   "upload-bucket",
    //   cdk.Fn.importValue("bucket-upload-arn"),
    // );
    // const privateBucket = cdk.aws_s3.Bucket.fromBucketArn(
    //   this,
    //   "private-bucket",
    //   cdk.Fn.importValue("bucket-private-arn"),
    // );
    // const publicBucket = cdk.aws_s3.Bucket.fromBucketArn(
    //   this,
    //   "public-bucket",
    //   cdk.Fn.importValue("bucket-public-arn"),
    // );

    new BackendStack(this, "Backend", {
      apigw: stateful.apigw,
      wsapi: frontend.wsapi,
      wsstage: frontend.wsstage,
      // hostedZone: props.hostedZone,
      appHostname: props.shared.appHostname,
      filesHostname: props.shared.filesHostname,
      wsdb: frontend.wsdb,
      uploadBucket: stateful.uploadBucket,
      publicBucket: stateful.publicBucket,
      privateBucket: stateful.privateBucket,
    });
  }
}

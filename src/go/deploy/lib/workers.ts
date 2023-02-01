import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { SubscriptionFilter } from "aws-cdk-lib/aws-sns";
import { QueueWorker, wireLambda } from "./utils";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

///
//
//
export class CreateWorkersFile extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      coreFile,
    }: {
      eventbus: cdk.aws_sns.Topic;
      coreFile: {
        uploadBucket: cdk.aws_s3.Bucket;
        privateBucket: cdk.aws_s3.Bucket;
        publicBucket: cdk.aws_s3.Bucket;
      };
    },
  ) {
    super(scope, id);

    const worker = new QueueWorker(this, "upload", {
      eventbus,
      memorySize: 512,
      path: "workers_file",
      env: {
        SQS_HANDLER: "event:file_uploaded",
        PRIVATE_BUCKET: coreFile.privateBucket.bucketName,
        PUBLIC_BUCKET: coreFile.publicBucket.bucketName,
      },
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.FileEventbus/FileUploaded"],
        }),
      },
    });

    wireLambda(this, worker.lambda, {
      s3buckets: [coreFile.uploadBucket, coreFile.publicBucket],
    });
  }
}

export class CreateWorkersUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    new QueueWorker(this, "password-changed", {
      eventbus,
      path: "workers_user",
      env: { SQS_HANDLER: "userSecurity/password_changed" },
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityPasswordChange"],
        }),
      },
    });

    new QueueWorker(this, "register-token", {
      eventbus,
      path: "workers_user",
      env: { SQS_HANDLER: "userSecurity/register" },
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityRegisterToken"],
        }),
      },
    });

    new QueueWorker(this, "forgot-request", {
      eventbus,
      path: "workers_user",
      env: { SQS_HANDLER: "userSecurity/forgot" },
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityForgotRequest"],
        }),
      },
    });

    new QueueWorker(this, "user-invite", {
      eventbus,
      path: "workers_user",
      env: { SQS_HANDLER: "userSecurity/invite" },
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityInviteToken"],
        }),
      },
    });
  }
}

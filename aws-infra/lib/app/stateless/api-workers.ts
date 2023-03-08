import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { SubscriptionFilter } from "aws-cdk-lib/aws-sns";
import { GoHandler } from "./api-utils";

export class CreateWorkersFile extends Construct {
  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      uploadBucket,
      publicBucket,
      privateBucket,
    }: {
      eventbus: cdk.aws_sns.Topic;
      uploadBucket: cdk.aws_s3.IBucket;
      privateBucket: cdk.aws_s3.IBucket;
      publicBucket: cdk.aws_s3.IBucket;
    },
  ) {
    super(scope, id);

    const handler = new GoHandler(this, "worker-file", ["workers", "workers_file"], {
      memorySize: 512,
      environment: {
        SQS_HANDLER: "event:file_uploaded",
        PRIVATE_BUCKET: privateBucket.bucketName,
        PUBLIC_BUCKET: publicBucket.bucketName,
      },
      s3buckets: [uploadBucket, publicBucket],
    });
    handler.eventQueue("worker-file", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.FileEventbus/FileUploaded"],
        }),
      },
    });
  }
}

export class CreateWorkersUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const path = ["workers", "workers_user"];

    const changed = new GoHandler(this, "password-changed", path, {
      environment: { SQS_HANDLER: "userSecurity/password_changed" },
    });
    changed.eventQueue("password-changed", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityPasswordChange"],
        }),
      },
    });

    const token = new GoHandler(this, "register-token", path, {
      environment: { SQS_HANDLER: "userSecurity/register" },
    });
    token.eventQueue("register-token", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityRegisterToken"],
        }),
      },
    });

    const forgot = new GoHandler(this, "register-forgot", path, {
      environment: { SQS_HANDLER: "userSecurity/forgot" },
    });
    forgot.eventQueue("register-forgot", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityForgotRequest"],
        }),
      },
    });

    const invite = new GoHandler(this, "register-invite", path, {
      environment: { SQS_HANDLER: "userSecurity/invite" },
    });
    invite.eventQueue("register-invite", eventbus, {
      filterPolicy: {
        "twirp.path": SubscriptionFilter.stringFilter({
          allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/SecurityInviteToken"],
        }),
      },
    });
  }
}

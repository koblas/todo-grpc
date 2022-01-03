import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";

export interface SNSPublishArgs {
  topicArn: pulumi.Input<aws.ARN>;
}

export class SNSPublishPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy;

  constructor(name: string, args: SNSPublishArgs, opts?: pulumi.ComponentResourceOptions) {
    super("aws:components:SNSPublishPolicy", name, args, opts);

    const defaultResourceOptions: pulumi.ResourceOptions = { parent: this };

    this.policy = new aws.iam.Policy(
      name,
      {
        description: `IAM policy for publishing to SNS`,
        policy: {
          Version: "2012-10-17",
          Statement: [
            {
              Effect: "Allow",
              Action: ["sns:Publish"],
              Resource: args.topicArn,
            },
          ],
        },
      },
      defaultResourceOptions,
    );

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } });
  }
}

import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

export interface SQSPolicyArgs {
  queueArn: pulumi.Input<aws.ARN>;
  sourceArn?: pulumi.Input<aws.ARN>;
}
export type SQSPublishPolicyArgs = SQSPolicyArgs;
export type SQSProcessPolicyArgs = SQSPolicyArgs;

export class SQSPublishPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy;
  constructor(name: string, args: SQSPublishPolicyArgs, opts?: pulumi.ComponentResourceOptions) {
    super("caya:SQSPublishPolicy", name, args, opts);
    const defaultParentOptions: pulumi.ResourceOptions = { parent: this };
    this.policy = new aws.iam.Policy(
      name,
      {
        policy: {
          Version: "2012-10-17",
          Statement: [
            {
              Effect: "Allow",
              Action: ["sqs:SendMessage"],
              Resource: args.queueArn,
              ...(args.sourceArn
                ? {
                    Condition: {
                      ArnEquals: {
                        "aws:SourceArn": args.sourceArn,
                      },
                    },
                  }
                : {}),
            },
          ],
        },
      },
      defaultParentOptions,
    );

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } });
  }
}

export class SQSProcessPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy;
  constructor(name: string, args: SQSProcessPolicyArgs, opts?: pulumi.ComponentResourceOptions) {
    super("aws:components:SQSProcessPolicy", name, args, opts);
    const defaultResourceOptions: pulumi.ResourceOptions = { parent: this };

    this.policy = new aws.iam.Policy(
      name,
      {
        policy: {
          Version: "2012-10-17",
          Statement: [
            {
              Effect: "Allow",
              Action: [
                "sqs:GetQueueUrl",
                "sqs:ReceiveMessage",
                "sqs:DeleteMessage",
                "sqs:GetQueueAttributes",
                "sqs:ChangeMessageVisibility",
              ],
              Resource: args.queueArn,
            },
          ],
        },
      },
      defaultResourceOptions,
    );

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } });
  }
}

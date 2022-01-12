import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";

export interface DynamoRWArgs {
  dbArn: pulumi.Output<string>;
}

export class DynamoRWPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy;

  constructor(name: string, args: DynamoRWArgs, opts?: pulumi.ComponentResourceOptions) {
    super("aws:components:DynamoRWPolicy", name, args, opts);

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
              Action: [
                "dynamodb:BatchGetItem",
                "dynamodb:GetItem",
                "dynamodb:Query",
                "dynamodb:Scan",
                "dynamodb:BatchWriteItem",
                "dynamodb:PutItem",
                "dynamodb:UpdateItem",
                "dynamodb:DeleteItem",
              ],
              Resource: args.dbArn.apply((v) => [v, `${v}/*`, `${v}/index/*`]),
            },
          ],
        },
      },
      defaultResourceOptions,
    );

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } });
  }
}

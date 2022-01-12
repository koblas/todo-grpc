import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";

export interface SSMParameterArgs {
  regionName: pulumi.Input<aws.GetRegionResult["name"]>;
  accountId: pulumi.Input<aws.GetCallerIdentityResult["accountId"]>;
}

export class SSMParameterPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy;

  constructor(name: string, args: SSMParameterArgs, opts?: pulumi.ComponentResourceOptions) {
    super("aws:components:SSMParameterPolicy", name, args, opts);

    const defaultResourceOptions: pulumi.ResourceOptions = { parent: this };

    this.policy = new aws.iam.Policy(
      name,
      {
        description: "IAM policy for reading SSM parameters",
        policy: {
          Version: "2012-10-17",
          Statement: [
            {
              Effect: "Allow",
              Action: ["ssm:DescribeParameters"],
              Resource: "*",
            },
            {
              Effect: "Allow",
              Action: ["ssm:GetParameter", "ssm:GetParameters", "ssm:GetParametersByPath"],
              Resource: `arn:aws:ssm:${args.regionName}:${args.accountId}:parameter/*`,
            },
          ],
        },
      },
      defaultResourceOptions,
    );

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } });
  }
}

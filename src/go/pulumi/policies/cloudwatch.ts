import * as aws from '@pulumi/aws'
import * as pulumi from '@pulumi/pulumi'

export interface LambdaCloudWatchPolicyArgs {
  /**
   * Lambda name
   */
  lambdaName: string

  /**
   * Specifies the number of days you want to retain log events in the specified log group.
   * Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.
   */
  logRetentionInDays?: number
}

/**
 * Cloudwatch policy for Lambda resource
 */
export class LambdaCloudWatchPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy
  readonly logGroup: aws.cloudwatch.LogGroup

  constructor(name: string, args: LambdaCloudWatchPolicyArgs, opts?: pulumi.ComponentResourceOptions) {
    super('aws:components:LambdaCloudWatchPolicy', name, args, opts)

    const defaultResourceOptions: pulumi.ResourceOptions = { parent: this }
    const { lambdaName } = args

    this.policy = new aws.iam.Policy(
      name,
      {
        name,
        description: `IAM policy for logging from ${lambdaName} lambda`,
        policy: {
          Version: '2012-10-17',
          Statement: [
            {
              Effect: 'Allow',
              Action: ['logs:CreateLogGroup', 'logs:CreateLogStream', 'logs:PutLogEvents'],
              Resource: `arn:aws:logs:*:*:log-group:/aws/lambda/${lambdaName}*`
            }
          ]
        }
      },
      defaultResourceOptions
    )

    this.logGroup = new aws.cloudwatch.LogGroup(
      `${name}-log-group`,
      {
        name: `/aws/lambda/${args.lambdaName}`,
        ...(args.logRetentionInDays
          ? {
              retentionInDays: args.logRetentionInDays
            }
          : {})
      },
      defaultResourceOptions
    )

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } })
  }
}
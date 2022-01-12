import * as aws from '@pulumi/aws'
import * as pulumi from '@pulumi/pulumi'

interface S3PolicyArgs {
  /**
   * resource ARN for bucket policy
   */
  bucketArn: pulumi.Output<aws.ARN>
}

type S3ReadPolicyArgs = S3PolicyArgs
type S3ReadWritePolicyArgs = S3PolicyArgs

export class S3ReadPolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy

  constructor(name: string, args: S3ReadPolicyArgs, opts?: pulumi.ComponentResourceOptions) {
    super('aws:components:S3ReadPolicy', name, args, opts)
    const { bucketArn } = args
    const defaultParentOptions: pulumi.ResourceOptions = { parent: this }
    this.policy = new aws.iam.Policy(
      name,
      {
        name,
        policy: {
          Version: '2012-10-17',
          Statement: [
            {
              Action: ['s3:GetObject'],
              Effect: 'Allow',
              Resource: pulumi.interpolate`${bucketArn}/*`
            }
          ]
        }
      },
      defaultParentOptions
    )

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } })
  }
}

export class S3ReadWritePolicy extends pulumi.ComponentResource {
  readonly policy: aws.iam.Policy
  constructor(name: string, args: S3ReadWritePolicyArgs, opts?: pulumi.ComponentResourceOptions) {
    super('aws:components:S3ReadWritePolicy', name, args, opts)
    const defaultParentOptions: pulumi.ResourceOptions = { parent: this }
    this.policy = new aws.iam.Policy(
      name,
      {
        policy: {
          Version: '2012-10-17',
          Statement: [
            {
              Effect: 'Allow',
              Action: ['s3:GetObject', 's3:PutObject', 's3:DeleteObject'],
              Resource: pulumi.interpolate`${args.bucketArn}/*`
            },
            {
              Action: ['s3:ListBucket'],
              Effect: 'Allow',
              Resource: args.bucketArn
            }
          ]
        }
      },
      defaultParentOptions
    )

    this.registerOutputs({ policy: { name: this.policy.name, arn: this.policy.arn } })
  }
}
import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";

import { LambdaCloudWatchPolicy } from "./policies";
import { attachPoliciesToRole } from "./utils";

/**
 * Arguments to LambdaFunction
 */
export interface LambdaFunctionArgs extends Omit<aws.lambda.FunctionArgs, "name" | "role"> {
  /**
   * Additional policies to attach to lambda role
   */
  policies?: aws.iam.Policy[];

  /**
   * The Lambda environment's configuration settings.
   */
  // environment?: {
  // [key: string]: pulumi.Input<string>;
  // };
}

/**
 * creates a lambda with cloudwatch log group policy.
 *
 * ```typescript
 * import { LambdaFunction, S3ReadPolicy } from 'pulumi-aws-components'
 *
 * const s3ReadPolicy = new S3ReadPolicy('', {
 *  bucketArn: `<S3 Bucket ARN>`
 * })
 *
 * const lambda = new LambdaFunction('my-lambda', {
 *  policyArns: [s3ReadPolicy.policy], # Additional policies to attach to lambda
 *  environment: {
 *    'keyA': 'valueA',
 *    ...
 *  },
 *  # ... other aws.lambda.FunctionArgs
 * })
 *
 * ```
 */
export class LambdaFunction extends pulumi.ComponentResource {
  readonly role: aws.iam.Role;
  readonly lambda: aws.lambda.Function;
  readonly roleAttachments: pulumi.Input<aws.iam.RolePolicyAttachment>[];
  readonly logGroup: pulumi.Input<aws.cloudwatch.LogGroup>;

  /**
   * Creates a new Lambda function with a default cloudwatch policy.
   *
   * @param name The _unique_ name of the resource.
   * @param args The arguments to configure the lambda.
   * @param opts A bag of options that control this resource's behavior.
   */
  constructor(name: string, args: LambdaFunctionArgs, opts?: pulumi.CustomResourceOptions) {
    super("aws:components:LambdaFunction", name, args, opts);

    // Default resource options for this component's child resources.
    const defaultResourceOptions: pulumi.ResourceOptions = { parent: this };

    const roleName = `${name}-role`;
    this.role = new aws.iam.Role(
      roleName,
      {
        name: roleName,
        assumeRolePolicy: aws.iam.assumeRolePolicyForPrincipal({
          Service: ["lambda.amazonaws.com"],
        }),
      },
      defaultResourceOptions,
    );

    this.lambda = new aws.lambda.Function(
      name,
      {
        memorySize: 1024,
        ...args,
        // environment: {
        // variables: {
        // ...(args.environment || {}),
        // },
        // },
        name,
        role: this.role.arn,
      },
      defaultResourceOptions,
    );

    // to manage the CloudWatch Log Group for the Lambda Function.
    const cloudWatchPolicy = new LambdaCloudWatchPolicy(`${name}-policy`, { lambdaName: name }, defaultResourceOptions);

    // Attach any additional policies
    this.roleAttachments = attachPoliciesToRole(
      this.role,
      [...(args.policies || []), cloudWatchPolicy.policy],
      defaultResourceOptions,
    );

    // Create the log group with default setup
    this.logGroup = new aws.cloudwatch.LogGroup(
      `${name}-cloudwatch`,
      {
        retentionInDays: 1,
        name: this.lambda.name.apply((v) => `/aws/lambda/${v}`),
      },
      defaultResourceOptions,
    );

    this.registerOutputs({
      lambda: { name: this.lambda.name, arn: this.lambda.arn },
      role: { name: this.role.name, arn: this.role.arn },
      roleAttachments: this.roleAttachments,
      logGroup: this.logGroup.arn,
    });
  }
}

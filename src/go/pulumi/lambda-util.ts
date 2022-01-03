import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";
import { PolicyArgs } from "@pulumi/aws/iam";
import { LambdaFunction } from "./lambda-function";
import * as policies from "./policies";

export type LambdaProps = Omit<aws.lambda.FunctionArgs, "role"> & {
  policies?: aws.iam.Policy[];
};

export async function buildLambda(
  name: string,
  props: LambdaProps,
): Promise<{ lambda: aws.lambda.Function; role: aws.iam.Role }> {
  // const lambda = new LambdaFunction(name, {
  //   code: props.code,
  //   environment: props.environment,
  //   handler: "app",
  //   runtime: "go1.x",
  //   policies: [
  //     new policies.SSMParameterPolicy(name, {
  //       regionName: (await aws.getRegion()).name,
  //       accountId: (await aws.getCallerIdentity()).accountId,
  //     }).policy,
  //   ],
  // });

  // return { lambda: lambda.lambda, role: lambda.role };
  // // Create the role for the Lambda to assume
  const lambdaRole = new aws.iam.Role(`${name}-assume-role`, {
    assumeRolePolicy: {
      Version: "2012-10-17",
      Statement: [
        {
          Action: "sts:AssumeRole",
          Principal: {
            Service: "lambda.amazonaws.com",
          },
          Effect: "Allow",
          Sid: "",
        },
      ],
    },
  });

  const policies: Record<string, aws.iam.PolicyDocument["Statement"]> = {
    [`${name}-ssm-access`]: [
      {
        Effect: "Allow",
        Action: ["ssm:DescribeParameters"],
        Resource: "*",
      },
      {
        Effect: "Allow",
        Action: ["ssm:GetParameter", "ssm:GetParameters", "ssm:GetParametersByPath"],
        Resource: `arn:aws:ssm:${(await aws.getRegion()).name}:${
          (await aws.getCallerIdentity()).accountId
        }:parameter/*`,
      },
    ],
    [`${name}-cloudwatch-write`]: [
      {
        Effect: "Allow",
        Action: ["logs:CreateLogGroup", "logs:CreateLogStream", "logs:PutLogEvents"],
        Resource: "*",
      },
    ],
    [`${name}-lambda-invoke`]: [
      {
        Effect: "Allow",
        Action: ["lambda:InvokeFunction"],
        Resource: `arn:aws:lambda:${(await aws.getRegion()).name}:${
          (await aws.getCallerIdentity()).accountId
        }:function:*`,
      },
    ],
  };

  Object.entries(policies).forEach(([policyName, statements]) => {
    const policy = new aws.iam.Policy(`${name}-${policyName}`, {
      name: policyName,
      policy: {
        Version: "2012-10-17",
        Statement: statements,
      },
    });

    new aws.iam.RolePolicyAttachment(`${name}-${policyName}`, {
      role: lambdaRole,
      policyArn: policy.arn,
    });
  });
  Object.entries(props.policies ?? []).forEach(([policyName, policy]) => {
    new aws.iam.RolePolicyAttachment(`${name}-${policyName}`, {
      role: lambdaRole,
      policyArn: policy.arn,
    });
  });

  const theFunction = new aws.lambda.Function(`${name}-lambda`, {
    name: name,
    code: props.code,
    // memorySize: 1024,
    memorySize: 512,
    environment: props.environment,
    handler: "app",
    runtime: "go1.x",
    role: lambdaRole.arn,
    timeout: props.timeout ?? 5,
  });

  // Create the log group with default setup
  new aws.cloudwatch.LogGroup(`${name}-cloudwatch`, {
    retentionInDays: 1,
    name: theFunction.name.apply((v) => `/aws/lambda/${v}`),
  });

  return { lambda: theFunction, role: lambdaRole };
}

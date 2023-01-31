import * as path from "path";
import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { GoFunction, GoFunctionProps } from "@aws-cdk/aws-lambda-go-alpha";
import { SqsToLambda } from "@aws-solutions-constructs/aws-sqs-lambda";
import { LambdaToDynamoDB } from "@aws-solutions-constructs/aws-lambda-dynamodb";
import { LambdaToSns } from "@aws-solutions-constructs/aws-lambda-sns";

const LAMBDA_DEFAULTS: Partial<GoFunctionProps> = {
  logRetention: cdk.Duration.days(3).toDays(),
  insightsVersion: cdk.aws_lambda.LambdaInsightsVersion.VERSION_1_0_135_0,
  architecture: cdk.aws_lambda.Architecture.ARM_64,
  bundling: {
    goBuildFlags: [`-ldflags="-s -w"`],
  },
};

export function goFunction(
  scope: Construct,
  name: string,
  paths: string[],
  params?: Partial<GoFunctionProps>,
): GoFunction {
  const lambda = new GoFunction(scope, [...paths, "handler"].join("-"), {
    functionName: name,
    entry: path.join(__dirname, "..", "..", "cmd", "lambda", ...paths),
    ...LAMBDA_DEFAULTS,
    ...params,
  });

  return lambda;
}

export function wireLambda(
  scope: Construct,
  lambda: cdk.aws_lambda.Function,
  {
    eventbus,
    parameters,
    dynamo,
    s3buckets,
  }: {
    dynamo?: cdk.aws_dynamodb.ITable;
    eventbus?: cdk.aws_sns.Topic;
    parameters?: string[];
    s3buckets?: cdk.aws_s3.Bucket[];
  },
) {
  if (parameters?.length) {
    cdk.aws_iam.Grant.addToPrincipal({
      grantee: lambda,
      actions: ["ssm:DescribeParameters", "ssm:GetParameters", "ssm:GetParameter", "ssm:GetParameterHistory"],
      resourceArns: parameters.map((p) =>
        cdk.Stack.of(scope).formatArn({
          service: "ssm",
          resource: `parameter${p}`,
        }),
      ),
    });
  }

  // Invoke and SendMessage are good downstream calls
  cdk.aws_iam.Grant.addToPrincipal({
    grantee: lambda,
    actions: ["lambda:InvokeFunction"],
    resourceArns: [
      cdk.Stack.of(scope).formatArn({
        service: "lambda",
        resource: "function:*",
      }),
    ],
  });

  cdk.aws_iam.Grant.addToPrincipal({
    grantee: lambda,
    actions: ["sqs:GetQueueUrl", "sqs:SendMessage"],
    resourceArns: [
      cdk.Stack.of(scope).formatArn({
        service: "sqs",
        resource: "*",
      }),
    ],
  });

  if (eventbus) {
    new LambdaToSns(scope, "sns-perms", {
      existingLambdaObj: lambda,
      existingTopicObj: eventbus,
    });
  }

  if (dynamo) {
    new LambdaToDynamoDB(scope, "dynamo-perms", {
      existingTableObj: dynamo as cdk.aws_dynamodb.Table,
      tablePermissions: "ReadWrite",
      existingLambdaObj: lambda,
    });
  }

  if (s3buckets) {
    s3buckets.forEach((bucket) => {
      bucket.grantReadWrite(lambda);
    });
  }
}

export class QueueWorker extends Construct {
  lambda: cdk.aws_lambda.Function;

  constructor(
    scope: Construct,
    id: string,
    {
      path,
      eventbus,
      env,
      filterPolicy,
      memorySize,
    }: {
      path: string;
      eventbus: cdk.aws_sns.Topic;
      env: Record<string, string>;
      memorySize?: GoFunctionProps["memorySize"];
      filterPolicy: NonNullable<cdk.aws_sns_subscriptions.SubscriptionProps["filterPolicy"]>;
    },
  ) {
    super(scope, id);

    this.lambda = goFunction(this, `core-workers-${id}`, ["workers", path], {
      environment: env,
      memorySize,
    });

    const worker = new QueueLambda(this, id, {
      eventbus,
      queueProps: {
        // SNS cannot deliver to encrypted SQS queues
        encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
      },
      existingLambdaObj: this.lambda,
    });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(worker.queue, {
        rawMessageDelivery: true,
        filterPolicy,
      }),
    );
  }
}

export class QueueLambda extends Construct {
  queue: cdk.aws_sqs.Queue;

  constructor(
    scope: Construct,
    id: string,
    {
      eventbus,
      lambdaFunctionProps,
      existingLambdaObj,
      queueProps,
    }: {
      eventbus: cdk.aws_sns.Topic;
      queueProps?: cdk.aws_sqs.QueueProps;
      lambdaFunctionProps?: cdk.aws_lambda.FunctionProps;
      existingLambdaObj?: cdk.aws_lambda.Function;
    },
  ) {
    super(scope, id);

    // Connect the queue
    const inst = new SqsToLambda(this, "sqs", {
      ...(lambdaFunctionProps
        ? {
            lambdaFunctionProps: {
              functionName: `worker-${id}`,
              logRetention: cdk.Duration.days(3).toDays(),
              ...lambdaFunctionProps,
            },
          }
        : {}),
      ...(existingLambdaObj ? { existingLambdaObj } : {}),
      deadLetterQueueProps: {
        queueName: `${id}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${id}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        ...queueProps,
      },
    });

    this.queue = inst.sqsQueue;

    // Make sure we can write to SNS
    wireLambda(this, inst.lambdaFunction, { eventbus, parameters: ["/common/*"] });
  }
}

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

export class GoHandler extends Construct {
  lambda: cdk.aws_lambda.Function;

  constructor(
    scope: Construct,
    name: string,
    paths: string[],
    params?: Partial<GoFunctionProps> & {
      dynamo?: cdk.aws_dynamodb.ITable;
      eventbus?: cdk.aws_sns.Topic;
      parameters?: string[];
      s3buckets?: cdk.aws_s3.IBucket[];
    },
  ) {
    super(scope, name);

    // this.lambda = new GoFunction(scope, [...paths, "handler"].join("-"), {
    this.lambda = new GoFunction(this, name, {
      functionName: name,
      entry: path.join(__dirname, "..", "..", "..", "..", "src", "go", "cmd", "lambda", ...paths),
      ...LAMBDA_DEFAULTS,
      ...params,
    });

    this.wireLambda({
      dynamo: params?.dynamo,
      eventbus: params?.eventbus,
      parameters: params?.parameters,
      s3buckets: params?.s3buckets,
    });
  }

  wireLambda({
    eventbus,
    parameters,
    dynamo,
    s3buckets,
  }: {
    dynamo?: cdk.aws_dynamodb.ITable;
    eventbus?: cdk.aws_sns.Topic;
    parameters?: string[];
    s3buckets?: cdk.aws_s3.IBucket[];
  }) {
    if (parameters?.length) {
      cdk.aws_iam.Grant.addToPrincipal({
        grantee: this.lambda,
        actions: ["ssm:DescribeParameters", "ssm:GetParameters", "ssm:GetParameter", "ssm:GetParameterHistory"],
        resourceArns: parameters.map((p) =>
          cdk.Stack.of(this).formatArn({
            service: "ssm",
            resource: `parameter${p}`,
          }),
        ),
      });
    }

    // Invoke and SendMessage are good downstream calls
    cdk.aws_iam.Grant.addToPrincipal({
      grantee: this.lambda,
      actions: ["lambda:InvokeFunction"],
      resourceArns: [
        cdk.Stack.of(this).formatArn({
          service: "lambda",
          resource: "function:*",
        }),
      ],
    });

    cdk.aws_iam.Grant.addToPrincipal({
      grantee: this.lambda,
      actions: ["sqs:GetQueueUrl", "sqs:SendMessage"],
      resourceArns: [
        cdk.Stack.of(this).formatArn({
          service: "sqs",
          resource: "*",
        }),
      ],
    });

    if (eventbus) {
      new LambdaToSns(this, "sns-perms", {
        existingLambdaObj: this.lambda,
        existingTopicObj: eventbus,
      });
    }

    if (dynamo) {
      new LambdaToDynamoDB(this, "dynamo-perms", {
        existingTableObj: dynamo as cdk.aws_dynamodb.Table,
        tablePermissions: "ReadWrite",
        existingLambdaObj: this.lambda,
      });
    }

    if (s3buckets) {
      // const policy = new cdk.aws_iam.PolicyStatement();
      // policy.addActions("s3:Get*", "s3:Put*", "s3:List*", "s3:Delete*");
      // policy.addResources(...s3buckets.map((b) => b.bucketArn));

      // console.log(policy.toJSON());
      // this.lambda.addToRolePolicy(policy);
      s3buckets.forEach((bucket) => {
        bucket.grantReadWrite(this.lambda);
      });
    }
  }

  listenQueue(name: string, queueProps?: cdk.aws_sqs.QueueProps) {
    const inst = new SqsToLambda(this, "sqs", {
      // ...(lambdaFunctionProps
      //   ? {
      //       lambdaFunctionProps: {
      //         functionName: `worker-${id}`,
      //         logRetention: cdk.Duration.days(3).toDays(),
      //         ...lambdaFunctionProps,
      //       },
      //     }
      //   : {}),
      // ...(existingLambdaObj ? { existingLambdaObj } : {}),
      existingLambdaObj: this.lambda,
      deadLetterQueueProps: {
        queueName: `${name}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${name}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        ...queueProps,
      },
    });

    return inst.sqsQueue;
  }

  eventQueue(name: string, eventbus: cdk.aws_sns.Topic, props?: cdk.aws_sns_subscriptions.SqsSubscriptionProps) {
    const queue = this.listenQueue(name, {
      // SNS cannot deliver to encrypted SQS queues
      encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
    });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(queue, {
        rawMessageDelivery: true,
        ...props,
      }),
    );

    return queue;
  }
}

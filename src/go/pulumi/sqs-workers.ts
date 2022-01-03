import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";
import { LambdaFunctionArgs, LambdaFunction } from "./lambda-function";
import { buildLambda } from "./lambda-util";
import { SQSProcessPolicy, SQSPublishPolicy } from "./policies";

export async function sqsWorkers(corestack: pulumi.StackReference) {
  await buildWorker(
    corestack,
    "user-security-password-changed",
    { SQS_HANDLER: "userSecurity/password_changed" },
    {
      stream: ["event:user_security"],
      action: ["USER_PASSWORD_CHANGE"],
    },
  );
  await buildWorker(
    corestack,
    "user-security-register",
    { SQS_HANDLER: "userSecurity/register" },
    {
      stream: ["event:user_security"],
      action: ["USER_REGISTER_TOKEN"],
    },
  );
  await buildWorker(
    corestack,
    "user-security-forgot",
    { SQS_HANDLER: "userSecurity/forgot" },
    {
      stream: ["event:user_security"],
      action: ["USER_FORGOT_REQUEST"],
    },
  );
  await buildWorker(
    corestack,
    "user-security-invite",
    { SQS_HANDLER: "userSecurity/invite" },
    {
      stream: ["event:user_security"],
      action: ["USER_INVITE_TOKEN"],
    },
  );
}

async function buildWorker(
  corestack: pulumi.StackReference,
  name: string,
  env: Record<string, string>,
  filter: Record<string, string[]>,
) {
  const queueDlq = new aws.sqs.Queue(`${name}-dlq`, {});
  const queue = new aws.sqs.Queue(name, {
    redrivePolicy: queueDlq.arn.apply((arn) =>
      JSON.stringify({
        deadLetterTargetArn: arn,
        maxReceiveCount: 2,
      }),
    ),
    visibilityTimeoutSeconds: 300,
  });

  new aws.sqs.QueuePolicy(name, {
    queueUrl: queue.id,
    policy: {
      Version: "2012-10-17",
      Statement: [
        {
          Effect: "Allow",
          Principal: {
            Service: "sns.amazonaws.com",
          },
          Action: ["sqs:SendMessage"],
          Resource: queue.arn,
          Condition: {
            ArnEquals: {
              "aws:SourceArn": await corestack.getOutputValue("entityTopicArn"),
            },
          },
        },
      ],
    },
  });

  new aws.sns.TopicSubscription(`${name}-subscription`, {
    protocol: "sqs",
    topic: await corestack.getOutputValue("entityTopicArn"),
    endpoint: queue.arn,
    filterPolicy: JSON.stringify(filter),
    rawMessageDelivery: true,
  });

  const { lambda } = await buildQueueLambda(name, {
    handler: "app",
    runtime: "go1.x",
    code: new pulumi.asset.AssetArchive({
      app: new pulumi.asset.FileAsset("../build/core-workers"),
    }),
    environment: {
      variables: env,
    },
    queue,
  });

  return { lambda };
}

export interface QueueLambdaArgs extends Omit<LambdaFunctionArgs, "role"> {
  queue: aws.sqs.Queue;
  queueBatchSize?: number;
}

async function buildQueueLambda(name: string, args: QueueLambdaArgs) {
  const { queue, queueBatchSize = 10, environment, ...lambdaArgs } = args;

  const sqsPolicyName = `${name}-policy-sqs`;
  const queuePolicy = new SQSProcessPolicy(sqsPolicyName, { queueArn: queue.arn });

  const { lambda } = await buildLambda(name, {
    ...lambdaArgs,
    // policies: [...(lambdaArgs.policies || []), queuePolicy.policy],
    policies: [queuePolicy.policy],
    environment,
  });

  queue.onEvent(`${name}-queue-event-subscription`, lambda, {
    batchSize: queueBatchSize,
  });

  return {
    lambda: lambda,
  };
}

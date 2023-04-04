import { Construct } from "constructs";
import * as fs from "fs";
import * as path from "path";
import * as zip from "adm-zip";
import * as aws from "@cdktf/provider-aws";
import { Architecture, Runtime, StackLambda, StackLambdaConfig } from "./lambda";
import { AssetType, TerraformAsset } from "cdktf";

interface Props extends Partial<StackLambdaConfig> {
  // environment?: Record<string, string>;
  path?: string[];

  apiTrigger?: aws.apigatewayv2Api.Apigatewayv2Api;
  dynamo?: aws.dynamodbTable.DynamodbTable;
  eventbus?: aws.snsTopic.SnsTopic;
  parameters?: string[];
  s3buckets?: aws.s3Bucket.S3Bucket[];

  // allowedTriggers?: StackLambdaConfig["allowedTriggers"];
  // attachPolicies?: StackLambdaConfig["attachPolicies"];
}

export class GoHandler extends Construct {
  public lambda: aws.lambdaFunction.LambdaFunction;
  public role?: aws.iamRole.IamRole;

  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    const handler = new StackLambda(this, "lambda", {
      functionName: id,
      asset: props.path ? createBootstrapAsset(this, props.path.join("-")) : createStubAsset(this, id),
      architecture: Architecture.ARM,
      runtime: Runtime.AL2,
      handler: "bootstrap",
      // environment: props.environment,
      cloudwatchLogsRetentionInDays: 3,
      ...(props.apiTrigger
        ? {
            allowedTriggers: [
              {
                action: "lambda:InvokeFunction",
                principal: "apigateway.amazonaws.com",
                sourceArn: `${props.apiTrigger.executionArn}/*/*`,
              },
            ],
          }
        : {}),
      ...props,
      attachPolicies: [
        ...createPolicies(this, id, {
          dynamo: props.dynamo,
          eventbus: props.eventbus,
          parameters: props.parameters,
          s3buckets: props.s3buckets,
        }),
        ...(props.attachPolicies ?? []),
      ],
      ...(props.path
        ? {}
        : {
            lifecycle: {
              ignoreChanges: ["source_code_hash", "filename"],
            },
          }),
    });

    if (!handler.lambda) {
      throw new Error("Lambda failed to create");
    }

    this.lambda = handler.lambda;
    this.role = handler.role;
  }

  public eventQueue(
    name: string,
    eventbus: aws.snsTopic.SnsTopic,
    params?: Partial<aws.snsTopicSubscription.SnsTopicSubscriptionConfig>,
  ) {
    /// add an
    const queue = this.listenQueue(name, {
      // SNS cannot deliver to encrypted SQS queues
      sqsManagedSseEnabled: false,
    });

    new aws.snsTopicSubscription.SnsTopicSubscription(this, `${name}-subscription`, {
      topicArn: eventbus.arn,
      protocol: "sqs",
      endpoint: queue.arn,
      rawMessageDelivery: true,
      ...params,
    });

    const doc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, `${name}-doc`, {
      statement: [
        {
          effect: "Allow",
          actions: ["sqs:SendMessage"],
          principals: [
            {
              type: "Service",
              identifiers: ["sns.amazonaws.com"],
            },
          ],
          resources: [queue.arn],
          condition: [
            {
              test: "ArnEquals",
              variable: "aws:SourceArn",
              values: [eventbus.arn],
            },
          ],
        },
      ],
    });

    new aws.sqsQueuePolicy.SqsQueuePolicy(this, `${name}-policy`, {
      queueUrl: queue.id,
      policy: doc.json,
    });

    return queue;
  }

  listenQueue(name: string, queueProps?: aws.sqsQueue.SqsQueueConfig) {
    const deadletter = new aws.sqsQueue.SqsQueue(this, `${name}-queue-dl`, {
      name: `${name}-dlq`,
      messageRetentionSeconds: 7 * 24 * 60 * 60, // 7 days
      visibilityTimeoutSeconds: 5 * 60, // 5 minutes
      sqsManagedSseEnabled: true,
      ...queueProps,
    });
    const queue = new aws.sqsQueue.SqsQueue(this, `${name}-queue`, {
      name,
      messageRetentionSeconds: 7 * 24 * 60 * 60, // 7 days
      visibilityTimeoutSeconds: 5 * 60, // 5 minutes
      redrivePolicy: JSON.stringify({
        deadLetterTargetArn: deadletter.arn,
        maxReceiveCount: 4,
      }),
      sqsManagedSseEnabled: true,
      ...queueProps,
    });

    // new aws.lambdaPermission.LambdaPermission(this, "sqs_perms", {
    //   functionName: this.lambda.functionName,
    //   action: "lambda:InvokeFunction",
    //   principal: "sqs.amazon.com",
    //   sourceArn: queue.arn,
    // });

    const doc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(this, "sqs_doc", {
      statement: [
        {
          effect: "Allow",
          actions: ["sqs:ReceiveMessage", "sqs:DeleteMessage", "sqs:GetQueueAttributes", "sqs:ChangeMessageVisibility"],
          resources: [queue.arn],
        },
      ],
    });
    const policy = new aws.iamPolicy.IamPolicy(this, "sqs_policy", {
      policy: doc.json,
    });

    if (this.role) {
      new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, "sqs_role_attach", {
        role: this.role.id,
        policyArn: policy.arn,
      });
    }

    new aws.lambdaEventSourceMapping.LambdaEventSourceMapping(this, "sqs_event", {
      eventSourceArn: queue.arn,
      functionName: this.lambda.arn,
      enabled: true,
      batchSize: 5,
    });

    return queue;
  }
}

function createStubAsset(scope: Construct, name: string): TerraformAsset {
  const archive = new zip();

  const output = path.join("/tmp/", `${name}-stub.zip`);
  archive.addFile("boostrap", Buffer.from(new Uint8Array()));
  const data = archive.toBuffer();

  fs.writeFileSync(output, data);

  return new TerraformAsset(scope, name, {
    path: output,
    type: AssetType.FILE,
  });
}

function createBootstrapAsset(scope: Construct, name: string): TerraformAsset {
  const archive = new zip();

  const input = path.join(__dirname, `../../../src/go/build/${name}`);
  archive.addLocalFile(input, "", "bootstrap");

  const output = path.join("/tmp/", `${name}.zip`);
  // archive.writeZip(output);
  const data = archive.toBuffer();

  fs.writeFileSync(output, data);

  return new TerraformAsset(scope, name, {
    path: output,
    type: AssetType.FILE,
  });
}

export function createPolicies(
  scope: Construct,
  id: string,
  {
    eventbus,
    parameters,
    dynamo,
    s3buckets,
  }: {
    dynamo?: aws.dynamodbTable.DynamodbTable;
    eventbus?: aws.snsTopic.SnsTopic;
    parameters?: string[];
    s3buckets?: aws.s3Bucket.S3Bucket[];
  },
): aws.iamPolicy.IamPolicy[] {
  const ident = new aws.dataAwsCallerIdentity.DataAwsCallerIdentity(scope, `${id}-identity`);
  const region = new aws.dataAwsRegion.DataAwsRegion(scope, `${id}-region`);

  function formatArn(p: {
    partition?: string;
    service?: string;
    region?: string;
    account?: string;
    resource?: string;
    sep?: string;
    resourceName?: string;
  }) {
    return `arn:${p.partition ?? "aws"}:${p.service ?? ""}:${p.region ?? region.name}:${p.account ?? ident.accountId}:${
      p.resource ?? ""
    }${p.sep ?? ""}${p.resourceName ?? ""}`;
  }

  const base = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(scope, "wire-base", {
    statement: [
      {
        effect: "Allow",
        actions: ["lambda:InvokeFunction"],
        resources: [formatArn({ service: "lambda", resource: "function:*" })],
      },
      {
        effect: "Allow",
        actions: ["sqs:GetQueueUrl", "sqs:SendMessage"],
        resources: [formatArn({ service: "sqs", resource: "*" })],
      },
    ],
  });

  const docs: aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument[] = [base];

  if (parameters?.length) {
    const doc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(scope, "wire-param", {
      statement: [
        {
          effect: "Allow",
          actions: ["ssm:DescribeParameters", "ssm:GetParameters", "ssm:GetParameter", "ssm:GetParameterHistory"],
          resources: parameters.map((p) => formatArn({ service: "ssm", resource: `parameter${p}` })),
        },
      ],
    });

    docs.push(doc);
  }

  if (eventbus) {
    const doc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(scope, "wire-topic", {
      statement: [
        {
          effect: "Allow",
          actions: ["sns:Publish"],
          resources: [eventbus.arn],
        },
      ],
    });

    docs.push(doc);
  }

  if (dynamo) {
    const doc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(scope, "wire-dynamo", {
      statement: [
        {
          effect: "Allow",
          actions: [
            "dynamodb:List*",
            "dynamodb:DescribeReservedCapacity*",
            "dynamodb:DescribeLimits",
            "dynamodb:DescribeTimeToLive",
          ],
          resources: ["*"],
        },
        {
          effect: "Allow",
          actions: [
            "dynamodb:BatchGet*",
            "dynamodb:DescribeStream",
            "dynamodb:DescribeTable",
            "dynamodb:Get*",
            "dynamodb:Query",
            "dynamodb:Scan",
            "dynamodb:BatchWrite*",
            "dynamodb:CreateTable",
            "dynamodb:Delete*",
            "dynamodb:Update*",
            "dynamodb:PutItem",
          ],
          resources: [formatArn({ service: "dynamodb", resource: `table/${dynamo.name}` })],
        },
      ],
    });

    docs.push(doc);
  }
  if (s3buckets) {
    const doc = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(scope, `wire-s3`, {
      statement: [
        {
          effect: "Allow",
          actions: ["s3:ListBucket"],
          // resources: s3buckets.map((bucket) => formatArn({ service: "s3", resource: bucket.bucket })),
          resources: s3buckets.map((bucket) => bucket.arn),
        },
        {
          effect: "Allow",
          actions: ["s3:PutObject", "s3:GetObject", "s3:PutObjectAcl", "s3:GetObjectAcl"],
          // resources: s3buckets.map((bucket) => formatArn({ service: "s3", resource: `${bucket.bucket}/*` })),
          resources: s3buckets.map((bucket) => `${bucket.arn}/*`),
        },
      ],
    });

    docs.push(doc);
  }

  const output = new aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocument(scope, "wire-output", {
    sourcePolicyDocuments: docs.map((d) => d.json),
  });

  return [
    new aws.iamPolicy.IamPolicy(scope, "attached-policy", {
      name: `${id}-lambda-service`,
      policy: output.json,
    }),
  ];
}

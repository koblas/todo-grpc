import * as cdk from "aws-cdk-lib";
import { WebSocketApi, WebSocketStage } from "@aws-cdk/aws-apigatewayv2-alpha";
import { Construct } from "constructs";
import { SqsToLambda } from "@aws-solutions-constructs/aws-sqs-lambda";
import { SubscriptionFilter } from "aws-cdk-lib/aws-sns";
import { goFunction, wireLambda } from "./utils";

export class WebsocketTodo extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const lambda = goFunction(this, "websocket-todo", ["websocket", "todo"], {});

    const queue = new SqsToLambda(this, "sqs", {
      existingLambdaObj: lambda,
      deadLetterQueueProps: {
        queueName: `${id}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${id}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        // SNS cannot deliver to encrypted SQS queues
        encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
      },
    });

    wireLambda(this, lambda, { parameters: ["/common/*"], eventbus });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(queue.sqsQueue, {
        rawMessageDelivery: true,
        filterPolicy: {
          "twirp.path": SubscriptionFilter.stringFilter({
            allowlist: ["/twirp/corepbv1.eventbus.TodoEventbus/TodoChange"],
          }),
        },
      }),
    );
  }
}

export class WebsocketUser extends Construct {
  constructor(scope: Construct, id: string, { eventbus }: { eventbus: cdk.aws_sns.Topic }) {
    super(scope, id);

    const lambda = goFunction(this, "websocket-user", ["websocket", "user"], {});

    const queue = new SqsToLambda(this, "sqs", {
      existingLambdaObj: lambda,
      deadLetterQueueProps: {
        queueName: `${id}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${id}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        // SNS cannot deliver to encrypted SQS queues
        encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
      },
    });

    wireLambda(this, lambda, { parameters: ["/common/*"], eventbus });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(queue.sqsQueue, {
        rawMessageDelivery: true,
        filterPolicy: {
          "twirp.path": SubscriptionFilter.stringFilter({
            allowlist: ["/twirp/corepbv1.eventbus.UserEventbus/UserChange"],
          }),
        },
      }),
    );
  }
}

export class WebsocketBroadcast extends Construct {
  constructor(
    scope: Construct,
    id: string,
    { eventbus, wsstage, wsapi }: { eventbus: cdk.aws_sns.Topic; wsstage: WebSocketStage; wsapi: WebSocketApi },
  ) {
    super(scope, id);

    const db = cdk.aws_dynamodb.Table.fromTableName(this, "conns", "ws-connection");

    const lambda = goFunction(this, "websocket-broadcast", ["websocket", "broadcast"], {
      environment: {
        CONN_DB: db.tableName,
        WS_ENDPOINT: wsstage.callbackUrl,
      },
    });

    wsstage.grantManagementApiAccess(lambda);

    const queue = new SqsToLambda(this, "sqs", {
      existingLambdaObj: lambda,
      deadLetterQueueProps: {
        queueName: `${id}-dlq`,
        retentionPeriod: cdk.Duration.days(7),
      },
      queueProps: {
        queueName: `${id}`,
        retentionPeriod: cdk.Duration.days(7),
        visibilityTimeout: cdk.Duration.minutes(5),
        // SNS cannot deliver to encrypted SQS queues
        encryption: cdk.aws_sqs.QueueEncryption.UNENCRYPTED,
      },
    });

    wireLambda(this, lambda, { parameters: ["/common/*"], dynamo: db });

    eventbus.addSubscription(
      new cdk.aws_sns_subscriptions.SqsSubscription(queue.sqsQueue, {
        rawMessageDelivery: true,
        filterPolicy: {
          "twirp.path": SubscriptionFilter.stringFilter({
            allowlist: ["/twirp/corepbv1.eventbus.BroadcastEventbus/Send"],
          }),
        },
      }),
    );
  }
}

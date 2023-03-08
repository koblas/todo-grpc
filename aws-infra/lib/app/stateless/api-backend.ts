import { HttpApi, IWebSocketApi, IWebSocketStage, WebSocketApi, WebSocketStage } from "@aws-cdk/aws-apigatewayv2-alpha";
import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { WebsocketBroadcast, WebsocketTodo, WebsocketUser } from "./api-websocket";
import { CreateWorkersFile, CreateWorkersUser } from "./api-workers";
import { PublicAuth, PublicFile, PublicTodo, PublicUser } from "./api-public";
import { CoreUser, CoreTodo, CoreOauthUser, CoreSendEmailQueue } from "./api-core";
import { TriggerS3 } from "./api-triggers";

export interface Props {
  apigw: HttpApi;
  wsdb: cdk.aws_dynamodb.ITable;
  wsapi: IWebSocketApi;
  wsstage: WebSocketStage;
  uploadBucket: cdk.aws_s3.IBucket;
  publicBucket: cdk.aws_s3.IBucket;
  privateBucket: cdk.aws_s3.IBucket;
  // hostedZone: cdk.aws_route53.IHostedZone;
  appHostname: string;
  filesHostname: string;
}

export class BackendStack extends Construct {
  // todo
  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    const eventbus = new cdk.aws_sns.Topic(this, "eventbus", {});

    // NOT USED const coreFile = new FileStorage(this, "core-file", { eventbus, hostedZone: props.hostedZone, hostname: "files" });
    // NOT USED new CoreSendEmail(this, "core-send-email", { eventbus });

    new CoreUser(this, "core-user", { eventbus });
    new CoreTodo(this, "core-todo", { eventbus });
    new CoreOauthUser(this, "core-oauth-user", { eventbus });
    new CoreSendEmailQueue(this, "send-email-queue", { eventbus });

    new PublicFile(this, "public-file", { eventbus, apigw: props.apigw, uploadBucket: props.uploadBucket });
    new PublicAuth(this, "public-auth", { eventbus, apigw: props.apigw });
    new PublicTodo(this, "public-todo", { eventbus, apigw: props.apigw });
    new PublicUser(this, "public-user", { eventbus, apigw: props.apigw });

    new CreateWorkersUser(this, "workers_user", { eventbus });
    new CreateWorkersFile(this, "workers_file", {
      eventbus,
      uploadBucket: props.uploadBucket,
      privateBucket: props.privateBucket,
      publicBucket: props.publicBucket,
    });

    new WebsocketTodo(this, "websocket-todo", { eventbus });
    new WebsocketUser(this, "websocket-user", { eventbus });
    new WebsocketBroadcast(this, "websocket-broadcast", {
      db: props.wsdb,
      eventbus,
      wsstage: props.wsstage,
      wsapi: props.wsapi,
    });

    new TriggerS3(this, "trigger-s3", { eventbus, bucket: props.uploadBucket });
  }
}

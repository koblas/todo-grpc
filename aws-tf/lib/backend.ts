import { Construct } from "constructs";
import * as aws from "@cdktf/provider-aws";
import { PublicAuth, PublicFile, PublicTodo, PublicUser } from "./api-public";
import { CoreOauthUser, CoreSendEmailQueue, CoreTodo, CoreUser } from "./api-core";
import { WebsocketBroadcast, WebsocketTodo, WebsocketUser } from "./api-websocket";
import { WorkerFile, WorkerUser } from "./api-workers";
import { TriggerS3 } from "./api-trigger";
import * as rand from "@cdktf/provider-random";

export interface Props {
  apigw: aws.apigatewayv2Api.Apigatewayv2Api;
  apiDomainName: string;
  wsdb: aws.dynamodbTable.DynamodbTable;
  wsapi: aws.apigatewayv2Api.Apigatewayv2Api;
  wsstage: aws.apigatewayv2Stage.Apigatewayv2Stage;
  uploadBucket: aws.s3Bucket.S3Bucket;
  publicBucket: aws.s3Bucket.S3Bucket;
  privateBucket: aws.s3Bucket.S3Bucket;
}

export class Backend extends Construct {
  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id);

    const rvalue = new rand.id.Id(this, "key", { byteLength: 8 });
    const eventbus = new aws.snsTopic.SnsTopic(this, "eventbus", {
      name: `eventbus-${rvalue.id}`,
    });

    // // NOT USED const coreFile = new FileStorage(this, "core-file", { eventbus, hostedZone: props.hostedZone, hostname: "files" });
    // // NOT USED new CoreSendEmail(this, "core-send-email", { eventbus });

    new CoreUser(this, "core-user", { eventbus });
    new CoreTodo(this, "core-todo", { eventbus });
    new CoreOauthUser(this, "core-oauth-user", { eventbus });
    new CoreSendEmailQueue(this, "send-email-queue", { eventbus });

    new PublicFile(this, "public-file", { apigw: props.apigw, bucket: props.uploadBucket });
    new PublicAuth(this, "public-auth", { apigw: props.apigw });
    new PublicTodo(this, "public-todo", { apigw: props.apigw });
    new PublicUser(this, "public-user", { apigw: props.apigw });

    new WorkerUser(this, "workers_user", { eventbus });
    new WorkerFile(this, "workers_file", {
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

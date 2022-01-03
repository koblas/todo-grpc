import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import { DynamoRWPolicy, SNSPublishPolicy } from "./policies";
import { buildLambda } from "./lambda-util";

export async function coreUser(corestack: pulumi.StackReference) {
  const db = new aws.dynamodb.Table("app-user", {
    name: "app-user",
    billingMode: "PAY_PER_REQUEST",

    attributes: [{ name: "pk", type: "S" }],
    hashKey: "pk",
  });

  const name = "core-user";

  const { policy: dbPolicy } = new DynamoRWPolicy(`${name}-dynamodb`, {
    dbArn: db.arn,
  });
  const { policy: snsPublish } = new SNSPublishPolicy(`${name}-sns`, {
    topicArn: corestack.getOutputValue("entityTopicArn"),
  });

  const { lambda, role } = await buildLambda("core-user", {
    code: new pulumi.asset.AssetArchive({
      app: new pulumi.asset.FileAsset("../build/core-user"),
    }),
    policies: [dbPolicy, snsPublish],
  });

  return { lambda, role, db };
}

// const dbPolicy = new aws.iam.Policy(`${name}-dynamodb`, {
//   name: `${name}-dynamodb`,
//   policy: {
//     Version: "2012-10-17",
//     Statement: [
//       {
//         Effect: "Allow",
//         Action: [
//           "dynamodb:BatchGetItem",
//           "dynamodb:GetItem",
//           "dynamodb:Query",
//           "dynamodb:Scan",
//           "dynamodb:BatchWriteItem",
//           "dynamodb:PutItem",
//           "dynamodb:UpdateItem",
//         ],
//         Resource: db.arn.apply((v) => [v, `${v}/*`, `${v}/index/*`]),
//       },
//     ],
//   },
// });

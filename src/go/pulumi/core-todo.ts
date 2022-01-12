import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import { buildLambda } from "./lambda-util";
import { DynamoRWPolicy, SNSPublishPolicy } from "./policies";

export async function coreTodo(corestack: pulumi.StackReference) {
  const name = "core-todo";

  const db = new aws.dynamodb.Table("app-todo", {
    name: "app-todo",
    billingMode: "PAY_PER_REQUEST",

    attributes: [
      { name: "pk", type: "S" },
      { name: "sk", type: "S" },
    ],

    hashKey: "pk",
    rangeKey: "sk",
  });

  const { policy: dbPolicy } = new DynamoRWPolicy(`${name}-dynamodb`, {
    dbArn: db.arn,
  });
  const { policy: snsPublish } = new SNSPublishPolicy(`${name}-sns`, {
    topicArn: corestack.getOutputValue("entityTopicArn"),
  });

  const { lambda, role } = await buildLambda("core-todo", {
    code: new pulumi.asset.AssetArchive({
      app: new pulumi.asset.FileAsset("../build/core-todo"),
    }),
    policies: [dbPolicy, snsPublish],
  });

  return { lambda, role, db };
}

// const policies: Record<string, aws.iam.PolicyDocument["Statement"]> = {
//   [`${name}-dynamodb`]: [
//     {
//       Effect: "Allow",
//       Action: [
//         "dynamodb:BatchGetItem",
//         "dynamodb:GetItem",
//         "dynamodb:Query",
//         "dynamodb:Scan",
//         "dynamodb:BatchWriteItem",
//         "dynamodb:PutItem",
//         "dynamodb:DeleteItem",
//         "dynamodb:UpdateItem",
//       ],
//       Resource: db.arn.apply((v) => [v, `${v}/*`, `${v}/index/*`]),
//     },
//   ],
// };

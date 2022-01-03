import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import { buildLambda } from "./lambda-util";
import { SNSPublishPolicy } from "./policies";

export async function coreOauthUser(corestack: pulumi.StackReference) {
  const name = "core-oauth-user";

  const { policy: snsPublish } = new SNSPublishPolicy(`${name}-sns`, {
    topicArn: corestack.getOutputValue("entityTopicArn"),
  });

  const { lambda, role } = await buildLambda(name, {
    code: new pulumi.asset.AssetArchive({
      app: new pulumi.asset.FileAsset("../build/core-oauth-user"),
    }),
    policies: [snsPublish],
  });

  return lambda;
}

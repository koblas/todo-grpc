import { publicAuth } from "./public-auth";
import { publicTodo } from "./public-todo";
import { coreUser } from "./core-user";
import { coreTodo } from "./core-todo";
import { coreOauthUser } from "./core-oauth-user";
import { coreSendEmail } from "./core-send-email";
import { sqsWorkers } from "./sqs-workers";
import * as pulumi from "@pulumi/pulumi";

async function main() {
  const corestack = new pulumi.StackReference(`koblas/devops/${pulumi.runtime.getStack()}`);

  await publicAuth(corestack);
  await publicTodo(corestack);
  await coreTodo(corestack);
  await coreUser(corestack);
  await coreOauthUser(corestack);
  await coreSendEmail(corestack);
  await sqsWorkers(corestack);
}

module.exports = main();

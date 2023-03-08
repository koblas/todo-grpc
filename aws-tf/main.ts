import { App, CloudBackend, NamedCloudWorkspace } from "cdktf";
import { BaseStack } from "./lib/stack";

const app = new App();

const stack = new BaseStack(app, "aws-tf");
new CloudBackend(stack, {
  hostname: "app.terraform.io",
  organization: "snaplabs",
  workspaces: new NamedCloudWorkspace("iqvine-prod"),
});

app.synth();

import * as path from "path";
import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { GoFunction, GoFunctionProps } from "@aws-cdk/aws-lambda-go-alpha";
import { exec } from "child_process";

const LAMBDA_DEFAULTS: Partial<GoFunctionProps> = {
  logRetention: cdk.Duration.days(3).toDays(),
  insightsVersion: cdk.aws_lambda.LambdaInsightsVersion.VERSION_1_0_135_0,
  architecture: cdk.aws_lambda.Architecture.ARM_64,
  bundling: {
    goBuildFlags: [`-ldflags="-s -w"`],
  },
};

export function goFunction(
  scope: Construct,
  name: string,
  paths: string[],
  params?: Partial<GoFunctionProps>,
): GoFunction {
  const lambda = new GoFunction(scope, [...paths, "handler"].join("-"), {
    functionName: name,
    entry: path.join(__dirname, "..", "..", "cmd", "lambda", ...paths),
    ...LAMBDA_DEFAULTS,
    ...params,
  });

  return lambda;
}

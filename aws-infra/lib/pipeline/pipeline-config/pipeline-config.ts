import * as dotenv from "dotenv";

import { Account, EnvironmentConfig, Region, Stage } from "../pipeline-types/pipeline-types";

dotenv.config();

const account = process.env.ACCOUNT || process.env.CDK_DEFAULT_ACCOUNT;
const region = process.env.REGION || process.env.CDK_DEFAULT_REGION;

export const environments: Record<Stage, EnvironmentConfig> = {
  // allow developers to spin up a quick branch for a given PR they are working on e.g. pr-124
  // this is done with an npm run develop, not through the pipeline, and uses the values in .env
  [Stage.develop]: {
    env: {
      ...(account ? { account } : {}),
      ...(region ? { region } : {}),
    },
    shared: {
      appHostname: "app",
      filesHostname: "files",
      publicBucketName: "iqvine-dev-public-files",
      privateBucketName: "iqvine-dev-private-files",
      uploadBucketName: "iqvine-dev-upload-files",
    },
    stateful: {
      domainName: "iqvine.com",
    },
    stateless: {
      domainName: "iqvine.com",
      lambdaMemorySize: parseInt(process.env.LAMBDA_MEM_SIZE || "128"),
    },
    stageName: process.env.PR_NUMBER || Stage.develop,
  },
  [Stage.featureDev]: {
    env: {
      ...(account ? { account } : {}),
      ...(region ? { region } : {}),
    },
    shared: {
      appHostname: "app",
      filesHostname: "files",
      publicBucketName: "iqvine-feature-public-files",
      privateBucketName: "iqvine-feature-private-files",
      uploadBucketName: "iqvine-feature-upload-files",
    },
    stateful: {
      domainName: "iqvine.com",
    },
    stateless: {
      domainName: "iqvine.com",
      lambdaMemorySize: 128,
    },
    stageName: Stage.featureDev,
  },
  [Stage.staging]: {
    env: {
      ...(account ? { account } : {}),
      ...(region ? { region } : {}),
    },
    shared: {
      appHostname: "app",
      filesHostname: "files",
      publicBucketName: "iqvine-stage-public-files",
      privateBucketName: "iqvine-stage-private-files",
      uploadBucketName: "iqvine-stage-upload-files",
    },
    stateful: {
      domainName: "iqvine.com",
    },
    stateless: {
      domainName: "iqvine.com",
      lambdaMemorySize: 512,
    },
    stageName: Stage.staging,
  },
  [Stage.prod]: {
    env: {
      ...(account ? { account } : {}),
      ...(region ? { region } : {}),
    },
    shared: {
      appHostname: "app",
      filesHostname: "files",
      publicBucketName: "iqvine-public-files",
      privateBucketName: "iqvine-private-files",
      uploadBucketName: "iqvine-upload-files",
    },
    stateful: {
      domainName: "iqvine.com",
    },
    stateless: {
      domainName: "iqvine.com",
      lambdaMemorySize: 1024,
    },
    stageName: Stage.prod,
  },
};

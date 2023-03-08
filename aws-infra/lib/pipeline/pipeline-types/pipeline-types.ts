export interface EnvironmentConfig {
  env: {
    account?: string;
    region?: string;
  };
  stageName: string;
  shared: {
    appHostname: string;
    filesHostname: string;
    publicBucketName: string;
    privateBucketName: string;
    uploadBucketName: string;
  };
  stateful: {
    domainName: string;
  };
  stateless: {
    lambdaMemorySize: number;
    domainName: string;
  };
}

export const enum Region {
  dublin = "eu-west-1",
  london = "eu-west-2",
  frankfurt = "eu-central-1",
}

export const enum Stage {
  featureDev = "featureDev",
  staging = "staging",
  prod = "prod",
  develop = "develop",
}

export const enum Account {
  featureDev = "11111111111",
  staging = "22222222222",
  prod = "33333333333",
}

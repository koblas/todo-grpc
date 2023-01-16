module github.com/koblas/grpc-todo

go 1.19

require (
	github.com/adjust/rmq/v5 v5.0.1
	github.com/aead/chacha20poly1305 v0.0.0-20201124145622-1a5aba2a8b29
	github.com/aws/aws-cdk-go/awscdk/v2 v2.45.0
	github.com/aws/aws-lambda-go v1.28.0
	github.com/aws/aws-sdk-go v1.42.22
	github.com/aws/aws-sdk-go-v2 v1.17.3
	github.com/aws/aws-sdk-go-v2/config v1.11.0
	github.com/aws/aws-sdk-go-v2/credentials v1.6.4
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.4.4
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.8.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.10.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.14.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.29.6
	github.com/aws/aws-sdk-go-v2/service/sns v1.12.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.14.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.17.1
	github.com/aws/constructs-go/constructs/v10 v10.1.124
	github.com/aws/jsii-runtime-go v1.69.0
	github.com/aws/smithy-go v1.13.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1
	github.com/fnproject/fdk-go v0.0.21
	github.com/go-playground/validator/v10 v10.11.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jaswdr/faker v1.8.0
	github.com/json-iterator/go v1.1.12
	github.com/o1egl/paseto v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/renstrom/shortuuid v3.0.0+incompatible
	github.com/robinjoseph08/redisqueue v1.1.0
	github.com/rs/xid v1.4.0
	github.com/stretchr/testify v1.8.0
	github.com/tencentyun/scf-go-lib v0.0.0-20211123032342-f972dcd16ff6
	github.com/twitchtv/twirp v8.1.1+incompatible
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	google.golang.org/genproto v0.0.0-20211102202547-e9cf271f7f2c
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	cloud.google.com/go v0.65.0 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/aead/poly1305 v0.0.0-20180717145839-3fee0db0b635 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.3.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.11.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cncf/xds/go v0.0.0-20211011173535-cb28da3451f1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/git-chglog/git-chglog v0.0.0-20190611050339-63a4e637021f // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-redis/redis v6.15.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/goveralls v0.0.2 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nats.go v1.22.1 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df // indirect
	github.com/urfave/cli v1.20.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/image v0.3.0 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.5 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/kyokomi/emoji.v1 v1.5.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

module github.com/koblas/grpc-todo

go 1.19

require (
	github.com/PullRequestInc/go-gpt3 v1.1.13
	github.com/adjust/rmq/v5 v5.1.1
	github.com/aead/chacha20poly1305 v0.0.0-20201124145622-1a5aba2a8b29
	github.com/aws/aws-lambda-go v1.38.0
	github.com/aws/aws-sdk-go v1.44.224
	github.com/aws/aws-sdk-go-v2 v1.17.6
	github.com/aws/aws-sdk-go-v2/config v1.18.18
	github.com/aws/aws-sdk-go-v2/credentials v1.13.17
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.10.18
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.11.5
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.19.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.30.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.30.6
	github.com/aws/aws-sdk-go-v2/service/sns v1.20.5
	github.com/aws/aws-sdk-go-v2/service/sqs v1.20.5
	github.com/aws/aws-sdk-go-v2/service/ssm v1.35.6
	github.com/aws/smithy-go v1.13.5
	github.com/bufbuild/connect-go v1.5.2
	github.com/bufbuild/connect-grpchealth-go v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/envoyproxy/go-control-plane v0.11.0
	github.com/fnproject/fdk-go v0.0.29
	github.com/go-faker/faker/v4 v4.1.0
	github.com/go-playground/validator/v10 v10.11.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gojuno/minimock/v3 v3.1.2
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/json-iterator/go v1.1.12
	github.com/minio/minio-go/v7 v7.0.49
	github.com/nats-io/nats.go v1.24.0
	github.com/o1egl/paseto v1.0.0
	github.com/oklog/ulid/v2 v2.1.0
	github.com/pkg/errors v0.9.1
	github.com/renstrom/shortuuid v3.0.0+incompatible
	github.com/robinjoseph08/redisqueue v1.1.0
	github.com/rs/xid v1.4.0
	github.com/stretchr/testify v1.8.2
	github.com/tencentyun/scf-go-lib v0.0.0-20211123032342-f972dcd16ff6
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.7.0
	golang.org/x/net v0.8.0
	golang.org/x/oauth2 v0.6.0
	golang.org/x/sync v0.1.0
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/AlecAivazis/survey/v2 v2.3.6 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/aead/poly1305 v0.0.0-20180717145839-3fee0db0b635 // indirect
	github.com/andygrunwald/go-jira v1.16.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.30 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.6 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cncf/xds/go v0.0.0-20230310173818-32f1caf87195 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.9.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/git-chglog/git-chglog v0.15.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.14 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kyokomi/emoji/v2 v2.2.12 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/goveralls v0.0.11 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nats-server/v2 v2.9.11 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/trivago/tgo v1.0.7 // indirect
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df // indirect
	github.com/urfave/cli/v2 v2.25.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/image v0.6.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

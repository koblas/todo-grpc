package filestore

import (
	"context"
	"io"
	"mime"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var InvalidSchemeError = errors.New("invalid scheme")

type AwsProvider struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	expiresIn     time.Duration
}

var _ Filestore = &AwsProvider{}

func NewAwsProvider() *AwsProvider {
	return &AwsProvider{
		expiresIn: time.Second * 60,
	}
}

func (provider *AwsProvider) buildClient(ctx context.Context) error {
	if provider.client != nil && provider.presignClient != nil {
		return nil
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	clientOptions := []func(o *s3.Options){}
	provider.client = s3.NewFromConfig(cfg)
	provider.presignClient = s3.NewPresignClient(provider.client, func(opt *s3.PresignOptions) {
		opt.ClientOptions = clientOptions
	})

	return nil
}

func (provider *AwsProvider) UploadUrl(ctx context.Context, params *FilePutParams) (*FilePutResponse, error) {
	key := buildObjectKey(params)

	if err := provider.buildClient(ctx); err != nil {
		return nil, err
	}

	request, err := provider.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(params.Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(params.ContentType),
	}, s3.WithPresignExpires(provider.expiresIn))
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't get a presigned request to put %v:%v\n",
			params.Bucket, key)
	}
	logger.FromContext(ctx).With(
		zap.String("bucket", params.Bucket),
		zap.String("objectKey", key),
	).Info("generated S3 path")

	return &FilePutResponse{request.URL}, nil
}

func (provider *AwsProvider) GetFile(ctx context.Context, params *FileGetParams) (io.Reader, error) {
	u, err := url.Parse(params.Path)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "s3" {
		return nil, InvalidSchemeError
	}
	if err := provider.buildClient(ctx); err != nil {
		return nil, err
	}
	response, err := provider.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(u.Host),
		Key:    aws.String(strings.TrimPrefix(u.Path, "/")),
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch from s3")
	}

	return response.Body, nil
}

func (provider *AwsProvider) PutFile(ctx context.Context, params *FilePutParams, reader io.Reader) (*FilePutResponse, error) {
	if err := provider.buildClient(ctx); err != nil {
		return nil, err
	}

	key := buildObjectKey(params)

	var contentType *string
	if value := mime.TypeByExtension(filepath.Ext(key)); value != "" {
		contentType = aws.String(value)
	}

	_, err := provider.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(params.Bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: contentType,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch from s3")
	}

	return &FilePutResponse{
		Url: "s3://" + params.Bucket + "/" + key,
	}, nil
}

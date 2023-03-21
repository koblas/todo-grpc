package filestore

import (
	"context"
	"io"
	"mime"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MinioProvider struct {
	Client    *minio.Client
	endpoint  string
	location  string
	expiresIn time.Duration
}

var _ Filestore = &MinioProvider{}

func NewMinioProvider(endpoint string) *MinioProvider {
	return &MinioProvider{
		endpoint:  endpoint,
		location:  "us-east-1",
		expiresIn: time.Second * 60,
	}
}

func (provider *MinioProvider) BuildClient(ctx context.Context) error {
	if provider.Client != nil {
		return nil
	}

	client, err := minio.New(provider.endpoint, &minio.Options{
		Creds: credentials.NewEnvMinio(),
		// Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		// Secure: useSSL,
	})
	if err != nil {
		return err
	}
	provider.Client = client

	return nil
}

func (provider *MinioProvider) VerifyBucket(ctx context.Context, bucketName string) error {
	exists, err := provider.Client.BucketExists(ctx, bucketName)
	// If there is an error or the bucket already exists (e.g. err == nil) return
	if err != nil || exists {
		return err
	}

	return provider.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: provider.location})
}

func (provider *MinioProvider) UploadUrl(ctx context.Context, params *FilePutParams) (*FilePutResponse, error) {
	if err := provider.BuildClient(ctx); err != nil {
		return nil, err
	}

	key, id := buildObjectKey(params)

	policy := minio.NewPostPolicy()
	policy.SetBucket(params.Bucket)
	policy.SetKey(key)
	policy.SetExpires(time.Now().UTC().Add(provider.expiresIn))

	if err := provider.VerifyBucket(ctx, params.Bucket); err != nil {
		return nil, err
	}

	result, err := provider.Client.PresignedPutObject(ctx, params.Bucket, key, provider.expiresIn)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't get a presigned request to put %v:%v\n",
			params.Bucket, key)
	}
	logger.FromContext(ctx).With(
		zap.String("bucket", params.Bucket),
		zap.String("objectKey", key),
	).Info("generated S3 path")

	return &FilePutResponse{result.String(), id}, nil
}

func (provider *MinioProvider) GetFile(ctx context.Context, params *FileGetParams) (io.Reader, error) {
	u, err := url.Parse(params.Path)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "minio" && u.Scheme != "s3" {
		return nil, InvalidSchemeError
	}
	if err := provider.BuildClient(ctx); err != nil {
		return nil, err
	}
	response, err := provider.Client.GetObject(ctx, u.Host, strings.TrimPrefix(u.Path, "/"), minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch from minio")
	}

	return response, nil
}

func (provider *MinioProvider) PutFile(ctx context.Context, params *FilePutParams, reader io.Reader) (*FilePutResponse, error) {
	if err := provider.BuildClient(ctx); err != nil {
		return nil, err
	}

	key, _ := buildObjectKey(params)

	if err := provider.VerifyBucket(ctx, params.Bucket); err != nil {
		return nil, err
	}
	contentType := mime.TypeByExtension(filepath.Ext(key))

	_, err := provider.Client.PutObject(ctx, params.Bucket, key, reader, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to store on minio")
	}

	return &FilePutResponse{
		Url: "minio://" + params.Bucket + "/" + key,
	}, nil
}

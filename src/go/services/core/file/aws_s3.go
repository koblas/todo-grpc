package file

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrorNotImplemented = errors.New("not implemented")

type s3Store struct {
	bucket        string // S3 Bucket name
	domain        string // Fully qualified domain name
	prefix        string // path prefix
	client        *s3.Client
	presignClient *s3.PresignClient
	expiresIn     time.Duration
}

var _ FileStore = (*s3Store)(nil)

func NewFileS3Store(bucket string, domain string, prefix string) FileStore {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	clientOptions := []func(o *s3.Options){}

	// if domain != "" {
	// 	resolver := s3.EndpointResolverFromURL("https://" + domain)
	// 	clientOptions = append(clientOptions, func(o *s3.Options) {
	// 		o.EndpointResolver = resolver
	// 	})
	// }

	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client, func(opt *s3.PresignOptions) {
		opt.ClientOptions = clientOptions
	})

	return &s3Store{
		bucket:        bucket,
		client:        client,
		prefix:        strings.Trim(prefix, "/"),
		domain:        domain,
		presignClient: presignClient,
		expiresIn:     60 * time.Second,
	}
}

func (store *s3Store) CreateUploadUrl(ctx context.Context, userId, fileType string) (string, error) {
	parts := []string{
		store.prefix,
		userId,
		fileType,
		uuid.NewString(),
	}
	objectKey := strings.TrimPrefix(strings.Join(parts, "/"), "/")

	request, err := store.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(store.bucket),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(store.expiresIn))
	if err != nil {
		return "", errors.Wrapf(err, "Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			store.bucket, objectKey)
	}
	logger.FromContext(ctx).With(
		zap.String("bucket", store.bucket),
		zap.String("objectKey", objectKey),
	).Info("generated S3 path")

	return request.URL, nil
}

func (store *s3Store) LookupUploadUrl(ctx context.Context, url string) (*FileInfo, error) {
	parts := strings.Split(url, "/")
	if len(parts) < 4 {
		return nil, ErrorLookupNotFound
	}

	lookupUrl := url
	_, err := store.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &store.bucket,
		Key:    aws.String(lookupUrl),
	})
	if err != nil {
		var responseError *awshttp.ResponseError
		if errors.As(err, &responseError) && responseError.ResponseError.HTTPStatusCode() == http.StatusNotFound {
			return nil, ErrorLookupNotFound
		}
		return nil, err
	}

	return &FileInfo{
		Id:          parts[len(parts)-1],
		UserId:      parts[len(parts)-3],
		FileType:    parts[len(parts)-2],
		Url:         url,
		InternalUrl: fmt.Sprintf("s3://%s/%s", store.bucket, lookupUrl),
	}, nil
}

func (store *s3Store) VerifyUploadUrl(ctx context.Context, url string, query string) error {
	return ErrorNotImplemented
}

func (store *s3Store) StoreFile(ctx context.Context, path string, bytes []byte) (string, error) {
	return "", ErrorNotImplemented
}

func (store *s3Store) GetFile(ctx context.Context, path string) ([]byte, error) {
	return nil, ErrorNotImplemented
}

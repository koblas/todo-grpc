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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrorNotImplemented = errors.New("not implemented")

type s3Store struct {
	bucket        string
	prefix        string
	client        *s3.Client
	presignClient *s3.PresignClient
	expiresIn     time.Duration
}

func NewFileS3Store(bucket string, prefix string) FileStore {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	return &s3Store{
		bucket:        bucket,
		client:        client,
		prefix:        strings.TrimPrefix(prefix, "/"),
		presignClient: presignClient,
		expiresIn:     time.Duration(60 * int64(time.Second)),
	}
}

func (store *s3Store) CreateUploadUrl(ctx context.Context, userId, fileType string) (string, error) {
	parts := []string{
		store.prefix,
		userId,
		fileType,
		uuid.NewString(),
	}
	objectKey := strings.Join(parts, "/")
	request, err := store.presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(store.bucket),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = store.expiresIn
	})
	if err != nil {
		return "", errors.Wrapf(err, "Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			store.bucket, objectKey)
	}

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

func (store *s3Store) StoreFile(ctx context.Context, url, query string, bytes []byte) (string, *FileInfo, error) {
	return "", nil, ErrorNotImplemented
}

func (store *s3Store) GetFile(ctx context.Context, path string) ([]byte, error) {
	return nil, ErrorNotImplemented
}

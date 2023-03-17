package workers_file

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"io"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/disintegration/imaging"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

func init() {
	workers = append(workers, Worker{
		Stream:    "event:file_uploaded",
		GroupName: "file.fileUploaded",
		Build:     NewFileUploaded,
	})
}

type fileUploaded struct {
	WorkerConfig
}

func NewFileUploaded(config WorkerConfig) http.Handler {
	svc := &fileUploaded{WorkerConfig: config}

	_, api := corev1connect.NewFileEventbusServiceHandler(svc)
	return api
}

func (cfg *fileUploaded) FileUploaded(ctx context.Context, msg *connect.Request[corev1.FileServiceUploadEvent]) (*connect.Response[corev1.FileEventbusFileUploadedResponse], error) {
	info := msg.Msg.Info

	log := logger.FromContext(ctx).With(
		zap.String("fileType", info.FileType),
		zap.Stringp("userId", info.UserId),
		zap.String("fileUrl", info.Url),
	)
	log.Info("in uploaded event handler")

	// We only handle profile images
	if info.FileType != "profile_image.upload" || info.UserId == nil {
		return connect.NewResponse(&corev1.FileEventbusFileUploadedResponse{}), nil
	}

	fileType := strings.TrimSuffix(info.FileType, ".upload")

	postComplete := func(errMsg string) {
		msgPtr := &errMsg
		if errMsg == "" {
			msgPtr = nil
		} else {
			log.Info(errMsg)
		}
		event := corev1.FileServiceCompleteEvent{
			IdemponcyId:  ulid.Make().String(),
			ErrorMessage: msgPtr,
			Info: &corev1.FileServiceUploadInfo{
				UserId:      info.UserId,
				FileType:    fileType,
				ContentType: nil,
				Url:         info.Url,
			},
		}

		if _, err := cfg.pubsub.FileComplete(ctx, connect.NewRequest(&event)); err != nil {
			log.With(zap.Error(err)).Error("unable to send ready")
		}
	}

	buf, err := cfg.fetchFromS3(ctx, log, msg.Msg)
	if err != nil {
		postComplete("unable to get data")
		return connect.NewResponse(&corev1.FileEventbusFileUploadedResponse{}), nil
	}

	writer, err := cfg.resizeImage(ctx, log, buf)
	if err != nil {
		postComplete("unable to get data")
		return connect.NewResponse(&corev1.FileEventbusFileUploadedResponse{}), nil
	}

	avatarUrl, err := cfg.saveToS3(ctx, log, *info.UserId, fileType, writer)
	if err != nil {
		postComplete("unable to save data")
		return connect.NewResponse(&corev1.FileEventbusFileUploadedResponse{}), nil
	}

	if err := cfg.updateUser(ctx, log, *info.UserId, avatarUrl); err != nil {
		postComplete("user update failed")
		return connect.NewResponse(&corev1.FileEventbusFileUploadedResponse{}), nil
	}

	postComplete("")
	return connect.NewResponse(&corev1.FileEventbusFileUploadedResponse{}), nil
}

func (cfg *fileUploaded) fetchFromS3(ctx context.Context, log logger.Logger, msg *corev1.FileServiceUploadEvent) (bytes.Buffer, error) {
	buf := bytes.Buffer{}

	_, span := otel.Tracer("upload").Start(ctx, "fetch_s3")
	defer span.End()
	reader, err := cfg.fileService.GetFile(ctx, &filestore.FileGetParams{
		Path: msg.Info.Url,
	})
	if err != nil {
		log.With(zap.Error(err)).Error("failed to get file data")
		return buf, errors.Wrap(err, "failed to get file data")
		// postComplete("unable to get data")
		// return &corepbv1.EventbusEmpty{}, nil
	}
	if _, err := io.Copy(&buf, reader); err != nil {
		log.With(zap.Error(err)).Error("failed to copy data")
		return buf, errors.Wrap(err, "failed to copy data")
		// postComplete("unable to get data")
		// return &corepbv1.EventbusEmpty{}, nil
	}

	return buf, nil
}

func (cfg *fileUploaded) resizeImage(ctx context.Context, log logger.Logger, buf bytes.Buffer) (bytes.Buffer, error) {
	writer := bytes.Buffer{}

	_, span := otel.Tracer("upload").Start(ctx, "resize")
	defer span.End()

	data := buf.Bytes()
	log.With(zap.Int("dataLen", len(data))).Info("Processing uploaded data")

	srcImage, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.With(zap.Error(err)).Info("unable to decode")
		return buf, errors.Wrap(err, "unable to decode")
	}
	log.With(zap.String("format", format)).Info("image decoded sucessfully")
	dstImage := imaging.Resize(srcImage, 128, 128, imaging.CatmullRom)
	log.Info("image resized sucessfully")

	writer = bytes.Buffer{}
	if err := png.Encode(&writer, dstImage); err != nil {
		log.With(zap.Error(err)).Info("unable to encode")
		return buf, errors.Wrap(err, "unable to encode")
	}

	return writer, nil
}

func (cfg *fileUploaded) saveToS3(ctx context.Context, log logger.Logger, userId, fileType string, buf bytes.Buffer) (string, error) {
	_, span := otel.Tracer("upload").Start(ctx, "save_s3")
	defer span.End()

	log.With(zap.Int("writeLen", buf.Len())).Info("Uploading resized data")
	putResult, err := cfg.fileService.PutFile(ctx, &filestore.FilePutParams{
		Bucket:   cfg.config.PublicBucket,
		UserId:   userId,
		FileType: fileType,
		Suffix:   ".png",
	}, &buf)
	if err != nil {
		log.With(zap.Error(err)).Info("failed to save to s3")
		return "", errors.Wrap(err, "unable to put data")
	}

	return putResult.Url, nil
}

func (cfg *fileUploaded) updateUser(ctx context.Context, log logger.Logger, userId string, avatarUrl string) error {
	_, span := otel.Tracer("upload").Start(ctx, "post_event")
	defer span.End()

	if _, err := cfg.userService.Update(ctx, connect.NewRequest(&corev1.UserServiceUpdateRequest{
		UserId:    userId,
		AvatarUrl: &avatarUrl,
	})); err != nil {
		return errors.Wrap(err, "user update failed")
	}

	return nil
}

func (cfg *fileUploaded) FileComplete(ctx context.Context, msg *connect.Request[corev1.FileServiceCompleteEvent]) (*connect.Response[corev1.FileEventbusFileCompleteResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("fileType", msg.Msg.Info.FileType))

	log.Info("in ready handler")
	return connect.NewResponse(&corev1.FileEventbusFileCompleteResponse{}), nil
}

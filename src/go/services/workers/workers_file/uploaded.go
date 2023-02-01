package workers_file

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"io"
	"strings"

	"github.com/disintegration/imaging"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/oklog/ulid/v2"

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

func NewFileUploaded(config WorkerConfig) corepbv1.TwirpServer {
	svc := &fileUploaded{WorkerConfig: config}

	return corepbv1.NewFileEventbusServer(svc)
}

func (cfg *fileUploaded) FileUploaded(ctx context.Context, msg *corepbv1.FileServiceUploadEvent) (*corepbv1.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(
		zap.String("fileType", msg.Info.FileType),
		zap.Stringp("userId", msg.Info.UserId),
		zap.String("fileUrl", msg.Info.Url),
	)
	log.Info("in uploaded event handler")

	// We only handle profile images
	if msg.Info.FileType != "profile_image.upload" || msg.Info.UserId == nil {
		return &corepbv1.EventbusEmpty{}, nil
	}

	fileType := strings.TrimSuffix(msg.Info.FileType, ".upload")

	postComplete := func(errMsg string) {
		msgPtr := &errMsg
		if errMsg == "" {
			msgPtr = nil
		} else {
			log.Info(errMsg)
		}
		event := corepbv1.FileServiceCompleteEvent{
			IdemponcyId:  ulid.Make().String(),
			ErrorMessage: msgPtr,
			Info: &corepbv1.FileServiceUploadInfo{
				UserId:      msg.Info.UserId,
				FileType:    fileType,
				ContentType: nil,
				Url:         msg.Info.Url,
			},
		}

		if _, err := cfg.pubsub.FileComplete(ctx, &event); err != nil {
			log.Error("unable to send ready")
		}
	}

	reader, err := cfg.fileService.GetFile(ctx, &filestore.FileGetParams{
		Path: msg.Info.Url,
	})
	if err != nil {
		log.With(zap.Error(err)).Error("failed to get file data")
		postComplete("unable to get data")
		return &corepbv1.EventbusEmpty{}, nil
	}
	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, reader); err != nil {
		log.With(zap.Error(err)).Error("failed to copy data")
		postComplete("unable to get data")
		return &corepbv1.EventbusEmpty{}, nil
	}

	data := buf.Bytes()
	log.With(zap.Int("dataLen", len(data))).Info("Processing uploaded data")

	srcImage, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		postComplete("unable to decode")
		return &corepbv1.EventbusEmpty{}, nil
	}
	log.With(zap.String("format", format)).Info("image decoded sucessfully")
	dstImage := imaging.Resize(srcImage, 128, 128, imaging.CatmullRom)
	log.Info("image resized sucessfully")

	writer := bytes.Buffer{}
	if err := png.Encode(&writer, dstImage); err != nil {
		postComplete("unable to encode")
		return &corepbv1.EventbusEmpty{}, nil
	}

	log.With(zap.Int("writeLen", writer.Len())).Info("Uploading resized data")
	putResult, err := cfg.fileService.PutFile(ctx, &filestore.FilePutParams{
		Bucket:   cfg.config.PublicBucket,
		UserId:   *msg.Info.UserId,
		FileType: fileType,
		Suffix:   ".png",
	}, &writer)
	if err != nil {
		postComplete("unable to put data")
		return &corepbv1.EventbusEmpty{}, nil
	}

	if _, err := cfg.userService.Update(ctx, &corepbv1.UserUpdateParam{
		UserId:    *msg.Info.UserId,
		AvatarUrl: &putResult.Url,
	}); err != nil {
		postComplete("user update failed")
		return &corepbv1.EventbusEmpty{}, nil
	}

	postComplete("")
	return &corepbv1.EventbusEmpty{}, nil
}

func (cfg *fileUploaded) FileComplete(ctx context.Context, msg *corepbv1.FileServiceCompleteEvent) (*corepbv1.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.String("fileType", msg.Info.FileType))

	log.Info("in ready handler")
	return &corepbv1.EventbusEmpty{}, nil
}

package workers_file

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

func init() {
	workers = append(workers, Worker{
		Stream:    "event:file_uploaded",
		GroupName: "file/fileUploaded",
		Build:     NewFileUploaded,
	})
}

type fileUploaded struct {
	WorkerConfig
}

func NewFileUploaded(config WorkerConfig) corepb.TwirpServer {
	svc := &fileUploaded{WorkerConfig: config}

	return corepb.NewFileEventbusServer(svc)
}

func (cfg *fileUploaded) FileUploaded(ctx context.Context, msg *corepb.FileUploadEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(
		zap.String("fileType", msg.Info.FileType),
		zap.Stringp("userId", msg.Info.UserId),
	)
	log.Info("in uploaded event handler")

	// We only handle profile images
	if msg.Info.FileType != "profile_image:upload" || msg.Info.UserId == nil {
		return &corepb.EventbusEmpty{}, nil
	}

	postComplete := func(errMsg string) {
		msgPtr := &errMsg
		if errMsg == "" {
			msgPtr = nil
		} else {
			log.Info(errMsg)
		}
		event := corepb.FileCompleteEvent{
			IdemponcyId:  uuid.NewString(),
			ErrorMessage: msgPtr,
			Info: &corepb.FileUploadInfo{
				UserId:      msg.Info.UserId,
				FileType:    strings.TrimSuffix(msg.Info.FileType, ":upload"),
				ContentType: nil,
				Url:         msg.Info.Url,
			},
		}

		if _, err := cfg.pubsub.FileComplete(ctx, &event); err != nil {
			log.Error("unable to send ready")
		}
	}

	fileData, err := cfg.fileService.Get(ctx, &corepb.FileGetParams{
		Path: msg.Info.Url,
	})
	if err != nil {
		log.With(zap.Error(err)).Error("unable to get image")
		return &corepb.EventbusEmpty{}, nil
	}

	srcImage, _, err := image.Decode(bytes.NewReader(fileData.Data))
	if err != nil {
		postComplete("unable to decode")
		return &corepb.EventbusEmpty{}, nil
	}
	dstImage := imaging.Resize(srcImage, 128, 128, imaging.CatmullRom)

	outData := []byte{}
	writer := bytes.NewBuffer(outData)
	if err := png.Encode(writer, dstImage); err != nil {
		postComplete("unable to encode")
		return &corepb.EventbusEmpty{}, nil
	}

	putResult, err := cfg.fileService.Put(ctx, &corepb.FilePutParams{
		UserId:   *msg.Info.UserId,
		FileType: "profile_image",
		Suffix:   ".png",
		Data:     writer.Bytes(),
	})
	if err != nil {
		postComplete("unable to save")
		return &corepb.EventbusEmpty{}, nil
	}

	if _, err := cfg.userService.Update(ctx, &corepb.UserUpdateParam{
		UserId:    *msg.Info.UserId,
		AvatarUrl: &putResult.Path,
	}); err != nil {
		postComplete("user update failed")
		return &corepb.EventbusEmpty{}, nil
	}

	postComplete("")
	return &corepb.EventbusEmpty{}, nil
}

func (cfg *fileUploaded) FileComplete(ctx context.Context, msg *corepb.FileCompleteEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.String("fileType", msg.Info.FileType))

	log.Info("in ready handler")
	return &corepb.EventbusEmpty{}, nil
}

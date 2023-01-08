package workers_file

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
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

func NewFileUploaded(config WorkerConfig) genpb.TwirpServer {
	svc := &fileUploaded{WorkerConfig: config}

	return genpb.NewFileEventbusServer(svc)
}

func (cfg *fileUploaded) FileUploaded(ctx context.Context, msg *genpb.FileUploadEvent) (*genpb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(
		zap.String("fileType", msg.Info.FileType),
		zap.Stringp("userId", msg.Info.UserId),
	)
	log.Info("in uploaded event handler")

	// We only handle profile images
	if msg.Info.FileType != "profile_image:upload" || msg.Info.UserId == nil {
		return &genpb.EventbusEmpty{}, nil
	}

	postComplete := func(errMsg string) {
		msgPtr := &errMsg
		if errMsg == "" {
			msgPtr = nil
		} else {
			log.Info(errMsg)
		}
		event := genpb.FileCompleteEvent{
			IdemponcyId:  uuid.NewString(),
			ErrorMessage: msgPtr,
			Info: &genpb.FileUploadInfo{
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

	fileData, err := cfg.fileService.Get(ctx, &genpb.FileGetParams{
		Path: msg.Info.Url,
	})
	if err != nil {
		log.With(zap.Error(err)).Error("unable to get image")
		return &genpb.EventbusEmpty{}, nil
	}

	srcImage, _, err := image.Decode(bytes.NewReader(fileData.Data))
	if err != nil {
		postComplete("unable to decode")
		return &genpb.EventbusEmpty{}, nil
	}
	dstImage := imaging.Resize(srcImage, 128, 128, imaging.CatmullRom)

	outData := []byte{}
	writer := bytes.NewBuffer(outData)
	if err := png.Encode(writer, dstImage); err != nil {
		postComplete("unable to encode")
		return &genpb.EventbusEmpty{}, nil
	}

	putResult, err := cfg.fileService.Put(ctx, &genpb.FilePutParams{
		UserId:   *msg.Info.UserId,
		FileType: "profile_image",
		Suffix:   ".png",
		Data:     outData,
	})
	if err != nil {
		postComplete("unable to save")
		return &genpb.EventbusEmpty{}, nil
	}

	if _, err := cfg.userService.Update(ctx, &genpb.UserUpdateParam{
		UserId:    *msg.Info.UserId,
		AvatarUrl: &putResult.Path,
	}); err != nil {
		postComplete("user update failed")
		return &genpb.EventbusEmpty{}, nil
	}

	postComplete("")
	return &genpb.EventbusEmpty{}, nil
}

func (cfg *fileUploaded) FileComplete(ctx context.Context, msg *genpb.FileCompleteEvent) (*genpb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.String("fileType", msg.Info.FileType))

	log.Info("in ready handler")
	return &genpb.EventbusEmpty{}, nil
}

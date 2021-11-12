package workers

import (
	"errors"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/robinjoseph08/redisqueue"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func extractBasic(log logger.Logger, msg *redisqueue.Message, m protoreflect.ProtoMessage) (string, error) {
	var action string

	metadata := genpb.Metadata{}
	badPayload := errors.New("bad payload")

	// Decode the metadata
	if iface, ok := msg.Values["metadata"]; !ok {
		log.Info("No meta data in payload")
		return action, badPayload
	} else if bdata, ok := iface.(string); !ok {
		log.Info("No meta data as string")
		return action, badPayload
	} else if err := proto.Unmarshal([]byte(bdata), &metadata); err != nil {
		return action, err
	} else {
		for _, item := range metadata.Metadata {
			if item.Key == "action" {
				action = item.Value
				break
			}
		}
		if action == "" {
			return action, badPayload
		}
	}

	// Decode the body
	if iface, ok := msg.Values["body"]; !ok {
		log.Info("No body in payload")
		return action, badPayload
	} else if bdata, ok := iface.(string); !ok {
		log.Info("No body as string")
		return action, badPayload
	} else if err := proto.Unmarshal([]byte(bdata), m); err != nil {
		return action, err
	}

	return action, nil
}

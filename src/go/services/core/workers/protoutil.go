package workers

import (
	"errors"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/logger"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func extractBasic(log logger.Logger, msg *eventbus.Message, m protoreflect.ProtoMessage) (string, error) {
	var action string

	action, found := msg.Attributes["action"]
	if !found {
		return "", errors.New("bad payload - missing action")
	}

	// Decode the body
	if len(msg.BodyBytes) == 0 {
		return "", errors.New("bad payload - empty body")
	}
	if err := proto.Unmarshal(msg.BodyBytes, m); err != nil {
		return action, err
	}

	return action, nil
}

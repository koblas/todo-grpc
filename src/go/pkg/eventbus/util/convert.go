package util

import (
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snstypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
	sqstypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/koblas/grpc-todo/pkg/eventbus"
)

const CONTENT_TRANSFER_ENCODING = "content-transfer-encoding"

func MessageToSns(topic string, msg eventbus.Message) sns.PublishInput {
	attr := map[string]snstypes.MessageAttributeValue{}

	strType := aws.String("String")
	for k, v := range msg.Attributes {
		attr[k] = snstypes.MessageAttributeValue{
			DataType:    strType,
			StringValue: aws.String(v),
		}
	}

	var mbody string
	if len(msg.BodyBytes) != 0 {
		attr[CONTENT_TRANSFER_ENCODING] = snstypes.MessageAttributeValue{
			StringValue: aws.String("base64"),
			DataType:    strType,
		}
		mbody = base64.StdEncoding.EncodeToString([]byte(msg.BodyBytes))
	} else {
		mbody = msg.BodyString
	}

	return sns.PublishInput{
		TopicArn:          aws.String(topic),
		MessageAttributes: attr,
		Message:           aws.String(mbody),
	}
}

func EventToSqs(item events.SQSMessage) sqstypes.Message {
	attr := map[string]string{}
	for k, v := range item.Attributes {
		attr[k] = v
	}
	mattr := map[string]sqstypes.MessageAttributeValue{}
	for k, v := range item.MessageAttributes {
		mattr[k] = sqstypes.MessageAttributeValue{
			DataType:         &v.DataType,
			StringValue:      v.StringValue,
			BinaryValue:      v.BinaryValue,
			BinaryListValues: v.BinaryListValues,
			StringListValues: v.StringListValues,
		}
	}

	return sqstypes.Message{
		Attributes:             attr,
		Body:                   &item.Body,
		MessageAttributes:      mattr,
		MessageId:              &item.MessageId,
		ReceiptHandle:          &item.ReceiptHandle,
		MD5OfBody:              &item.Md5OfBody,
		MD5OfMessageAttributes: &item.Md5OfMessageAttributes,
	}
}

func SqsToMessage(item sqstypes.Message) (eventbus.Message, error) {
	message := eventbus.Message{
		ID:         *item.MessageId,
		Attributes: map[string]string{},
	}

	body := item.Body
	for k, v := range item.MessageAttributes {
		if *v.DataType == "String" {
			message.Attributes[k] = *v.StringValue
		}
	}
	for k, v := range item.Attributes {
		message.Attributes[k] = v
	}
	// Use the decoded attributes
	if cte, found := message.Attributes[CONTENT_TRANSFER_ENCODING]; found && cte == "base64" {
		dec, err := base64.StdEncoding.DecodeString(*body)
		if err != nil {
			return eventbus.Message{}, err
		}

		message.BodyBytes = dec
	} else {
		message.BodyString = *body
	}

	return message, nil
}

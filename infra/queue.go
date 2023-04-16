package infra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSAPI interface {
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
}

type queue struct {
	api SQSAPI
}

func (sqs *queue) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	return sqs.api.ReceiveMessage(ctx, params, optFns...)
}

func NewQueueInfra(client *sqs.Client) SQSAPI {
	return &queue{
		api: client,
	}
}

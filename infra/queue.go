package infra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type QueueInfra interface {
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

type sqsInfra struct {
	api QueueInfra
}

func (sqs *sqsInfra) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	return sqs.api.ReceiveMessage(ctx, params, optFns...)
}

func (sqs *sqsInfra) DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error) {
	return sqs.api.DeleteMessage(ctx, params, optFns...)
}

func NewQueueInfra(client *sqs.Client) QueueInfra {
	return &sqsInfra{
		api: client,
	}
}

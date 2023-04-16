package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/stretchr/testify/assert"
)

func TestMetadaService(t *testing.T) {

	t.Run("It should return sqs error when there is", func(t *testing.T) {

		service := metadataService{
			queue: &sqsMock{
				ReceiveMessageFnMock: func() (*sqs.ReceiveMessageOutput, error) {
					return nil, fmt.Errorf("failed")
				},
			},
		}

		gotErr := service.Proccess(context.TODO())
		assert.NotNil(t, gotErr)
		assert.Equal(t, "failed", gotErr.Error())

	})

	t.Run("It should receive a msg from sqs and successfully persist", func(t *testing.T) {

		service := metadataService{
			database: &dynamoMock{
				PutItemFnMock: func() (*dynamodb.PutItemOutput, error) {
					return nil, nil
				},
			},
			queue: &sqsMock{ReceiveMessageFnMock: func() (*sqs.ReceiveMessageOutput, error) {
				return &sqs.ReceiveMessageOutput{
					Messages: []types.Message{
						{
							Body: aws.String(string([]byte(`{"id":"dummy","content":"dummy","createdAt":"2023-04-02T15:04:05Z07:00"}`))),
						},
					},
				}, nil
			}},
		}

		gotErr := service.Proccess(context.TODO())
		assert.Nil(t, gotErr)

	})

}

type sqsMock struct {
	ReceiveMessageFnMock func() (*sqs.ReceiveMessageOutput, error)
}

func (sqs *sqsMock) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	return sqs.ReceiveMessageFnMock()
}

type dynamoMock struct {
	PutItemFnMock func() (*dynamodb.PutItemOutput, error)
}

func (m *dynamoMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m.PutItemFnMock()
}

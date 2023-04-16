package infra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type db struct {
	api DynamoAPI
}

func (i *db) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return i.api.PutItem(ctx, params, optFns...)
}

func NewDatabase(client *dynamodb.Client) DynamoAPI {
	return &db{
		api: client,
	}
}

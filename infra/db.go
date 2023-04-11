package infra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DbInfra interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type databaseInfra struct {
	api DbInfra
}

func (db *databaseInfra) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return db.api.GetItem(ctx, params, optFns...)
}

func (db *databaseInfra) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return db.api.PutItem(ctx, params, optFns...)
}

func NewDatabase(client *dynamodb.Client) DbInfra {
	return &databaseInfra{
		api: client,
	}
}

package service

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/brbarmex/golang-unit-test-aws-sdk-v2/infra"
)

type MetadataInput struct {
	Id        string `json:"id"`
	Value     string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type metadataService struct {
	database infra.DynamoAPI
	queue    infra.SQSAPI
}

func (service *metadataService) Proccess(ctx context.Context) error {

	output, err := service.queue.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: aws.String("brbarmex-sqs"),
	})

	if err != nil {
		return err
	}

	for _, message := range output.Messages {

		var metadataInput MetadataInput
		if err := json.Unmarshal([]byte(*message.Body), &metadataInput); err != nil {
			log.Println(err.Error())
			continue
		}

		if strings.TrimSpace(metadataInput.CreatedAt) == "" {
			log.Println("the month is invalid")
			continue
		}

		item, err := attributevalue.MarshalMap(metadataInput)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		_, err = service.database.PutItem(ctx, &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String("Metadatas"),
		})

		if err != nil {
			log.Println(err.Error())
			continue
		}
	}

	return nil
}

func NewCustomerService(queue infra.SQSAPI, db infra.DynamoAPI) *metadataService {
	return &metadataService{
		queue:    queue,
		database: db,
	}
}

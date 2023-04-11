package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/brbarmex/golang-unit-test-aws-sdk-v2/infra"
)

type CustomerInput struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Contact     string `json:"createdAt"`
	Birthday    string `json:"birthday"`
}

type CustomerService struct {
	database infra.DbInfra
	queue    infra.QueueInfra
}

func (service *CustomerService) Proccess(ctx context.Context) error {

	result, err := service.queue.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: aws.String("brbarmex-sqs"),
	})

	if err != nil {
		return err
	}

	for _, message := range result.Messages {
		fmt.Println(message)
	}

	return nil
}

func NewCustomerService(queue infra.QueueInfra, db infra.DbInfra) *CustomerService {
	return &CustomerService{
		queue:    queue,
		database: db,
	}

}

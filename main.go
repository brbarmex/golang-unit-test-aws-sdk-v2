package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/brbarmex/golang-unit-test-aws-sdk-v2/infra"
	"github.com/brbarmex/golang-unit-test-aws-sdk-v2/service"
)

func main() {

	ctx := context.Background()

	awsCfg, err := configLoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	infraDynamo := infra.NewDatabase(dynamodb.NewFromConfig(awsCfg))
	infraSqs := infra.NewQueueInfra(sqs.NewFromConfig(awsCfg))
	service := service.NewCustomerService(infraSqs, infraDynamo)

	err = service.Proccess(ctx)
	fmt.Println(err)

}

func configLoadDefaultConfig(ctx context.Context) (cfg aws.Config, err error) {

	const (
		localstackAwsRegion   = "sa-east-1"
		localstackAwsEndpoint = "http://localhost:4566"
	)

	endpointLocalStackResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           localstackAwsEndpoint,
			SigningRegion: localstackAwsRegion,
		}, nil
	})

	return config.LoadDefaultConfig(ctx,
		config.WithRegion(localstackAwsRegion),
		config.WithEndpointResolverWithOptions(endpointLocalStackResolver),
	)
}

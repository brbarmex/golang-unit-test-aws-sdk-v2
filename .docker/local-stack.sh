
#!/bin/bash
echo "########### Setting up localstack profile ###########"
aws configure set aws_access_key_id dummy --profile=default
aws configure set aws_secret_access_key dummy --profile=default
aws configure set region sa-east-1 --profile=default

echo "########### Setting default profile ###########"
export AWS_DEFAULT_PROFILE=default

echo "########### Setting SQS names as env variables ###########"
export TEST_SQS=brbarmex-sqs

echo "########### Creating queues ###########"
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name $TEST_SQS

echo "########### Listing queues ###########"
aws --endpoint-url=http://localhost:4566 sqs list-queues

echo "########### Putting one message to the queue ###########"
aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url=http://localhost:4566/000000000000/$TEST_SQS \
    --message-body='{"id":"266fbb42-aee6-4402-a965-e55f5273c142","content":"266fbb42-aee6-4402-a965-e55f5273c142","createdAt":"2023-04-02T15:04:05Z07:00"}'

aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url=http://localhost:4566/000000000000/$TEST_SQS

aws dynamodb --endpoint-url=http://localhost:4566 create-table \
    --table-name Metadatas \
    --attribute-definitions \
        AttributeName=Id,AttributeType=S \
        AttributeName=Value,AttributeType=S \
    --key-schema \
        AttributeName=Id,KeyType=HASH \
        AttributeName=Value,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=100,WriteCapacityUnits=100

aws --endpoint-url=http://localhost:4566 --region=sa-east-1  dynamodb list-tables

aws dynamodb scan --endpoint-url=http://localhost:4566 --table-name Metadatas
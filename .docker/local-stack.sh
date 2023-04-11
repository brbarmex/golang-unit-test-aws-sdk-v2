
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
    --message-body={'{"id": "266fbb42-aee6-4402-a965-e55f5273c142",
                      "fullName": "Bruno Melo",
                      "phoneNumber": "(11) 99528-3697",
                      "address": "Suite 011 37845 Kub Flat, Trompmouth, WY 29493",
                      "createdAt": "2015-10-06",
                      "purchaseTransactions": [
                              {
                              "id": "96db4816-7f40-43e5-ba63-573730e73e1f",
                              "paymentType": "JCB",
                              "amount": 60.47,
                              "createdAt": "2017-07-01"
                                }
                              ]}'}

aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url=http://localhost:4566/000000000000/$TEST_SQS


aws dynamodb --endpoint-url=http://localhost:4566 create-table \
    --table-name CustomerEligibility \
    --attribute-definitions \
        AttributeName=Customer,AttributeType=S \
        AttributeName=Detail,AttributeType=S \
    --key-schema \
        AttributeName=Customer,KeyType=HASH \
        AttributeName=Detail,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=100,WriteCapacityUnits=100

aws --endpoint-url=http://localhost:4566 dynamodb put-item \
    --table-name CustomerEligibility  \
    --item \
        '{"Customer": {"S": "xiriu"}, "Detail": {"S": "Call Me Today"}, "AlbumTitle": {"S": "Somewhat Famous"}, "Awards": {"N": "1"}}'


aws --endpoint-url=http://localhost:4566 --region=sa-east-1  dynamodb list-tables

aws dynamodb scan --endpoint-url=http://localhost:4566 --table-name CustomerEligibility
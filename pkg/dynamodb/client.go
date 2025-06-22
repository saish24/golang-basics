package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var TableName = "my-table"
var DB = &DatabaseClient{}

func InitDynamoClient(ctx context.Context) {
	dynamoCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("ap-south-1"),
		config.WithRetryMaxAttempts(4),
	)
	if err != nil {
		panic(err)
	}

	DB.client = dynamodb.NewFromConfig(dynamoCfg)
}

func GetDynamoClient(ctx context.Context) *DatabaseClient {
	return DB
}

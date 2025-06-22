package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
)

type DatabaseClient struct {
	client *dynamodb.Client
}

type ReadItemInput struct {
	BaseItemInput
	ExpressionAttributeNames map[string]string
	ProjectionExpression     *string
	ConsistentRead           *bool
}

type ReadItemOutput struct {
	BaseItemOutput
	Item map[string]types.AttributeValue
}

type WriteItemInput struct {
	BaseItemInput
}

type WriteItemOutput struct {
	BaseItemOutput
}

type BaseItemInput struct {
	TableName string
	Key       map[string]interface{}
}

type BaseItemOutput struct {
	ConsumedCapacity *types.ConsumedCapacity
	ResultMetadata   middleware.Metadata
}

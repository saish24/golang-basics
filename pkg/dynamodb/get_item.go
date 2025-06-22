package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type GetItemInput struct {
	*ReadItemInput
}

type GetItemOutput struct {
	*ReadItemOutput
}

func (db *DatabaseClient) GetItem(ctx context.Context, input *GetItemInput, outputItem interface{}) (*GetItemOutput, error) {
	// Marshal Go value type to a map of AttributeValues
	key, err := attributevalue.MarshalMap(input.Key)
	if err != nil {
		return nil, err
	}

	res, err := db.client.GetItem(ctx, &dynamodb.GetItemInput{
		Key:                      key,
		TableName:                aws.String(input.TableName),
		ExpressionAttributeNames: input.ExpressionAttributeNames,
		ProjectionExpression:     input.ProjectionExpression,
		ConsistentRead:           input.ConsistentRead,
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal a slice of AttributeValues
	if err = attributevalue.UnmarshalMap(res.Item, outputItem); err != nil {
		return nil, err
	}

	return &GetItemOutput{
		ReadItemOutput: &ReadItemOutput{
			BaseItemOutput: BaseItemOutput{ConsumedCapacity: res.ConsumedCapacity},
			Item:           res.Item,
		},
	}, nil
}

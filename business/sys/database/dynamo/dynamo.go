package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"
)

func New() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

func PutItem(ctx context.Context, log *zap.SugaredLogger, client dynamodb.Client, data any) {

	/*	if _, ok := data.(struct{}); ok { //todo type checking

		}*/

	item, err := attributevalue.MarshalMap(data)
	if err != nil {
		//todo
	}

	_, _ = client.PutItem(ctx, &dynamodb.PutItemInput{ //todo
		Item:      item,
		TableName: aws.String(""), //todo
	})
}

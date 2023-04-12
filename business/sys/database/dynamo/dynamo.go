package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func New() *dynamodb.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return dynamodb.NewFromConfig(cfg)
}

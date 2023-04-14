package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/torenken/data-puddle-crm/app/handlers/agreementgrp"
	"github.com/torenken/data-puddle-crm/business/core/agreement"
	"github.com/torenken/data-puddle-crm/business/core/agreement/stores/agreementdb"
	"github.com/torenken/data-puddle-crm/business/sys/database/dynamo"
	"github.com/torenken/data-puddle-crm/foundation/logger"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
	db  *dynamodb.Client
)

func init() {
	log, _ = logger.New("create-agreement")
	db = dynamo.New()
}

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	handlers := agreementgrp.New(agreement.NewCore(agreementdb.NewStore(log, db)))
	return handlers.Create(ctx, req)
}

func main() {
	lambda.Start(handle)
}

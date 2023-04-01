package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/torenken/data-puddle-mock-env/app/handlers/agreementgrp"
	"github.com/torenken/data-puddle-mock-env/business/core/agreement"
	"github.com/torenken/data-puddle-mock-env/business/core/agreement/stores/agreementbucket"
	"github.com/torenken/data-puddle-mock-env/business/sys/storage/s3"
	"github.com/torenken/data-puddle-mock-env/foundation/logger"
)

/*
todo requirements
- wrapper for middle-methods like ServiceLogger
*/

/*
handle
- initiation of lambda env os.Getenv()
- init logger
-
*/

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//init
	s3c := storage.NewS3()
	log, _ := logger.New("provide-data-url")
	agtCore := agreement.NewCore(agreementbucket.NewStore(log, s3c))

	log.Info("Lets rock")

	agh := agreementgrp.Handlers{
		Agreement: agtCore,
	}
	agh.Create(ctx, req)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}

func main() {
	lambda.Start(handle)
}

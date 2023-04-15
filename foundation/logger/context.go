package logger

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
)

func AddRequestId(cxt context.Context, log *zap.SugaredLogger) *zap.SugaredLogger {
	lambdaCtx, ok := lambdacontext.FromContext(cxt)
	if ok {
		return log.With(zap.String("requestId", lambdaCtx.AwsRequestID))
	}
	return log
}

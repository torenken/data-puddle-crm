package agreementgrp

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

type LoggingHandler struct {
	next Handler
	log  *zap.SugaredLogger
}

func NewLoggingHandler(next Handler, log *zap.SugaredLogger) Handler {
	return &LoggingHandler{
		next: next,
		log:  log,
	}
}

func (h LoggingHandler) Create(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	h.log.Infof("request started: method[%s]: path[%s]", req.HTTPMethod, req.Path)

	defer func(start time.Time) {
		h.log.Infof("request completed: path[%s]: statuscode[%d]: since[%v]: method[%s]",
			req.Path, res.StatusCode, time.Since(start), req.HTTPMethod)
	}(time.Now())

	return h.next.Create(ctx, req)
}

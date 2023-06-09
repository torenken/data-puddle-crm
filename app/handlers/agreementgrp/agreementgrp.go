package agreementgrp

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/torenken/data-puddle-crm/business/core/agreement"
	"github.com/torenken/data-puddle-crm/foundation/event"
	"github.com/torenken/data-puddle-crm/foundation/web"
)

type Handler interface {
	Create(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type AgtHandlers struct {
	agreement *agreement.Core
}

func New(core *agreement.Core) Handler {
	return AgtHandlers{
		agreement: core,
	}
}

func (h AgtHandlers) Create(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var app AppNewAgreement
	if err := event.Decode([]byte(req.Body), &app); err != nil {
		//todo return ...
	}

	na, err := toCoreNewAgreement(app)
	if err != nil {
		//todo return ...
	}

	agt, err := h.agreement.Create(ctx, na)
	if err != nil {
		//todo return ...
	}
	return web.Response(agt, http.StatusOK)
}

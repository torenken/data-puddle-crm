package agreementgrp

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/torenken/data-puddle-crm/app/handlers/event"
	"github.com/torenken/data-puddle-crm/business/core/agreement"
)

type Handlers struct {
	agreement *agreement.Core
}

func New(core *agreement.Core) Handlers {
	return Handlers{
		agreement: core,
	}
}

func (h Handlers) Create(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var app AppNewAgreement
	if err := event.Decode([]byte(req.Body), &app); err != nil {
		return event.FailedResponse(err)
	}

	na, err := toCoreNewAgreement(app)
	if err != nil {
		//todo return ...
	}

	agt, err := h.agreement.Create(ctx, na)
	if err != nil {
		//todo return ...
	}
	return event.Response(agt, 200)
}

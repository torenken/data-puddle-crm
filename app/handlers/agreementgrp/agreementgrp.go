package agreementgrp

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/torenken/data-puddle-mock-env/business/core/agreement"
	v1Web "github.com/torenken/data-puddle-mock-env/business/web/v1"
	"github.com/torenken/data-puddle-mock-env/foundation/web"
)

type Handlers struct {
	Agreement *agreement.Core
}

// Create adds a new product to the system.
func (h *Handlers) Create(ctx context.Context, req events.APIGatewayProxyRequest) error {
	var app AppNewAgreement
	if err := web.Decode(strings.NewReader(req.Body), &app); err != nil {
		return err
	}

	na, err := toCoreNewAgreement(app)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	agt, err := h.Agreement.Create(ctx, na)
	if err != nil {
		return fmt.Errorf("create: app[%+v]: %w", app, err)
	}

	return web.Respond(ctx, agt, http.StatusCreated)
}

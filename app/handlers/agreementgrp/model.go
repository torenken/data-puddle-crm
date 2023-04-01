package agreementgrp

import (
	"time"

	"github.com/torenken/data-puddle-mock-env/business/core/agreement"
)

// AppAgreement represents an individual product.
type AppAgreement struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppProduct(core agreement.Agreement) AppAgreement {
	return AppAgreement{
		ID:          core.ID.String(),
		Name:        core.Name,
		DateCreated: core.DateCreated.Format(time.RFC3339),
		DateUpdated: core.DateUpdated.Format(time.RFC3339),
	}
}

// AppNewAgreement is what we require from clients when adding a Product.
type AppNewAgreement struct {
	Name string `json:"name" validate:"required"`
}

func toCoreNewAgreement(app AppNewAgreement) (agreement.NewAgreement, error) {
	core := agreement.NewAgreement{
		Name: app.Name,
	}

	return core, nil
}

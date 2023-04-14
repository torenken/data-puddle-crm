package agreementgrp

import (
	"github.com/torenken/data-puddle-crm/business/core/agreement"
	"github.com/torenken/data-puddle-crm/business/sys/validate"
)

type AppNewAgreement struct {
	Name string `json:"name" validate:"required"`
}

func toCoreNewAgreement(app AppNewAgreement) (agreement.NewAgreement, error) {
	core := agreement.NewAgreement{
		Name: app.Name,
	}
	return core, nil
}

func (app AppNewAgreement) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

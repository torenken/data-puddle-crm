package web

import (
	"encoding/json"
	"io"
)

type validator interface {
	Validate() error
}

func Decode(r io.Reader, val any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

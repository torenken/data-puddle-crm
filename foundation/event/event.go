package event

import "encoding/json"

type validator interface {
	Validate() error
}

func Decode(data []byte, val any) error {
	if err := json.Unmarshal(data, val); err != nil {
		return err
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

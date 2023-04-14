package validate

import "encoding/json"

type FieldErrors []FieldError

type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

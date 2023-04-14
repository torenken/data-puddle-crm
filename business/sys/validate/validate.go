package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

//see https://github.com/go-playground/validator

// use a single instance , it caches struct info
var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	trans, _ = ut.New(en.New(), en.New()).GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Check(val any) error {
	if err := validate.Struct(val); err != nil {
		valErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		var fields FieldErrors
		for _, valError := range valErrors {
			field := FieldError{
				Field: valError.Field(),
				Err:   valError.Translate(trans),
			}
			fields = append(fields, field)
		}
		return fields
	}
	return nil
}

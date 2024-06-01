package validation

import (
	"encoding/json"
	"errors"

	"github.com/BrunoPolaski/go-crud/src/configuration/restErr"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var transl ut.Translator

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validationError error,
) *restErr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return restErr.NewBadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []restErr.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := restErr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return restErr.NewBadRequestValidationError("One or more fields are invalid", errorCauses)
	} else {
		return restErr.NewBadRequestError("Error trying to convert fields")
	}
}

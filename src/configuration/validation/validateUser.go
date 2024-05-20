package validation

import (
	"encoding/json"
	"errors"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
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
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("One or more fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}

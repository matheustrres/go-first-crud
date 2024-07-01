package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	err error,
) *rest_errors.RestError {

	var jsonUnmarshallErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(err, &jsonUnmarshallErr) {
		return rest_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(err, &jsonValidationErr) {
		errCauses := []rest_errors.Cause{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := rest_errors.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errCauses = append(errCauses, cause)
		}

		return rest_errors.NewBadRequestValidationError("Some fields are invalid", errCauses)
	} else {
		return rest_errors.NewBadRequestError("Error trying to convert fields")
	}
}

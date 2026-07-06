package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Message  string `json:"message"`
	Tag      string `json:"tag"`
	Value    any    `json:"value"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError

	var ver validator.ValidationErrors
	if errors.As(err, &ver) {
		for _, e := range ver {
			var el ValidationError
			el.Property = e.Field()
			el.Tag = e.Tag()
			el.Value = e.Value()
			el.Message = e.Error()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}

	return nil
}
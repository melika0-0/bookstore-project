package validation

import ("github.com/go-playground/validator/v10")
type ValidationError struct {
property string `json:"property"` //what we wanna validate
message string `json:"message"`
tag string `json:"tag"`
value any `json:"value"`
}

func GetValidationErrors(err error) *[]ValidationError{
	var validationErrors []ValidationError
	var ver Validator.ValidationErrors
	if errors.As(err, &ver) {
		for_, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.property = err.field()
			el.tag = err.tag()
			el.value = err.Param()
			validationErrors = append(validationErrors, el) //append the errors

		}
		return &validationErrors //pointer to the slice
	}
	return nil //no validation errors
}
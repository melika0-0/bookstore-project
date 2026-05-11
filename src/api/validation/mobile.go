package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/melika0-0/bookstore-project/common"
)

func IranMobileNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	return common.IsValidIranMobileNumber(value)
}

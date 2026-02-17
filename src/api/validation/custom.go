package validation

import "github.com/go-playground/validator/v10"


func IrnMobileNumberValidator(fld validator.FieldLevel)bool {
	value,ok := fil.Field().Interface().(string) 
	if != ok {
		return false
	}
res , err := regexp.MatchString(`^(09\d{9}|\+989\d{9})$`, value)
if err != nil {
	log.Print(err.Error())
}
return res

}


package common

import (
	"log"
	"regexp"
)
const iranMobileNumberPattern = `^(09\d{9}|\+989\d{9})$`

func IsValidIranMobileNumber(mobile string) bool {
	res , err := regexp.MatchString(iranMobileNumberPattern, mobile)
	if err != nil{
		log.Print(err.Error())
	}
	return res
}
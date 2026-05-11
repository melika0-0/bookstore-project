package helper

import (
	"github.com/melika0-0/bookstore-project/api/validation"

)

type BaseHttpResponse struct {
	Success bool    `json:"success"`
	Result any `json:"result"`
	ResultCode int `json:"resultCode"`
	ValidationErrors  *[]validation.ValidationError `json:validationErrors`
	Error any `json:"error"`


}
// where did the error actually happend

func GetBaseHttpResponse(result any,success bool,resultcode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,Result: result,ResultCode: resultcode,
	}

}
func GetBaseHttpResponseWithError(result any,success bool,resultcode int,err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Result: result,
		ResultCode: resultcode,
		Error: err.Error(),
	}
}

func GetBaseHttpResponseWithValidationErrors(result any,success bool,resultcode int,err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Result: result,
		ResultCode: resultcode,
		Error: "validation error",
		ValidationErrors: validation.GetValidationErrors(err), //give err to validation package and get the validation errorsss
	}
}
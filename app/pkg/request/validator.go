package request

import (
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

// GetErrorMsg 获取字段验证错误信息
func GetErrorMsg(request Validator, err error) string {
	if validatorErrs, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		for _, v := range validatorErrs {
			if message, exist := request.GetMessages()[v.Field()+"."+v.Tag()]; exist {
				return message
			}
		}

	}

	return "Parameter error"
}

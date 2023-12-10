package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func AppendError(existErr, newErr error) error {
	if existErr == nil {
		return newErr
	}

	return fmt.Errorf("%v, %w", existErr, newErr)
}
func ParseValidateErrors(errs validator.ValidationErrors, target interface{}) error {
	var errResult error

	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()

	for _, fieldErr := range errs {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

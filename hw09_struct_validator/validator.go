package hw09structvalidator

import (
	"errors"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	res := ""
	for _, validationError := range v {
		res = res + validationError.Field + ": " + validationError.Err.Error() + "\n"
	}

	return strings.TrimRight(res, "\n")
}

func Validate(v interface{}) error {
	validationErrors := make(ValidationErrors, 0)
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return nil
	}

	trv := rv.Type()
	for i := 0; i < trv.NumField(); i++ {
		frv := trv.Field(i)
		vTags, ok := frv.Tag.Lookup("validate")
		if !ok {
			continue
		}

		val := rv.Field(i)
		switch val.Kind() {
		case reflect.String:
			err := ValidatorStr(val.String(), frv.Name, vTags)
			if err == nil {
				continue
			}
			if !errors.As(err, &ValidationErrors{}) {
				return err
			}
			validationErrors = append(validationErrors, err.(ValidationErrors)...)
		case reflect.Int:
			err := ValidatorInt(val.Int(), frv.Name, vTags)
			if err == nil {
				continue
			}
			if !errors.As(err, &ValidationErrors{}) {
				return err
			}
			validationErrors = append(validationErrors, err.(ValidationErrors)...)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

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
	res := strings.Builder{}
	for _, validationError := range v {
		res.WriteString(validationError.Field)
		res.WriteString(": ")
		res.WriteString(validationError.Err.Error())
		res.WriteString("\n")
	}

	return strings.TrimRight(res.String(), "\n")
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors
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

		var err error
		value := rv.Field(i)
		switch value.Kind() { // nolint:exhaustive
		case reflect.String:
			err = ValidatorStr(value.String(), frv.Name, vTags)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			err = ValidatorInt(value.Int(), frv.Name, vTags)
		case reflect.Array, reflect.Slice:
			err = ValidatorArray(value.Interface(), frv.Name, vTags)
		}

		if err == nil {
			continue
		}
		if !errors.As(err, &ValidationErrors{}) {
			return err
		}
		validationErrors = append(validationErrors, err.(ValidationErrors)...) // nolint:errorlint
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

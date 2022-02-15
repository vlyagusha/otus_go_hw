package hw09structvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
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
		res = res + validationError.Field + " " + validationError.Err.Error() + "\n"
	}

	return res
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
			fieldVal := val.String()
			for _, vTag := range strings.Split(vTags, "|") {
				tagParts := strings.SplitN(vTag, ":", 2)
				tagName := tagParts[0]
				tagVal := tagParts[1]
				switch tagName {
				case "len":
					length, err := strconv.Atoi(tagVal)
					if err != nil {
						return err
					}
					err = strLenValidator(length, fieldVal)
					if err != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: frv.Name,
							Err:   err,
						})
					}
				case "regexp":
					compile, err := regexp.Compile(tagVal)
					if err != nil {
						return err
					}

					err = strRegexpValidator(compile, fieldVal)
					if err != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: frv.Name,
							Err:   err,
						})
					}
				case "in":
					err := strInValidator(strings.Split(tagVal, ","), fieldVal)
					if err != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: frv.Name,
							Err:   err,
						})
					}
				}
			}
		case reflect.Int:
			fieldVal := val.Int()
			for _, vTag := range strings.Split(vTags, "|") {
				tagParts := strings.SplitN(vTag, ":", 2)
				tagName := tagParts[0]
				tagVal := tagParts[1]

				switch tagName {
				case "min":
					minVal, err := strconv.Atoi(tagVal)
					if err != nil {
						return err
					}
					err = intMinValidator(int64(minVal), fieldVal)
					if err != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: frv.Name,
							Err:   err,
						})
					}
				case "max":
					maxVal, err := strconv.Atoi(tagVal)
					if err != nil {
						return err
					}
					err = intMaxValidator(int64(maxVal), fieldVal)
					if err != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: frv.Name,
							Err:   err,
						})
					}
				case "in":
					strs := strings.Split(tagVal, ",")
					ints := make([]int64, len(strs))
					for i, s := range strs {
						intVal, err := strconv.Atoi(s)
						if err != nil {
							return err
						}
						ints[i] = int64(intVal)
					}
					err := intInValidator(ints, fieldVal)
					if err != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: frv.Name,
							Err:   err,
						})
					}
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func strLenValidator(length int, val string) error {
	if len(val) == length {
		return nil
	}
	return errors.New("str length validation error")
}

func strRegexpValidator(compile *regexp.Regexp, val string) error {
	if compile.Match([]byte(val)) {
		return nil
	}
	return errors.New("str regexp validation error")
}

func strInValidator(values []string, val string) error {
	for _, s := range values {
		if s == val {
			return nil
		}
	}
	return errors.New("str in validation error")
}

func intMinValidator(minVal, val int64) error {
	if val >= minVal {
		return nil
	}
	return errors.New("int min validation error")
}

func intMaxValidator(maxVal, val int64) error {
	if val <= maxVal {
		return nil
	}
	return errors.New("int max validation error")
}

func intInValidator(values []int64, val int64) error {
	for _, t := range values {
		if t == val {
			return nil
		}
	}
	return errors.New("int in validation error")
}

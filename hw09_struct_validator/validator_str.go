package hw09structvalidator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrStrLen    = errors.New("str length validation error")
	ErrStrRegexp = errors.New("str regexp validation error")
	ErrStrIn     = errors.New("str in validation error")
)

func ValidatorStr(fieldVal, fieldName, vTags string) error {
	validationErrors := make(ValidationErrors, 0)
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
			if err = strLenValidator(length, fieldVal); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
					Err:   err,
				})
			}
		case "regexp":
			compile, err := regexp.Compile(tagVal)
			if err != nil {
				return err
			}
			if err = strRegexpValidator(compile, fieldVal); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
					Err:   err,
				})
			}
		case "in":
			if err := strInValidator(strings.Split(tagVal, ","), fieldVal); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
					Err:   err,
				})
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
	return ErrStrLen
}

func strRegexpValidator(compile *regexp.Regexp, val string) error {
	if compile.Match([]byte(val)) {
		return nil
	}
	return ErrStrRegexp
}

func strInValidator(values []string, val string) error {
	for _, s := range values {
		if s == val {
			return nil
		}
	}
	return ErrStrIn
}

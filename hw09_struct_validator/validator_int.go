package hw09structvalidator

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrIntMin = errors.New("int min validation error")
	ErrIntMax = errors.New("int max validation error")
	ErrIntIn  = errors.New("int in validation error")
)

func ValidatorInt(fieldVal int64, fieldName, vTags string) error {
	validationErrors := make(ValidationErrors, 0)
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
			if err = intMinValidator(int64(minVal), fieldVal); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
					Err:   err,
				})
			}
		case "max":
			maxVal, err := strconv.Atoi(tagVal)
			if err != nil {
				return err
			}
			if err = intMaxValidator(int64(maxVal), fieldVal); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
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
			if err := intInValidator(ints, fieldVal); err != nil {
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

func intMinValidator(minVal, val int64) error {
	if val >= minVal {
		return nil
	}
	return ErrIntMin
}

func intMaxValidator(maxVal, val int64) error {
	if val <= maxVal {
		return nil
	}
	return ErrIntMax
}

func intInValidator(values []int64, val int64) error {
	for _, t := range values {
		if t == val {
			return nil
		}
	}
	return ErrIntIn
}

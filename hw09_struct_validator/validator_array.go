package hw09structvalidator

import "errors"

func ValidatorArray(value interface{}, fieldName, vTags string) error {
	var validationErrors ValidationErrors

	strSlice, ok := value.([]string)
	if ok {
		for _, val := range strSlice {
			err := ValidatorStr(val, fieldName, vTags)
			if err == nil {
				continue
			}
			if !errors.As(err, &ValidationErrors{}) {
				return err
			}
			validationErrors = append(validationErrors, err.(ValidationErrors)...) // nolint:errorlint
		}
	}

	intSlice, ok := value.([]int)
	if ok {
		for _, val := range intSlice {
			err := ValidatorInt(int64(val), fieldName, vTags)
			if err == nil {
				continue
			}
			if !errors.As(err, &ValidationErrors{}) {
				return err
			}
			validationErrors = append(validationErrors, err.(ValidationErrors)...) // nolint:errorlint
		}
	}

	return validationErrors
}

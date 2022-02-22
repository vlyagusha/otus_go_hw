package hw09structvalidator

import "errors"

func ValidatorArray(value interface{}, fieldName, vTags string) error {
	var validationErrors ValidationErrors

	switch v := value.(type) {
	case []string:
		for _, val := range v {
			err := ValidatorStr(val, fieldName, vTags)
			if err == nil {
				continue
			}
			if !errors.As(err, &ValidationErrors{}) {
				return err
			}
			validationErrors = append(validationErrors, err.(ValidationErrors)...) // nolint:errorlint
		}

		return validationErrors
	case []int:
		for _, val := range v {
			err := ValidatorInt(int64(val), fieldName, vTags)
			if err == nil {
				continue
			}
			if !errors.As(err, &ValidationErrors{}) {
				return err
			}
			validationErrors = append(validationErrors, err.(ValidationErrors)...) // nolint:errorlint
		}

		return validationErrors
	default:
		return errors.New("unknown array type")
	}
}

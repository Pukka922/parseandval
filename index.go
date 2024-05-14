package parseandval

import "github.com/go-playground/validator/v10"

func ParseAndValidate[T any](parserLib parserLib) (*T, []string) {
	var errors []string
	parsed, err := parse[T](parserLib)

	if err != nil {
		errors = append(errors, "Ungültige Daten")

		return nil, errors
	}

	err = validatorInstance.Struct(parsed)

	if err == nil {
		return parsed, nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Namespace())
	}

	return nil, errors
}

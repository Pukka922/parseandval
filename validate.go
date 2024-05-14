package parseandval

import "github.com/go-playground/validator/v10"

var validatorInstance *validator.Validate

func Init() {
	validatorInstance = validator.New()
}

func RegisterValidation(name string, fn validator.Func) {
	validatorInstance.RegisterValidation(name, fn)
}

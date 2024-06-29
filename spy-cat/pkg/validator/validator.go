package validator

import gpv "github.com/go-playground/validator/v10"

// use a single instance of Validate, it caches struct info
var validate Validator

type Validator interface {
	Validate(obj any) error
}

func GetValidator() Validator {
	if validate == nil {
		validate = New()
	}

	return validate
}

type validator struct {
	validate *gpv.Validate
}

func (v *validator) Validate(obj any) error {
	return v.validate.Struct(obj)
}

func New() Validator {
	return &validator{
		validate: gpv.New(gpv.WithRequiredStructEnabled()),
	}
}

package template

type validatorTemplate struct {
}

func (c *validatorTemplate) Text() []byte {
	return []byte(`package validator

import v "github.com/go-playground/validator/v10"

func New() *v.Validate {
	return v.New()
}
`)
}

func InternalPkgValidatorTemplate() Template {
	return &validatorTemplate{}
}

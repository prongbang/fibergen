package template

type configurationEnvironmentTemplate struct {
}

func (e *configurationEnvironmentTemplate) Text() []byte {
	return []byte(`package configuration

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)
`)
}

func ConfigurationEnvironmentTemplate() Template {
	return &configurationEnvironmentTemplate{}
}

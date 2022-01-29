package template

type docsSwaggerYamlTemplate struct {
}

func (s *docsSwaggerYamlTemplate) Text() []byte {
	return []byte("")
}

func DocsSwaggerYamlTemplate() Template {
	return &docsSwaggerYamlTemplate{}
}

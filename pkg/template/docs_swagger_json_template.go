package template

type docsSwaggerTemplate struct {
}

func (s *docsSwaggerTemplate) Text() []byte {
	return []byte(`{}`)
}

func DocsSwaggerJsonTemplate() Template {
	return &docsSwaggerTemplate{}
}

package template

type makefileTemplate struct {
	Project string
}

func (m *makefileTemplate) Text() []byte {
	return []byte(`
swaggen:
	swag init -g cmd/` + m.Project + `/main.go -o docs/apispec

api-dev:
	go run cmd/` + m.Project + `/main.go -env development

api-prod:
	go run cmd/` + m.Project + `/main.go -env production
`)
}

func MakefileTemplate(project string) Template {
	return &makefileTemplate{
		Project: project,
	}
}

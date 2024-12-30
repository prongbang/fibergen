package template

type makefileTemplate struct {
	Project string
}

func (m *makefileTemplate) Text() []byte {
	return []byte(`
swaggen:
	swag init -g cmd/api/main.go -o docs/apispec

api-dev:
	go run cmd/api/main.go -env development

api-prod:
	go run cmd/api/main.go -env production
`)
}

func MakefileTemplate() Template {
	return &makefileTemplate{}
}

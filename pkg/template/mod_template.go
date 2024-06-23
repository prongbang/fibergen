package template

type modTemplate struct {
	Module string
}

func (m *modTemplate) Text() []byte {
	return []byte(`module ` + m.Module + `

require (
	github.com/casbin/casbin/v2 v2.87.1
	github.com/fsnotify/fsnotify v1.7.0
	github.com/arsmn/fiber-swagger/v2 v2.31.1
	github.com/swaggo/swag v1.7.0
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.10.0
	github.com/gofiber/fiber/v2 v2.24.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/wire v0.5.0
	github.com/innotechdevops/mgo-driver v1.0.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/pkg/errors v0.9.1
	github.com/prongbang/callx v1.2.5
	github.com/prongbang/fiber-casbinrest v1.0.5
	github.com/prongbang/fibererror v1.1.1
	github.com/spf13/viper v1.10.1
	github.com/urfave/cli/v2 v2.3.0
	go.mongodb.org/mongo-driver v1.8.2
)`)
}

func ModTemplate(module string) Template {
	return &modTemplate{
		Module: module,
	}
}

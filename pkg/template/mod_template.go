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
	github.com/go-playground/validator/v10 v10.22.1
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/goccy/go-json v0.10.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/wire v0.6.0
	github.com/innotechdevops/mgo-driver v1.0.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/pkg/errors v0.9.1
	github.com/prongbang/callx v1.3.3
	github.com/prongbang/fiber-casbinrest v1.0.5
	github.com/prongbang/fibererror v1.1.1
	github.com/prongbang/goerror v1.0.1
	github.com/prongbang/sqlxwrapper v1.0.3
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

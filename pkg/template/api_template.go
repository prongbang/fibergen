package template

type apiTemplate struct {
	Module string
}

func (a *apiTemplate) Text() []byte {
	return []byte(`package api

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"` + a.Module + `/internal/pkg/casbinx"
	"github.com/spf13/viper"
	_ "time/tzdata"
)

type API interface {
	Register()
}

type api struct {
	Router Routers
}

func (a *api) Register() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(casbinx.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "X-Platform, X-Api-Key, Authorization, Access-Control-Allow-Credentials, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
	}))
	app.Use("/swagger", swagger.Handler)

	// Routers
	a.Router.Initials(app)

	// Serve
	_ = app.Listen(fmt.Sprintf(":%d", viper.GetInt("api.port")))
}

func NewAPI(router Routers) API {
	return &api{
		Router: router,
	}
}`)
}

func ApiTemplate(module string) Template {
	return &apiTemplate{
		Module: module,
	}
}

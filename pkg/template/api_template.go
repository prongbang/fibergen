package template

type apiTemplate struct {
	Module string
}

func (a *apiTemplate) Text() []byte {
	return []byte(`package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	_ "time/tzdata"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"` + a.Module + `/configuration"
	"` + a.Module + `/internal/pkg/casbinx"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "X-Platform, X-Api-Key, Authorization, Access-Control-Allow-Credentials, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
	}))
	app.Use("/swagger", swagger.HandlerDefault)
	app.Use(casbinx.New())

	app.Hooks().OnShutdown(func() error {
		log.Info("On server shutting down")
		return nil
	})

	// Routers
	a.Router.Initials(app)

	// Serve
	go func() {
		log.Fatal(app.Listen(fmt.Sprintf(":%d", configuration.Config.API.Port)))
	}()

	wait(func(sig os.Signal) {
		log.Info("Gracefully shutting down...")
		log.Info("Waiting for all request to finish")
		err := app.Shutdown()
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Running cleanup tasks...")
		log.Info("Server was successful shutdown.")
	})
}

func wait(fn func(os.Signal)) {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	sig := <-termChan
	fn(sig)
}

func New(router Routers) API {
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

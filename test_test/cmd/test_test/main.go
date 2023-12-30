package main

import (
	"github.com/prongbang/mvp/test_test/configuration"
	_ "github.com/prongbang/mvp/test_test/docs/apispec"
	"github.com/prongbang/mvp/test_test/internal/test_test/api"
	"github.com/prongbang/mvp/test_test/internal/test_test/database"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	_ "time/tzdata"
)

// @title Test_test API
// @version 1.0
// @description This is a swagger for Test_test API
// @termsOfService https://swagger.io/terms/
// @contact.name API Support
// @contact.url https://company.com/support
// @contact.email info@company.com
// @host localhost:9001
// @BasePath /
// @securityDefinitions.apikey APIKeyAuth
// @in header
// @name X-API-KEY
// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "",
				Usage: "-env development/production",
			},
		},
		Action: func(c *cli.Context) error {
			env := c.String("env")
			if env == configuration.EnvProduction {
				configuration.Load(env)
			} else {
				configuration.Load(configuration.EnvDevelopment)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	dbDriver := database.NewDatabaseDriver()
	apis := api.CreateAPI(dbDriver)
	apis.Register()
}
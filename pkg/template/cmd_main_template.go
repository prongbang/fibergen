package template

import (
	"strings"
)

type cmdMainTemplate struct {
	Module  string
	Project string
}

func (m *cmdMainTemplate) Text() []byte {
	return []byte(`package main

import (
	"` + m.Module + `/configuration"
	_ "` + m.Module + `/docs/apispec"
	"` + m.Module + `"
	"` + m.Module + `/internal/` + m.Project + `/database"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	_ "time/tzdata"
)

// @title ` + strings.ToUpper(m.Project[0:1]) + m.Project[1:] + ` API
// @version 1.0
// @description This is a swagger for ` + strings.ToUpper(m.Project[0:1]) + m.Project[1:] + ` API
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
	apps := ` + m.Project + `.CreateApp(dbDriver)
	apps.StartAPI()
}`)
}

func CmdMainTemplate(module string, project string) Template {
	return &cmdMainTemplate{
		Module:  module,
		Project: project,
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/prongbang/fibergen/pkg/arch"
	"github.com/prongbang/fibergen/pkg/command"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/tools"
	"github.com/urfave/cli/v2"
)

type Flags struct {
	N    string
	M    string
	F    string
	CRUD string
}

func (f Flags) Project() string {
	return strings.ReplaceAll(strings.ReplaceAll(f.N, " ", "_"), "-", "_")
}

func (f Flags) Module() string {
	return fmt.Sprintf("%s/%s", f.M, f.Project())
}

func (f Flags) Feature() string {
	return strings.ReplaceAll(strings.ReplaceAll(f.F, " ", "_"), "-", "_")
}

func main() {
	flags := Flags{}

	app := &cli.App{
		Name:      "fibergen",
		Usage:     "Generate a Clean Architecture for REST API with support for the Fiber Web Framework in Golang",
		Version:   "v1.0.6",
		Compiled:  time.Now(),
		Copyright: "(c) 2023 prongbang",
		Authors: []*cli.Author{
			{
				Name:  "prongbang",
				Email: "github.com/prongbang",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "new",
				Aliases:     []string{"n"},
				Usage:       "-n project-name",
				Destination: &flags.N,
			},
			&cli.StringFlag{
				Name:        "mod",
				Aliases:     []string{"m"},
				Usage:       "-m github.com/prongbang/module-name",
				Destination: &flags.M,
			},
			&cli.StringFlag{
				Name:        "feature",
				Aliases:     []string{"f"},
				Usage:       "-f auth",
				Destination: &flags.F,
			},
			&cli.StringFlag{
				Name:        "crud",
				Usage:       "-crud auth",
				Destination: &flags.CRUD,
			},
		},
		Action: func(*cli.Context) error {
			opt := option.Options{
				Project: flags.Project(),
				Module:  flags.Module(),
				Feature: flags.Feature(),
				Crud:    flags.CRUD,
			}
			cmd := command.New()
			arc := arch.New()
			wireInstaller := tools.NewWireInstaller(cmd)
			wireRunner := tools.NewWireRunner(cmd)
			gen := genx.NewGenerator(
				filex.NewFileX(),
				tools.New(
					wireInstaller,
					tools.NewSqlcInstaller(cmd, arc),
					tools.NewDbmlInstaller(cmd, arc),
				),
				wireInstaller,
				wireRunner,
			)
			gen.GenerateAll(opt)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

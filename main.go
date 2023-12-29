package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/urfave/cli/v2"
)

type Flags struct {
	N string
	M string
	F string
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
		},
		Action: func(*cli.Context) error {
			if flags.F == "" && flags.N == "" && flags.M == "" {
				return errors.New("please use: fibergen -f feature-name")
			}

			if flags.N == "" && flags.M == "" && flags.F == "" {
				return errors.New("please use: fibergen -n project-name -m github.com/prongbang/module-name")
			}

			opt := option.Options{Project: flags.Project(), Module: flags.Module(), Feature: flags.Feature()}
			gen := genx.NewGenerator(filex.NewFileX())
			gen.GenerateAll(opt)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

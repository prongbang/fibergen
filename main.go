package main

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/creator"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/generate"
	"os"
	"strings"
	"time"

	"github.com/ettle/strcase"
	"github.com/prongbang/fibergen/pkg/arch"
	"github.com/prongbang/fibergen/pkg/command"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/tools"
	"github.com/urfave/cli/v2"
)

type Flags struct {
	ProjectName string
	ModuleName  string
	FeatureName string
	SharedName  string
	Crud        string
	Spec        string
	Driver      string
}

func (f Flags) Project() string {
	return strcase.ToKebab(strings.ReplaceAll(f.ProjectName, " ", "_"))
}

func (f Flags) Module() string {
	return fmt.Sprintf("%s/%s", f.ModuleName, strcase.ToKebab(f.Project()))
}

func (f Flags) Feature() string {
	if f.FeatureName != "" {
		return strcase.ToSnake(strings.ReplaceAll(f.FeatureName, " ", ""))
	}
	return ""
}

func (f Flags) Shared() string {
	if f.SharedName != "" {
		return strcase.ToSnake(strings.ReplaceAll(f.SharedName, " ", ""))
	}
	return ""
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
				Destination: &flags.ProjectName,
			},
			&cli.StringFlag{
				Name:        "mod",
				Aliases:     []string{"m"},
				Usage:       "-m github.com/prongbang/module-name",
				Destination: &flags.ModuleName,
			},
			&cli.StringFlag{
				Name:        "feature",
				Aliases:     []string{"f"},
				Usage:       "-f auth",
				Destination: &flags.FeatureName,
			},
			&cli.StringFlag{
				Name:        "shared",
				Aliases:     []string{"sh"},
				Usage:       "-sh auth",
				Destination: &flags.SharedName,
			},
			&cli.StringFlag{
				Name:        "spec",
				Aliases:     []string{"s"},
				Usage:       "-s auth.json",
				Destination: &flags.Spec,
			},
			&cli.StringFlag{
				Name:        "driver",
				Aliases:     []string{"d"},
				Usage:       "-d mariadb",
				Destination: &flags.Driver,
			},
		},
		Action: func(*cli.Context) error {
			opt := option.Options{
				Project: flags.Project(),
				Module:  flags.Module(),
				Feature: flags.Feature(),
				Shared:  flags.Shared(),
				Spec:    flags.Spec,
				Driver:  flags.Driver,
			}
			cmd := command.New()
			arc := arch.New()
			wireInstaller := tools.NewWireInstaller(cmd)
			wireRunner := tools.NewWireRunner(cmd)
			fileX := filex.NewFileX()
			creatorX := creator.New(fileX)
			installer := tools.New(
				wireInstaller,
				tools.NewSqlcInstaller(cmd, arc),
				tools.NewDbmlInstaller(cmd, arc),
			)
			featureBinding := generate.NewFeatureBinding(fileX)
			sharedBinding := generate.NewSharedBinding(fileX)
			projectGenerator := generate.NewProjectGenerator(fileX)
			featureGenerator := generate.NewFeatureGenerator(fileX, creatorX, installer, wireInstaller, wireRunner, featureBinding)
			sharedGenerator := generate.NewSharedGenerator(fileX, creatorX, installer, wireInstaller, wireRunner, sharedBinding)
			gen := generate.NewGenerator(projectGenerator, featureGenerator, sharedGenerator)
			return gen.Generate(opt)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("[fibergen]", err.Error())
	}
}

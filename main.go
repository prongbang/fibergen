package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/prongbang/fibergen/pkg/option"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
)

func main() {
	projectName := flag.String("new", "", "-new project-name")
	modName := flag.String("mod", "", "-mod module-name")
	featureName := flag.String("f", "", "-f featureName-name")
	flag.Parse()

	if *featureName == "" && *projectName == "" && *modName == "" {
		fmt.Println("Please use: fibergen -f featureName-name")
		return
	}

	if *projectName == "" && *modName == "" && *featureName == "" {
		fmt.Println("Please use: fibergen -new project-name -mod github.com/prongbang/name")
		return
	}

	project := strings.ReplaceAll(strings.ReplaceAll(*projectName, " ", "_"), "-", "_")
	opt := option.Options{
		Project: project,
		Module:  fmt.Sprintf("%s/%s", *modName, project),
		Feature: *featureName,
	}

	fx := filex.NewFileX()
	gen := genx.NewGenerator(fx)
	gen.GenerateAll(opt)
}

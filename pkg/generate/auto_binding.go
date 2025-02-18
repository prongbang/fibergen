package generate

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/config"
	"github.com/prongbang/fibergen/pkg/option"
	"strings"

	"github.com/ettle/strcase"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/tocase"
	"github.com/pterm/pterm"
)

func AutoBinding(fx filex.FileX, pkg option.Package) {
	changeToRoot := "../../../"
	pwd, _ := fx.Getwd()

	wirePath := ""
	if pkg.Module.AppPath == config.AppPath {
		// Binding wire
		// Change to root directory
		_ = fx.Chdir(changeToRoot)
		pwdRoot, _ := fx.Getwd()

		wirePath = "/" + pwdRoot + "/wire.go"
	} else {
		// Reset root path
		changeToRoot = ""
		wirePath = "/" + pwd + "/wire.go"
	}

	wireB := fx.ReadFile(wirePath)
	wireText := wireB
	wireImpPat1 := "//+fibergen:import wire:package"
	wireImpPat2 := "// +fibergen:import wire:package"
	wireImp := fmt.Sprintf(
		`"%s/%s/api/%s"
	%s`, pkg.Module.Module, pkg.Module.AppPath, tocase.ToLower(pkg.Name), wireImpPat1,
	)
	wireText = strings.Replace(wireText, wireImpPat1, wireImp, 1)
	wireText = strings.Replace(wireText, wireImpPat2, wireImp, 1)

	wireBuildPat1 := "//+fibergen:func wire:build"
	wireBuildPat2 := "// +fibergen:func wire:build"
	wireBuild := fmt.Sprintf(
		`%s.ProviderSet,
		%s`, tocase.ToLower(pkg.Name), wireBuildPat1,
	)
	wireText = strings.Replace(wireText, wireBuildPat1, wireBuild, 1)
	wireText = strings.Replace(wireText, wireBuildPat2, wireBuild, 1)

	spinnerBindWire, _ := pterm.DefaultSpinner.Start("Binding file wire.go")
	if err := fx.WriteFile(wirePath, []byte(wireText)); err == nil {
		spinnerBindWire.Success()
	} else {
		spinnerBindWire.Fail()
	}

	// Binding routers
	// Change to api directory
	_ = fx.Chdir(pwd)
	routerPath := "/" + pwd + "/routers.go"
	routerB := fx.ReadFile(routerPath)
	routerText := routerB
	routerImpPat1 := "//+fibergen:import routers:package"
	routerImpPat2 := "// +fibergen:import routers:package"
	routerImp := fmt.Sprintf(
		`"%s/%s/api/%s"
	%s`, pkg.Module.Module, pkg.Module.AppPath, tocase.ToLower(pkg.Name), routerImpPat1,
	)
	routerText = strings.Replace(routerText, routerImpPat1, routerImp, 1)
	routerText = strings.Replace(routerText, routerImpPat2, routerImp, 1)

	routerStructPat1 := "//+fibergen:struct routers"
	routerStructPat2 := "// +fibergen:struct routers"
	routerStruct := fmt.Sprintf(
		`%sRoute %s.Router
	%s`, tocase.UpperCamelName(pkg.Name), tocase.ToLower(pkg.Name), routerStructPat1,
	)
	routerText = strings.Replace(routerText, routerStructPat1, routerStruct, 1)
	routerText = strings.Replace(routerText, routerStructPat2, routerStruct, 1)

	routerInitPat1 := "//+fibergen:func initials"
	routerInitPat2 := "// +fibergen:func initials"
	routerInit := fmt.Sprintf(
		`r.%sRoute.Initial(app)
	%s`, tocase.UpperCamelName(pkg.Name), routerInitPat1,
	)
	routerText = strings.Replace(routerText, routerInitPat1, routerInit, 1)
	routerText = strings.Replace(routerText, routerInitPat2, routerInit, 1)

	routerNewPat1 := "//+fibergen:func new:routers"
	routerNewPat2 := "// +fibergen:func new:routers"
	routerNew := fmt.Sprintf(
		`	%sRoute %s.Router,
	%s`, strcase.ToCamel(pkg.Name), tocase.ToLower(pkg.Name), routerNewPat1,
	)
	routerText = strings.Replace(routerText, routerNewPat1, routerNew, 1)
	routerText = strings.Replace(routerText, routerNewPat2, routerNew, 1)

	routerBindPat1 := "//+fibergen:return &routers"
	routerBindPat2 := "// +fibergen:return &routers"
	routerBind := fmt.Sprintf(
		`%sRoute: %sRoute,
		%s`, tocase.UpperCamelName(pkg.Name), strcase.ToCamel(pkg.Name), routerBindPat1,
	)
	routerText = strings.Replace(routerText, routerBindPat1, routerBind, 1)
	routerText = strings.Replace(routerText, routerBindPat2, routerBind, 1)

	spinnerBindRouter, _ := pterm.DefaultSpinner.Start("Binding file routers.go")
	if err := fx.WriteFile(routerPath, []byte(routerText)); err == nil {
		spinnerBindRouter.Success()
	} else {
		spinnerBindRouter.Fail()
	}

	// Change to root directory
	if changeToRoot != "" {
		_ = fx.Chdir(changeToRoot)
	}
}

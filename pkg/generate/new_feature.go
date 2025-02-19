package generate

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/mod"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/prongbang/fibergen/pkg/tools"
)

func NewFeature(fx filex.FileX, opt option.Options, wireInstaller tools.Installer, wireRunner tools.Runner) {
	// Install library
	_ = wireInstaller.Install()

	module := mod.GetModule(fx)
	pkg := option.Package{Name: opt.Feature, Module: module}
	for filename, tmpl := range featureTemplates(pkg) {
		FeatureGenerate(fx, pkg, filename, string(tmpl))
	}
	AutoBinding(fx, pkg)

	_ = wireRunner.Run()
}

func featureTemplates(pkg option.Package) map[string][]byte {

	dsTmpl, _ := template.RenderText(template.PrototypeDatasourceTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module, Path: pkg.Module.NewAppPath()})
	hdTmpl, _ := template.RenderText(template.PrototypeHandlerTemplate, template.Project{Name: pkg.Name})
	pdTmpl, _ := template.RenderText(template.PrototypeProviderTemplate, template.Project{Name: pkg.Name})
	rpTmpl, _ := template.RenderText(template.PrototypeRepositoryTemplate, template.Project{Name: pkg.Name})
	rtTmpl, _ := template.RenderText(template.PrototypeRouterTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})
	ucTmpl, _ := template.RenderText(template.PrototypeUsecaseTemplate, template.Project{Name: pkg.Name})
	mdTmpl, _ := template.RenderText(template.PrototypeModelTemplate, template.Project{Name: pkg.Name})

	return map[string][]byte{
		"datasource.go":                dsTmpl,
		"handler.go":                   hdTmpl,
		"provider.go":                  pdTmpl,
		"repository.go":                rpTmpl,
		"router.go":                    rtTmpl,
		"usecase.go":                   ucTmpl,
		fmt.Sprintf("%s.go", pkg.Name): mdTmpl,
	}
}

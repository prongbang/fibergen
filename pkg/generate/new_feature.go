package generate

import (
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
	for filename, tmpl := range template.FeatureTemplates(pkg) {
		FeatureGenerate(fx, pkg, filename, tmpl)
	}
	AutoBinding(fx, pkg)

	_ = wireRunner.Run()
}

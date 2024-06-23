package genx

import (
	"fmt"

	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/tools"

	"github.com/prongbang/fibergen/pkg/filex"
)

// Generator is the interface
type Generator interface {
	Generate(opt option.Options)
}

type generator struct {
	Fx            filex.FileX
	Installer     tools.Installer
	WireInstaller tools.Installer
	WireRunner    tools.Runner
}

func (f *generator) Generate(opt option.Options) {
	if opt.Project != "" && opt.Module != "" {
		NewProject(f.Fx, opt)
	} else if opt.Feature != "" {
		NewFeature(f.Fx, opt, f.WireInstaller, f.WireRunner)
	} else if opt.Crud != "" {
		NewFeatureCrud(f.Fx, opt, f.Installer, f.WireRunner)
	} else {
		fmt.Println("Not Supported")
	}
}

// NewGenerator is new instance with func
func NewGenerator(fx filex.FileX, installer tools.Installer, wireInstaller tools.Installer, wireRunner tools.Runner) Generator {
	return &generator{
		Fx:            fx,
		Installer:     installer,
		WireInstaller: wireInstaller,
		WireRunner:    wireRunner,
	}
}

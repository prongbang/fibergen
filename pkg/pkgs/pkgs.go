package pkgs

import (
	"github.com/prongbang/fibergen/pkg/mod"
	"github.com/prongbang/fibergen/pkg/option"
)

// Pkg is struct
type Pkg struct {
	Imports []string
	Name    string
	Module  mod.Mod
	Spec    option.Spec
}

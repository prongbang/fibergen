package option

import (
	"github.com/prongbang/fibergen/pkg/mod"
)

// Package is struct
type Package struct {
	Name   string
	Module mod.Mod
	Spec   Spec
}

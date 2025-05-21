package generate

import "github.com/prongbang/fibergen/pkg/option"

type Binding interface {
	Bind(pkg option.Package) error
}

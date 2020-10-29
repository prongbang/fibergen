package main

import (
	"flag"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
)

func main() {
	feature := flag.String("f", "", "-f=feature-name")
	flag.Parse()

	fx := filex.NewFileX()
	gen := genx.NewGenerator(fx)
	gen.GenerateAll(*feature)
}

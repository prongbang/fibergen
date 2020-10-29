package main

import (
	"flag"
	"fmt"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
)

func main() {
	feature := flag.String("f", "", "-f feature-name")
	flag.Parse()

	if *feature == "" {
		fmt.Println("Please use: fibergen -f feature-name")
		return
	}

	fx := filex.NewFileX()
	gen := genx.NewGenerator(fx)
	gen.GenerateAll(*feature)
}

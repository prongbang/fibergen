package main

import "flag"

func main() {
	feature := flag.String("f", "", "-f=feature-name")
	flag.Parse()

	fx := NewFileX()
	gen := NewGenerator(fx)
	gen.GenerateAll(*feature)
}

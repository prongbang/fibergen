package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// current dir
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	// find module
	os.Chdir("../../../")
	root, _ := os.Getwd()

	bt, _ := ioutil.ReadFile(root + "/go.mod")
	text := string(bt)

	m := "module "
	s := strings.Index(text, m)
	e := strings.Index(text, "\n")

	mod := text[s+len(m) : e]
	fmt.Println(mod)

	// find dir path
	i := strings.LastIndex(mod, "/")
	pj := mod[i:]
	fmt.Println(pj)
	ign := "/api"
	c := strings.LastIndex(pwd, pj)
	p := pwd[c+len(pj) : len(pwd)-len(ign)]
	fmt.Println(p)
}

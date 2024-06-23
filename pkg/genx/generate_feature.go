package genx

import (
	"fmt"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/pkgs"
	"github.com/pterm/pterm"
)

func GenerateFeature(fx filex.FileX, pkg pkgs.Pkg, filename string, tmpl string) {
	spinnerGenFile, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Generate file %s", filename))
	currentDir, err := fx.Getwd()
	if err != nil {
		fmt.Println(err)
		spinnerGenFile.Fail()
		return
	}
	currentDir = fmt.Sprintf("%s/%s", currentDir, pkg.Name)
	err = fx.EnsureDir(currentDir)
	if err != nil {
		fmt.Println(err)
		spinnerGenFile.Fail()
		return
	}
	target := fmt.Sprintf("%s/%s", currentDir, filename)
	if err := fx.WriteFile(target, []byte(tmpl)); err != nil {
		spinnerGenFile.Fail()
	} else {
		spinnerGenFile.Success()
	}
}

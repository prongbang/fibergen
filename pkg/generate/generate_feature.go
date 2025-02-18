package generate

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/tocase"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/pterm/pterm"
)

func FeatureGenerate(fx filex.FileX, pkg option.Package, filename string, tmpl string) {
	spinnerGenFile, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Generate file %s", filename))
	currentDir, err := fx.Getwd()
	if err != nil {
		fmt.Println(err)
		spinnerGenFile.Fail()
		return
	}
	currentDir = fmt.Sprintf("%s/%s", currentDir, tocase.ToLower(pkg.Name))
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

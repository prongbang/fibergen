package genx

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ettle/strcase"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/mod"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/pkgs"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/prongbang/fibergen/pkg/tools"
	"github.com/prongbang/fibergen/pkg/typeof"
)

func NewFeatureCrud(fx filex.FileX, opt option.Options, installer tools.Installer, wireRunner tools.Runner) {
	// Load spec from JSON file
	jsonSpec := fx.ReadFile(opt.Spec)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonSpec), &result)
	if err != nil {
		log.Fatal("JSON format invalid:", err)
	}

	spec := option.Spec{
		Driver: opt.Driver,
	}
	alias := strings.ToLower(opt.Crud)[:1]
	queryColumns := []string{}
	insertValues := []string{}
	insertFields := []string{}
	insertQuestions := []string{}
	updateSets := []string{}
	fields := []string{}
	for key, value := range result {
		column := strcase.ToSnake(key)
		vars := strcase.ToPascal(key)
		typeValue := typeof.ValueOf(value)

		// Fields
		fields = append(fields, fmt.Sprintf("\t%s\t%s `json:\"%s\" db:\"%s\"`", vars, typeValue, column, column))

		// Query
		queryColumns = append(queryColumns, fmt.Sprintf("%s.%s", alias, column))

		// Pk
		if strings.ToUpper(key) == "ID" {
			spec.Pk = typeValue
		} else {
			// Insert
			insertValues = append(insertValues, fmt.Sprintf("\tdata.%s,\n", vars))
			insertFields = append(insertFields, column)
			insertQuestions = append(insertQuestions, "?")

			// Update
			updateSets = append(updateSets, fmt.Sprintf(`if data.%s != "" {
		set += ", %s=:%s"
		params["%s"] = data.%s
	}`, vars, column, column, column, vars))
		}
	}
	spec.QueryColumns = strings.Join(queryColumns, ", ")
	spec.InsertValues = strings.Join(insertValues, "")
	spec.InsertFields = strings.Join(insertFields, ", ")
	spec.InsertQuestions = strings.Join(insertQuestions, ", ")
	spec.UpdateSets = strings.Join(updateSets, "\n")
	spec.Fields = strings.Join(fields, "\n")

	// Install library
	if err := installer.Install(); err == nil {
		mod := mod.GetModule(fx)
		pkg := pkgs.Pkg{
			Name:   opt.Crud,
			Module: mod,
			Spec:   spec,
		}
		for filename, tmpl := range template.FeatureCrudTemplates(pkg) {
			GenerateFeature(fx, pkg, filename, tmpl)
		}
		AutoBinding(fx, pkg)

		_ = wireRunner.Run()
	}
}

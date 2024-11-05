package genx

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/prongbang/fibergen/pkg/typer"

	"github.com/ettle/strcase"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/mod"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/pkgs"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/prongbang/fibergen/pkg/tools"
)

func NewFeatureCrud(fx filex.FileX, opt option.Options, installer tools.Installer, wireRunner tools.Runner) {
	// Load spec from JSON file
	jsonSpec := fx.ReadFile(opt.Spec)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonSpec), &result)
	if err != nil {
		log.Fatal("JSON format invalid:", err)
	}

	imports := []string{}
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
	columns := []string{}
	for key, value := range result {
		snakeTag := strcase.ToSnake(key)
		camelTag := strcase.ToCamel(key)
		vars := strcase.ToPascal(key)
		typeValue := typer.Get(value)

		// Imports
		if strings.Contains(typeValue, "time.Time") {
			if len(imports) == 0 {
				imports = append(imports, `"time"`)
			} else {
				for _, imp := range imports {
					if imp != `"time"` {
						imports = append(imports, `"time"`)
					}
				}
			}
		}

		// Columns
		columns = append(columns, snakeTag)

		// Fields
		fields = append(fields, fmt.Sprintf("\t%s\t%s `json:\"%s\" db:\"%s\"`", vars, typeValue, camelTag, snakeTag))

		// Query
		queryColumns = append(queryColumns, fmt.Sprintf("%s.%s", alias, snakeTag))

		// Pk
		if strings.ToUpper(key) == "ID" {
			spec.Pk = typeValue
		} else {
			// Insert
			insertValues = append(insertValues, fmt.Sprintf("\tdata.%s,\n", vars))
			insertFields = append(insertFields, snakeTag)
			insertQuestions = append(insertQuestions, "?")

			// Update
			operate := typer.Operate(typeValue)
			operand := typer.Value(typeValue)
			updateSets = append(updateSets, fmt.Sprintf(`if data.%s %s %s {
		set += ", %s=:%s"
		params["%s"] = data.%s
	}`, vars, operate, operand, snakeTag, snakeTag, snakeTag, vars))
		}
	}
	spec.QueryColumns = strings.Join(queryColumns, ", ")
	spec.InsertValues = strings.Join(insertValues, "")
	spec.InsertFields = strings.Join(insertFields, ", ")
	spec.InsertQuestions = strings.Join(insertQuestions, ", ")
	spec.UpdateSets = strings.Join(updateSets, "\n\t")
	spec.Fields = strings.Join(fields, "\n")
	spec.Columns = columns

	// Install library
	if err := installer.Install(); err == nil {
		module := mod.GetModule(fx)
		pkg := pkgs.Pkg{
			Imports: imports,
			Name:    opt.Crud,
			Module:  module,
			Spec:    spec,
		}
		for filename, tmpl := range template.FeatureCrudTemplates(pkg) {
			GenerateFeature(fx, pkg, filename, tmpl)
		}
		AutoBinding(fx, pkg)

		_ = wireRunner.Run()
	}
}

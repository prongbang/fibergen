package generate

import (
	"encoding/json"
	"fmt"
	"github.com/prongbang/fibergen/pkg/config"
	"github.com/prongbang/fibergen/pkg/template"
	"log"
	"strings"

	"github.com/prongbang/fibergen/pkg/typer"

	"github.com/ettle/strcase"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/mod"
	"github.com/prongbang/fibergen/pkg/option"
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
		pkg := option.Package{
			Imports: imports,
			Name:    opt.Crud,
			Module:  module,
			Spec:    spec,
		}
		for filename, tmpl := range featureCrudTemplates(pkg) {
			FeatureGenerate(fx, pkg, filename, string(tmpl))
		}
		AutoBinding(fx, pkg)

		_ = wireRunner.Run()
	}
}

func featureCrudTemplates(pkg option.Package) map[string][]byte {
	appPath := pkg.Module.AppPath
	if appPath == config.AppPath {
		appPath = config.InternalPath
	}

	dsTmpl, _ := template.RenderText(template.CrudDatasourceTemplate, template.Project{Pk: pkg.Spec.Pk, Name: pkg.Name, Module: pkg.Module.Module, Path: pkg.Module.AppPath})
	hdTmpl, _ := template.RenderText(template.CrudHandlerTemplate, template.Project{Name: pkg.Name})
	pdTmpl, _ := template.RenderText(template.CrudProviderTemplate, template.Project{Name: pkg.Name})
	pmTmpl, _ := template.RenderText(template.CrudPermissionTemplate, template.Project{Name: pkg.Name})
	rpTmpl, _ := template.RenderText(template.CrudRepositoryTemplate, template.Project{Name: pkg.Name})
	rtTmpl, _ := template.RenderText(template.CrudRouterTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})
	ucTmpl, _ := template.RenderText(template.CrudUseCaseTemplate, template.Project{Name: pkg.Name})
	mdTmpl, _ := template.RenderText(template.CrudModelTemplate, template.Project{
		Imports: pkg.Imports,
		Fields:  pkg.Spec.Fields,
		Name:    pkg.Name,
	})

	return map[string][]byte{
		"datasource.go":                dsTmpl,
		"handler.go":                   hdTmpl,
		"provider.go":                  pdTmpl,
		"permission.go":                pmTmpl,
		"repository.go":                rpTmpl,
		"router.go":                    rtTmpl,
		"usecase.go":                   ucTmpl,
		fmt.Sprintf("%s.go", pkg.Name): mdTmpl,
		//"datasource.go": DataSourceCrud(
		//	pkg.Name,
		//	pkg.Module.Module,
		//	appPath,
		//	pkg.Spec.Pk,
		//	pkg.Spec.Driver,
		//	pkg.Spec.QueryColumns,
		//	pkg.Spec.InsertValues,
		//	pkg.Spec.InsertFields,
		//	pkg.Spec.InsertQuestions,
		//	pkg.Spec.UpdateSets,
		//),
		//"handler.go": HandlerCrud(
		//	pkg.Name,
		//	pkg.Module.Module,
		//	pkg.Spec.Pk,
		//),
		//"provider.go": Provider(pkg.Name),
		//"repository.go": RepositoryCrud(
		//	pkg.Name,
		//	pkg.Spec.Pk,
		//),
		//"router.go": RouterCrud(
		//	pkg.Name,
		//	pkg.Module.Module,
		//),
		//"usecase.go": UseCaseCrud(
		//	pkg.Name,
		//	pkg.Spec.Pk,
		//),
		//fmt.Sprintf("%s.go", pkg.Name): ModelCrud(
		//	pkg.Imports,
		//	pkg.Module.Module,
		//	pkg.Spec.Pk,
		//	pkg.Name,
		//	pkg.Spec.Fields,
		//	pkg.Spec.Columns,
		//),
	}
}

package generate

import (
	"encoding/json"
	"fmt"
	"github.com/ettle/strcase"
	"github.com/prongbang/fibergen/pkg/config"
	"github.com/prongbang/fibergen/pkg/creator"
	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/mod"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/prongbang/fibergen/pkg/tools"
	"github.com/prongbang/fibergen/pkg/typer"
	"strings"
)

func featureTemplates(pkg option.Package) map[string][]byte {
	dsTmpl, _ := template.RenderText(template.PrototypeDatasourceTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module, Path: pkg.Module.NewAppPath()})
	hdTmpl, _ := template.RenderText(template.PrototypeHandlerTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})
	pdTmpl, _ := template.RenderText(template.PrototypeProviderTemplate, template.Project{Name: pkg.Name})
	rpTmpl, _ := template.RenderText(template.PrototypeRepositoryTemplate, template.Project{Name: pkg.Name})
	rtTmpl, _ := template.RenderText(template.PrototypeRouterTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})
	ucTmpl, _ := template.RenderText(template.PrototypeUseCaseTemplate, template.Project{Name: pkg.Name})
	mdTmpl, _ := template.RenderText(template.PrototypeModelTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})

	return map[string][]byte{
		"datasource.go":                dsTmpl,
		"handler.go":                   hdTmpl,
		"provider.go":                  pdTmpl,
		"repository.go":                rpTmpl,
		"router.go":                    rtTmpl,
		"usecase.go":                   ucTmpl,
		fmt.Sprintf("%s.go", pkg.Name): mdTmpl,
	}
}

func featureCrudTemplates(pkg option.Package) map[string][]byte {
	appPath := pkg.Module.AppPath
	if appPath == config.AppPath {
		appPath = config.InternalPath
	}

	dsTmpl, _ := template.RenderText(template.CrudDatasourceTemplate, template.Project{Name: pkg.Name, PrimaryField: pkg.Spec.PrimaryField, Module: pkg.Module.Module, Path: appPath})
	hdTmpl, _ := template.RenderText(template.CrudHandlerTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})
	pdTmpl, _ := template.RenderText(template.CrudProviderTemplate, template.Project{Name: pkg.Name})
	pmTmpl, _ := template.RenderText(template.CrudPermissionTemplate, template.Project{Name: pkg.Name})
	rpTmpl, _ := template.RenderText(template.CrudRepositoryTemplate, template.Project{Name: pkg.Name, PrimaryField: pkg.Spec.PrimaryField})
	rtTmpl, _ := template.RenderText(template.CrudRouterTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module})
	ucTmpl, _ := template.RenderText(template.CrudUseCaseTemplate, template.Project{Name: pkg.Name, Module: pkg.Module.Module, Fields: pkg.Spec.Fields})
	mdTmpl, _ := template.RenderText(template.CrudModelTemplate, template.Project{Imports: pkg.Spec.Imports, Module: pkg.Module.Module, Fields: pkg.Spec.Fields, PrimaryField: pkg.Spec.PrimaryField, Name: pkg.Name})

	return map[string][]byte{
		"datasource.go":                dsTmpl,
		"handler.go":                   hdTmpl,
		"provider.go":                  pdTmpl,
		"permission.go":                pmTmpl,
		"repository.go":                rpTmpl,
		"router.go":                    rtTmpl,
		"usecase.go":                   ucTmpl,
		fmt.Sprintf("%s.go", pkg.Name): mdTmpl,
	}
}

type featureGenerator struct {
	FileX          filex.FileX
	Creator        creator.Creator
	Installer      tools.Installer
	WireInstaller  tools.Installer
	WireRunner     tools.Runner
	FeatureBinding Binding
}

func (f *featureGenerator) Generate(opt option.Options) error {
	if opt.Driver != "" {
		return f.generateFeatureCrud(opt)
	}
	return f.generateFeaturePrototype(opt)
}

func (f *featureGenerator) generateFeaturePrototype(opt option.Options) error {
	_ = f.WireInstaller.Install()
	pkg := option.Package{Name: opt.Feature, Module: mod.GetModule(f.FileX)}
	for filename, tmpl := range featureTemplates(pkg) {
		_ = f.Creator.Create(creator.Config{Pkg: pkg, Filename: filename, Template: tmpl})
	}
	_ = f.FeatureBinding.Bind(pkg)
	return f.WireRunner.Run()
}

func (f *featureGenerator) generateFeatureCrud(opt option.Options) error {
	spec, err := generateSpec(f.FileX, opt)
	if err != nil {
		return err
	}

	// Install library
	if err := f.Installer.Install(); err == nil {
		module := mod.GetModule(f.FileX)
		pkg := option.Package{
			Name:   opt.Feature,
			Module: module,
			Spec:   spec,
		}
		for filename, tmpl := range featureCrudTemplates(pkg) {
			_ = f.Creator.Create(creator.Config{Pkg: pkg, Filename: filename, Template: tmpl})
		}
		_ = f.FeatureBinding.Bind(pkg)
		return f.WireRunner.Run()
	}
	return nil
}

func generateSpec(fileX filex.FileX, opt option.Options) (option.Spec, error) {
	// Load spec from JSON file
	jsonSpec := fileX.ReadFile(opt.Spec)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonSpec), &result)
	if err != nil {
		return option.Spec{}, fmt.Errorf("JSON format invalid: %s", err.Error())
	}

	imports := []string{}
	spec := option.Spec{
		Driver: opt.Driver,
	}
	alias := strings.ToLower(opt.Feature)[:1]
	queryColumns := []string{}
	insertValues := []string{}
	insertFields := []string{}
	insertQuestions := []string{}
	updateSets := []string{}
	fields := []template.Field{
		{
			Name:    "CreatedBy",
			Type:    "string",
			JsonTag: "createdBy",
			DbTag:   "created_by",
			Update:  false,
			Create:  false,
		},
		{
			Name:    "UpdatedBy",
			Type:    "string",
			JsonTag: "updatedBy",
			DbTag:   "updated_by",
			Update:  false,
			Create:  false,
		},
	}
	columns := []string{}
	for key, value := range result {
		snakeTag := strcase.ToSnake(key)
		camelTag := strcase.ToCamel(key)
		vars := strcase.ToPascal(key)
		typeValue := typer.Get(value)
		isPrimaryKey := strings.ToUpper(key) == "ID"

		// Imports
		if strings.Contains(typeValue, "time.Time") {
			if len(imports) == 0 {
				imports = append(imports, "time")
			} else {
				for _, imp := range imports {
					if imp != "time" {
						imports = append(imports, "time")
					}
				}
			}
		}

		// Columns
		columns = append(columns, snakeTag)

		// Fields
		fields = append(fields, template.Field{PrimaryKey: isPrimaryKey, Name: vars, Type: typeValue, JsonTag: camelTag, DbTag: snakeTag, Update: true, Create: true})

		// Query
		queryColumns = append(queryColumns, fmt.Sprintf("%s.%s", alias, snakeTag))

		// Pk
		if isPrimaryKey {
			spec.PrimaryField = template.PrimaryField{
				Name:    "Id",
				Type:    typeValue,
				JsonTag: "id",
			}
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
	spec.Fields = fields
	spec.Columns = columns
	spec.Imports = imports
	return spec, nil
}

func NewFeatureGenerator(
	fx filex.FileX,
	creator creator.Creator,
	installer tools.Installer,
	wireInstaller tools.Installer,
	wireRunner tools.Runner,
	featureBinding Binding,
) Generator {
	return &featureGenerator{
		FileX:          fx,
		Creator:        creator,
		Installer:      installer,
		WireInstaller:  wireInstaller,
		WireRunner:     wireRunner,
		FeatureBinding: featureBinding,
	}
}

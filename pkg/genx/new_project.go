package genx

import (
	"fmt"
	"github.com/ettle/strcase"
	"github.com/prongbang/fibergen/pkg/generate"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/pterm/pterm"
)

func NewProject(fx filex.FileX, opt option.Options) {
	var err error
	opt.Project = strcase.ToKebab(opt.Project)

	spinnerGenProject, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Create project \"%s\"", opt.Project))

	currentDir, _ := fx.Getwd()

	// Create project directory
	currentDir = fmt.Sprintf("%s/%s", currentDir, opt.Project)
	err = fx.EnsureDir(currentDir)

	// Create go.mod
	err = generate.Execute(fx, fmt.Sprintf("%s/go.mod", currentDir), template.ModTemplate, template.Project{Module: opt.Module})

	// Create cmd
	cmdPath := fmt.Sprintf("%s/cmd/api", currentDir)
	err = fx.EnsureDir(cmdPath)

	// main.go
	err = generate.Execute(fx, fmt.Sprintf("%s/main.go", cmdPath), template.CmdMainTemplate, template.Project{Name: opt.Project, Module: opt.Module})

	// Create docs
	docsDir := fmt.Sprintf("%s/docs/apispec", currentDir)
	err = fx.EnsureDir(docsDir)

	// docs.go
	err = generate.Execute(fx, fmt.Sprintf("%s/docs.go", docsDir), template.DocsTemplate, template.Any{})

	//// swagger.json
	err = generate.Execute(fx, fmt.Sprintf("%s/swagger.json", docsDir), template.DocsSwaggerJSONTemplate, template.Any{})

	// swagger.yaml
	err = generate.Execute(fx, fmt.Sprintf("%s/swagger.yaml", docsDir), template.DocsSwaggerYAMLTemplate, template.Any{})

	// Create app
	appDir := fmt.Sprintf("%s/internal", currentDir)
	err = fx.EnsureDir(appDir)

	// app.go
	err = generate.Execute(fx, fmt.Sprintf("%s/app.go", appDir), template.AppTemplate, template.Project{Module: opt.Module})

	// Create api
	apiDir := fmt.Sprintf("%s/internal/api", currentDir)
	err = fx.EnsureDir(apiDir)

	// api.go
	err = generate.Execute(fx, fmt.Sprintf("%s/api.go", apiDir), template.ApiTemplate, template.Project{Module: opt.Module})

	// routers.go
	err = generate.Execute(fx, fmt.Sprintf("%s/routers.go", apiDir), template.ApiRoutersTemplate, template.Project{Module: opt.Module})

	// wire.go
	err = generate.Execute(fx, fmt.Sprintf("%s/wire.go", currentDir), template.WireTemplate, template.Project{Module: opt.Module, Name: opt.Project})

	// wire_gen.go
	err = generate.Execute(fx, fmt.Sprintf("%s/wire_gen.go", currentDir), template.WireGenTemplate, template.Project{Module: opt.Module, Name: opt.Project})

	// Create shared pkg/core
	databaseDir := fmt.Sprintf("%s/internal/database", currentDir)
	err = fx.EnsureDir(databaseDir)

	// drivers.go
	err = generate.Execute(fx, fmt.Sprintf("%s/drivers.go", databaseDir), template.DatabaseDriversTemplate, template.Any{})

	// mongodb.go
	err = generate.Execute(fx, fmt.Sprintf("%s/mongodb.go", databaseDir), template.DatabaseMongoDBTemplate, template.Any{})

	// wire.go
	err = generate.Execute(fx, fmt.Sprintf("%s/wire.go", databaseDir), template.DatabaseWireTemplate, template.Any{})

	// wire_gen.go
	err = generate.Execute(fx, fmt.Sprintf("%s/wire_gen.go", databaseDir), template.DatabaseWireGenTemplate, template.Any{})

	// Create deployments
	deploymentsDir := fmt.Sprintf("%s/deployments", currentDir)
	err = fx.EnsureDir(deploymentsDir)

	// Dockerfile
	err = generate.Execute(fx, fmt.Sprintf("%s/Dockerfile", deploymentsDir), template.DeploymentsDockerfileTemplate, template.Project{Module: opt.Module, Name: opt.Project})

	// api-prod.yml
	err = generate.Execute(fx, fmt.Sprintf("%s/api-prod.yml", deploymentsDir), template.DeploymentsAPIComposeTemplate, template.Project{Name: opt.Project})

	// Create shared pkg/core
	coreDir := fmt.Sprintf("%s/pkg/core", currentDir)
	err = fx.EnsureDir(coreDir)

	// handler.go
	err = generate.Execute(fx, fmt.Sprintf("%s/handler.go", coreDir), template.CoreHandlerTemplate, template.Any{})

	// paging.go
	err = generate.Execute(fx, fmt.Sprintf("%s/paging.go", coreDir), template.CorePagingTemplate, template.Any{})

	// params.go
	err = generate.Execute(fx, fmt.Sprintf("%s/params.go", coreDir), template.CoreParamsTemplate, template.Any{})

	// request.go
	err = generate.Execute(fx, fmt.Sprintf("%s/request.go", coreDir), template.CoreRequestTemplate, template.Any{})

	// response.go
	err = generate.Execute(fx, fmt.Sprintf("%s/response.go", coreDir), template.CoreResponseTemplate, template.Any{})

	// router.go
	err = generate.Execute(fx, fmt.Sprintf("%s/router.go", coreDir), template.CoreRouterTemplate, template.Any{})

	// jwt.go
	err = generate.Execute(fx, fmt.Sprintf("%s/jwt.go", coreDir), template.CoreJWTTemplate, template.Any{})

	// flag.go
	err = generate.Execute(fx, fmt.Sprintf("%s/flag.go", coreDir), template.CoreFlagTemplate, template.Any{})

	// sorting.go
	err = generate.Execute(fx, fmt.Sprintf("%s/sorting.go", coreDir), template.CoreSortingTemplate, template.Any{})

	// header.go
	err = generate.Execute(fx, fmt.Sprintf("%s/header.go", coreDir), template.CoreHeaderTemplate, template.Any{})

	// Create shared pkg/requestx
	requestxDir := fmt.Sprintf("%s/pkg/requestx", currentDir)
	err = fx.EnsureDir(requestxDir)

	// requestx.go
	err = generate.Execute(fx, fmt.Sprintf("%s/requestx.go", requestxDir), template.RequestXRequestTemplate, template.Project{Module: opt.Module})

	// Create shared pkg/structx
	structxDir := fmt.Sprintf("%s/pkg/structx", currentDir)
	err = fx.EnsureDir(structxDir)

	// struct.go
	err = generate.Execute(fx, fmt.Sprintf("%s/struct.go", structxDir), template.StructXTemplate, template.Any{})

	// Create shared pkg/schema
	schemaDir := fmt.Sprintf("%s/pkg/schema", currentDir)
	err = fx.EnsureDir(schemaDir)

	// sql.go
	err = generate.Execute(fx, fmt.Sprintf("%s/sql.go", schemaDir), template.SchemaSQLTemplate, template.Any{})

	// Create policy
	casbinPolicyDir := fmt.Sprintf("%s/policy", currentDir)
	err = fx.EnsureDir(casbinPolicyDir)

	// model.conf
	err = generate.Execute(fx, fmt.Sprintf("%s/model.conf", casbinPolicyDir), template.CasbinModelTemplate, template.Any{})

	// policy.csv
	err = generate.Execute(fx, fmt.Sprintf("%s/policy.csv", casbinPolicyDir), template.CasbinPolicyTemplate, template.Any{})

	// Create Makefile
	err = generate.Execute(fx, fmt.Sprintf("%s/Makefile", currentDir), template.MakefileTemplate, template.Any{})

	// Create configuration
	configurationDir := fmt.Sprintf("%s/configuration", currentDir)
	err = fx.EnsureDir(configurationDir)

	// configuration.go
	err = generate.Execute(fx, fmt.Sprintf("%s/configuration.go", configurationDir), template.ConfigurationTemplate, template.Any{})

	// environment.go
	err = generate.Execute(fx, fmt.Sprintf("%s/environment.go", configurationDir), template.ConfigurationEnvironmentTemplate, template.Any{})

	// development.yml
	err = generate.Execute(fx, fmt.Sprintf("%s/development.yml", configurationDir), template.ConfigurationDevelopmentTemplate, template.Any{})

	// production.yml
	err = generate.Execute(fx, fmt.Sprintf("%s/production.yml", configurationDir), template.ConfigurationProductionTemplate, template.Any{})

	// Create internal/pkg
	internalPkgDir := fmt.Sprintf("%s/internal/pkg", currentDir)
	err = fx.EnsureDir(internalPkgDir)

	// Create internal/pkg/casbinx
	internalPkgCasbinxDir := fmt.Sprintf("%s/internal/pkg/casbinx", currentDir)
	err = fx.EnsureDir(internalPkgCasbinxDir)

	// casbinx.go
	err = generate.Execute(fx, fmt.Sprintf("%s/casbinx.go", internalPkgCasbinxDir), template.InternalPkgCasbinxTemplate, template.Any{})

	// Create internal/pkg/response
	internalPkgResponseDir := fmt.Sprintf("%s/internal/pkg/response", currentDir)
	err = fx.EnsureDir(internalPkgResponseDir)

	// response.go
	err = generate.Execute(fx, fmt.Sprintf("%s/response.go", internalPkgResponseDir), template.InternalPkgResponseTemplate, template.Any{})

	// Create internal/pkg/validator
	internalPkgValidatorDir := fmt.Sprintf("%s/internal/pkg/validator", currentDir)
	err = fx.EnsureDir(internalPkgValidatorDir)

	// validator.go
	err = generate.Execute(fx, fmt.Sprintf("%s/validator.go", internalPkgValidatorDir), template.InternalPkgValidatorTemplate, template.Any{})

	// Update status
	if err == nil {
		spinnerGenProject.Success()
	} else {
		spinnerGenProject.Fail()
	}
}

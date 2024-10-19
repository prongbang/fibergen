package genx

import (
	"fmt"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/pterm/pterm"
)

func NewProject(fx filex.FileX, opt option.Options) {
	spinnerGenProject, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Create project \"%s\"", opt.Project))

	currentDir, _ := fx.Getwd()

	// Create project directory
	currentDir = fmt.Sprintf("%s/%s", currentDir, opt.Project)
	_ = fx.EnsureDir(currentDir)

	// Create go.mod
	modPath := fmt.Sprintf("%s/go.mod", currentDir)
	modTemplate := template.ModTemplate(opt.Module)
	_ = fx.WriteFile(modPath, modTemplate.Text())

	// Create cmd
	cmdPath := fmt.Sprintf("%s/cmd/%s", currentDir, opt.Project)
	_ = fx.EnsureDir(cmdPath)

	// main.go
	cmdMainPath := fmt.Sprintf("%s/main.go", cmdPath)
	cmdMainTemplate := template.CmdMainTemplate(opt.Module, opt.Project)
	_ = fx.WriteFile(cmdMainPath, cmdMainTemplate.Text())

	// Create docs
	docsDir := fmt.Sprintf("%s/docs/apispec", currentDir)
	_ = fx.EnsureDir(docsDir)

	// docs.go
	docsPath := fmt.Sprintf("%s/docs.go", docsDir)
	docsTemplate := template.DocsTemplate()
	_ = fx.WriteFile(docsPath, docsTemplate.Text())

	// swagger.json
	docsSwaggerPath := fmt.Sprintf("%s/swagger.json", docsDir)
	docsSwaggerTemplate := template.DocsSwaggerJsonTemplate()
	_ = fx.WriteFile(docsSwaggerPath, docsSwaggerTemplate.Text())

	// swagger.yaml
	docsSwaggerYamlPath := fmt.Sprintf("%s/swagger.yaml", docsDir)
	docsSwaggerYamlTemplate := template.DocsSwaggerYamlTemplate()
	_ = fx.WriteFile(docsSwaggerYamlPath, docsSwaggerYamlTemplate.Text())

	// Create api
	apiDir := fmt.Sprintf("%s/internal/%s/api", currentDir, opt.Project)
	_ = fx.EnsureDir(apiDir)

	// api.go
	apiPath := fmt.Sprintf("%s/api.go", apiDir)
	apiTemplate := template.ApiTemplate(opt.Module)
	_ = fx.WriteFile(apiPath, apiTemplate.Text())

	// routers.go
	apiRoutersPath := fmt.Sprintf("%s/routers.go", apiDir)
	apiRoutersTemplate := template.ApiRoutersTemplate(opt.Module)
	_ = fx.WriteFile(apiRoutersPath, apiRoutersTemplate.Text())

	// wire.go
	wireApiPath := fmt.Sprintf("%s/wire.go", apiDir)
	wireApiTemplate := template.WireApiTemplate(opt.Module, opt.Project)
	_ = fx.WriteFile(wireApiPath, wireApiTemplate.Text())

	// wire_gen.go
	wireGenApiPath := fmt.Sprintf("%s/wire_gen.go", apiDir)
	wireGenApiTemplate := template.WireGenApiTemplate(opt.Module, opt.Project)
	_ = fx.WriteFile(wireGenApiPath, wireGenApiTemplate.Text())

	// Create shared pkg/core
	databaseDir := fmt.Sprintf("%s/internal/%s/database", currentDir, opt.Project)
	_ = fx.EnsureDir(databaseDir)

	// drivers.go
	databaseDriversPath := fmt.Sprintf("%s/drivers.go", databaseDir)
	databaseDriversTemplate := template.DatabaseDriversTemplate()
	_ = fx.WriteFile(databaseDriversPath, databaseDriversTemplate.Text())

	// mongodb.go
	databaseMongodbPath := fmt.Sprintf("%s/mongodb.go", databaseDir)
	databaseMongodbTemplate := template.DatabaseMongodbTemplate()
	_ = fx.WriteFile(databaseMongodbPath, databaseMongodbTemplate.Text())

	// wire.go
	databaseWirePath := fmt.Sprintf("%s/wire.go", databaseDir)
	databaseWireTemplate := template.DatabaseWireTemplate()
	_ = fx.WriteFile(databaseWirePath, databaseWireTemplate.Text())

	// wire_gen.go
	databaseWireGenPath := fmt.Sprintf("%s/wire_gen.go", databaseDir)
	databaseWireGenTemplate := template.DatabaseWireGenTemplate()
	_ = fx.WriteFile(databaseWireGenPath, databaseWireGenTemplate.Text())

	// Create deployments
	deploymentsDir := fmt.Sprintf("%s/deployments", currentDir)
	_ = fx.EnsureDir(deploymentsDir)

	// Dockerfile
	deploymentsDockerfilePath := fmt.Sprintf("%s/Dockerfile", deploymentsDir)
	deploymentsDockerfileTemplate := template.DeploymentsDockerfileTemplate(opt.Module, opt.Project)
	_ = fx.WriteFile(deploymentsDockerfilePath, deploymentsDockerfileTemplate.Text())

	// api-prod.yml
	deploymentsApiComposePath := fmt.Sprintf("%s/api-prod.yml", deploymentsDir)
	deploymentsApiComposeTemplate := template.DeploymentsApiComposeTemplate(opt.Project)
	_ = fx.WriteFile(deploymentsApiComposePath, deploymentsApiComposeTemplate.Text())

	// Create shared pkg/core
	coreDir := fmt.Sprintf("%s/pkg/core", currentDir)
	_ = fx.EnsureDir(coreDir)

	// handler.go
	coreHandlerPath := fmt.Sprintf("%s/handler.go", coreDir)
	coreHandlerTemplate := template.CoreHandlerTemplate()
	_ = fx.WriteFile(coreHandlerPath, coreHandlerTemplate.Text())

	// paging.go
	corePagingPath := fmt.Sprintf("%s/paging.go", coreDir)
	corePagingTemplate := template.CorePagingTemplate()
	_ = fx.WriteFile(corePagingPath, corePagingTemplate.Text())

	// params.go
	coreParamsPath := fmt.Sprintf("%s/params.go", coreDir)
	coreParamsTemplate := template.CoreParamsTemplate()
	_ = fx.WriteFile(coreParamsPath, coreParamsTemplate.Text())

	// request.go
	coreRequestPath := fmt.Sprintf("%s/request.go", coreDir)
	coreRequestTemplate := template.CoreRequestTemplate()
	_ = fx.WriteFile(coreRequestPath, coreRequestTemplate.Text())

	// response.go
	coreResponsePath := fmt.Sprintf("%s/response.go", coreDir)
	coreResponseTemplate := template.CoreResponseTemplate()
	_ = fx.WriteFile(coreResponsePath, coreResponseTemplate.Text())

	// router.go
	coreRouterPath := fmt.Sprintf("%s/router.go", coreDir)
	coreRouterTemplate := template.CoreRouterTemplate()
	_ = fx.WriteFile(coreRouterPath, coreRouterTemplate.Text())

	// validate.go
	coreValidatePath := fmt.Sprintf("%s/validate.go", coreDir)
	coreValidateTemplate := template.CoreValidateTemplate()
	_ = fx.WriteFile(coreValidatePath, coreValidateTemplate.Text())

	// jwt.go
	coreJwtPath := fmt.Sprintf("%s/jwt.go", coreDir)
	coreJwtTemplate := template.CoreJwtTemplate()
	_ = fx.WriteFile(coreJwtPath, coreJwtTemplate.Text())

	// flag.go
	coreFlagPath := fmt.Sprintf("%s/flag.go", coreDir)
	coreFlagTemplate := template.CoreFlagTemplate()
	_ = fx.WriteFile(coreFlagPath, coreFlagTemplate.Text())

	// header.go
	coreHeaderPath := fmt.Sprintf("%s/header.go", coreDir)
	coreHeaderTemplate := template.CoreHeaderTemplate()
	_ = fx.WriteFile(coreHeaderPath, coreHeaderTemplate.Text())

	// Create policy
	casbinPolicyDir := fmt.Sprintf("%s/policy", currentDir)
	_ = fx.EnsureDir(casbinPolicyDir)

	// model.conf
	casbinModelPath := fmt.Sprintf("%s/model.conf", casbinPolicyDir)
	casbinModelTemplate := template.CasbinModelTemplate()
	_ = fx.WriteFile(casbinModelPath, casbinModelTemplate.Text())

	// policy.csv
	casbinPolicyPath := fmt.Sprintf("%s/policy.csv", casbinPolicyDir)
	casbinPolicyTemplate := template.CasbinPolicyTemplate()
	_ = fx.WriteFile(casbinPolicyPath, casbinPolicyTemplate.Text())

	// Create Makefile
	makefilePath := fmt.Sprintf("%s/Makefile", currentDir)
	makefileTemplate := template.MakefileTemplate(opt.Project)
	_ = fx.WriteFile(makefilePath, makefileTemplate.Text())

	// Create configuration
	configurationDir := fmt.Sprintf("%s/configuration", currentDir)
	_ = fx.EnsureDir(configurationDir)

	// configuration.go
	configurationPath := fmt.Sprintf("%s/configuration.go", configurationDir)
	configurationTemplate := template.ConfigurationTemplate()
	_ = fx.WriteFile(configurationPath, configurationTemplate.Text())

	// environment.go
	configurationEnvironmentPath := fmt.Sprintf("%s/environment.go", configurationDir)
	configurationEnvironmentTemplate := template.ConfigurationEnvironmentTemplate()
	_ = fx.WriteFile(configurationEnvironmentPath, configurationEnvironmentTemplate.Text())

	// development.yml
	configurationDevelopmentPath := fmt.Sprintf("%s/development.yml", configurationDir)
	configurationDevelopmentTemplate := template.ConfigurationDevelopmentTemplate()
	_ = fx.WriteFile(configurationDevelopmentPath, configurationDevelopmentTemplate.Text())

	// production.yml
	configurationProductionPath := fmt.Sprintf("%s/production.yml", configurationDir)
	configurationProductionTemplate := template.ConfigurationProductionTemplate()
	_ = fx.WriteFile(configurationProductionPath, configurationProductionTemplate.Text())

	// Create internal/pkg
	internalPkgDir := fmt.Sprintf("%s/internal/pkg", currentDir)
	_ = fx.EnsureDir(internalPkgDir)

	// Create internal/pkg/casbinx
	internalPkgCasbinxDir := fmt.Sprintf("%s/internal/pkg/casbinx", currentDir)
	_ = fx.EnsureDir(internalPkgCasbinxDir)

	// casbinx.go
	internalPkgCasbinxPath := fmt.Sprintf("%s/casbinx.go", internalPkgCasbinxDir)
	internalPkgCasbinxTemplate := template.InternalPkgCasbinxTemplate()
	_ = fx.WriteFile(internalPkgCasbinxPath, internalPkgCasbinxTemplate.Text())

	// Create internal/pkg/response
	internalPkgResponseDir := fmt.Sprintf("%s/internal/pkg/response", currentDir)
	_ = fx.EnsureDir(internalPkgResponseDir)

	// response.go
	internalPkgResponsePath := fmt.Sprintf("%s/response.go", internalPkgResponseDir)
	internalPkgResponseTemplate := template.InternalPkgResponseTemplate(opt.Module)
	_ = fx.WriteFile(internalPkgResponsePath, internalPkgResponseTemplate.Text())

	// Create internal/pkg/validator
	internalPkgValidatorDir := fmt.Sprintf("%s/internal/pkg/validator", currentDir)
	_ = fx.EnsureDir(internalPkgValidatorDir)

	// validator.go
	internalPkgValidatorPath := fmt.Sprintf("%s/validator.go", internalPkgValidatorDir)
	internalPkgValidatorTemplate := template.InternalPkgValidatorTemplate()
	err := fx.WriteFile(internalPkgValidatorPath, internalPkgValidatorTemplate.Text())

	// Update status
	if err == nil {
		spinnerGenProject.Success()
	} else {
		spinnerGenProject.Fail()
	}
}

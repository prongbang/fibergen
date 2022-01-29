package genx

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/template"
	"strings"

	"github.com/prongbang/fibergen/pkg/filex"
)

// Pkg is struct
type Pkg struct {
	Name   string
	Module Mod
}

// Mod is struct
type Mod struct {
	Module  string
	AppPath string
}

// Generator is the interface
type Generator interface {
	GenerateAll(opt option.Options)
	NewProject(opt option.Options)
	Generate(pkg Pkg, filename string)
	AutoBinding(pkg Pkg)
	DataSourceTemplate(pkg Pkg) string
	HandlerTemplate(pkg Pkg) string
	ProviderTemplate(pkg Pkg) string
	RepositoryTemplate(pkg Pkg) string
	RouterTemplate(pkg Pkg) string
	UseCaseTemplate(pkg Pkg) string
	ValidateTemplate(pkg Pkg) string
	ModelTemplate(pkg Pkg) string
	GetTemplate(pkg Pkg, filename string) string
	ModelName(feature string) string
	Templates(pkg Pkg) map[string]string
	GetModule() Mod
}

type generator struct {
	Fx filex.FileX
}

func (f *generator) NewProject(opt option.Options) {
	currentDir, _ := f.Fx.Getwd()

	// Create project directory
	currentDir = fmt.Sprintf("%s/%s", currentDir, opt.Project)
	_ = f.Fx.EnsureDir(currentDir)

	// Create go.mod
	modPath := fmt.Sprintf("%s/go.mod", currentDir)
	modTemplate := template.ModTemplate(opt.Module)
	_ = f.Fx.WriteFile(modPath, modTemplate.Text())

	// Create cmd
	cmdPath := fmt.Sprintf("%s/cmd/%s", currentDir, opt.Project)
	_ = f.Fx.EnsureDir(cmdPath)

	// main.go
	cmdMainPath := fmt.Sprintf("%s/main.go", cmdPath)
	cmdMainTemplate := template.CmdMainTemplate(opt.Module, opt.Project)
	_ = f.Fx.WriteFile(cmdMainPath, cmdMainTemplate.Text())

	// Create docs
	docsDir := fmt.Sprintf("%s/docs/apispec", currentDir)
	_ = f.Fx.EnsureDir(docsDir)

	// docs.go
	docsPath := fmt.Sprintf("%s/docs.go", docsDir)
	docsTemplate := template.DocsTemplate()
	_ = f.Fx.WriteFile(docsPath, docsTemplate.Text())

	// swagger.json
	docsSwaggerPath := fmt.Sprintf("%s/swagger.json", docsDir)
	docsSwaggerTemplate := template.DocsSwaggerJsonTemplate()
	_ = f.Fx.WriteFile(docsSwaggerPath, docsSwaggerTemplate.Text())

	// swagger.yaml
	docsSwaggerYamlPath := fmt.Sprintf("%s/swagger.yaml", docsDir)
	docsSwaggerYamlTemplate := template.DocsSwaggerYamlTemplate()
	_ = f.Fx.WriteFile(docsSwaggerYamlPath, docsSwaggerYamlTemplate.Text())

	// Create api
	apiDir := fmt.Sprintf("%s/internal/%s/api", currentDir, opt.Project)
	_ = f.Fx.EnsureDir(apiDir)

	// api.go
	apiPath := fmt.Sprintf("%s/api.go", apiDir)
	apiTemplate := template.ApiTemplate(opt.Module)
	_ = f.Fx.WriteFile(apiPath, apiTemplate.Text())

	// routers.go
	apiRoutersPath := fmt.Sprintf("%s/routers.go", apiDir)
	apiRoutersTemplate := template.ApiRoutersTemplate(opt.Module)
	_ = f.Fx.WriteFile(apiRoutersPath, apiRoutersTemplate.Text())

	// wire.go
	wireApiPath := fmt.Sprintf("%s/wire.go", apiDir)
	wireApiTemplate := template.WireApiTemplate(opt.Module, opt.Project)
	_ = f.Fx.WriteFile(wireApiPath, wireApiTemplate.Text())

	// wire_gen.go
	wireGenApiPath := fmt.Sprintf("%s/wire_gen.go", apiDir)
	wireGenApiTemplate := template.WireGenApiTemplate(opt.Module, opt.Project)
	_ = f.Fx.WriteFile(wireGenApiPath, wireGenApiTemplate.Text())

	// Create shared pkg/core
	databaseDir := fmt.Sprintf("%s/internal/%s/database", currentDir, opt.Project)
	_ = f.Fx.EnsureDir(databaseDir)

	// drivers.go
	databaseDriversPath := fmt.Sprintf("%s/drivers.go", databaseDir)
	databaseDriversTemplate := template.DatabaseDriversTemplate()
	_ = f.Fx.WriteFile(databaseDriversPath, databaseDriversTemplate.Text())

	// mongodb.go
	databaseMongodbPath := fmt.Sprintf("%s/mongodb.go", databaseDir)
	databaseMongodbTemplate := template.DatabaseMongodbTemplate()
	_ = f.Fx.WriteFile(databaseMongodbPath, databaseMongodbTemplate.Text())

	// wire.go
	databaseWirePath := fmt.Sprintf("%s/wire.go", databaseDir)
	databaseWireTemplate := template.DatabaseWireTemplate()
	_ = f.Fx.WriteFile(databaseWirePath, databaseWireTemplate.Text())

	// wire_gen.go
	databaseWireGenPath := fmt.Sprintf("%s/wire_gen.go", databaseDir)
	databaseWireGenTemplate := template.DatabaseWireGenTemplate()
	_ = f.Fx.WriteFile(databaseWireGenPath, databaseWireGenTemplate.Text())

	// Create deployments
	deploymentsDir := fmt.Sprintf("%s/deployments", currentDir)
	_ = f.Fx.EnsureDir(deploymentsDir)

	// Dockerfile
	deploymentsDockerfilePath := fmt.Sprintf("%s/Dockerfile", deploymentsDir)
	deploymentsDockerfileTemplate := template.DeploymentsDockerfileTemplate(opt.Module, opt.Project)
	_ = f.Fx.WriteFile(deploymentsDockerfilePath, deploymentsDockerfileTemplate.Text())

	// api-prod.yml
	deploymentsApiComposePath := fmt.Sprintf("%s/api-prod.yml", deploymentsDir)
	deploymentsApiComposeTemplate := template.DeploymentsApiComposeTemplate(opt.Project)
	_ = f.Fx.WriteFile(deploymentsApiComposePath, deploymentsApiComposeTemplate.Text())

	// Create shared pkg/core
	coreDir := fmt.Sprintf("%s/pkg/core", currentDir)
	_ = f.Fx.EnsureDir(coreDir)

	// handler.go
	coreHandlerPath := fmt.Sprintf("%s/handler.go", coreDir)
	coreHandlerTemplate := template.CoreHandlerTemplate()
	_ = f.Fx.WriteFile(coreHandlerPath, coreHandlerTemplate.Text())

	// paging.go
	corePagingPath := fmt.Sprintf("%s/paging.go", coreDir)
	corePagingTemplate := template.CorePagingTemplate()
	_ = f.Fx.WriteFile(corePagingPath, corePagingTemplate.Text())

	// params.go
	coreParamsPath := fmt.Sprintf("%s/params.go", coreDir)
	coreParamsTemplate := template.CoreParamsTemplate()
	_ = f.Fx.WriteFile(coreParamsPath, coreParamsTemplate.Text())

	// request.go
	coreRequestPath := fmt.Sprintf("%s/request.go", coreDir)
	coreRequestTemplate := template.CoreRequestTemplate()
	_ = f.Fx.WriteFile(coreRequestPath, coreRequestTemplate.Text())

	// response.go
	coreResponsePath := fmt.Sprintf("%s/response.go", coreDir)
	coreResponseTemplate := template.CoreResponseTemplate()
	_ = f.Fx.WriteFile(coreResponsePath, coreResponseTemplate.Text())

	// router.go
	coreRouterPath := fmt.Sprintf("%s/router.go", coreDir)
	coreRouterTemplate := template.CoreRouterTemplate()
	_ = f.Fx.WriteFile(coreRouterPath, coreRouterTemplate.Text())

	// validate.go
	coreValidatePath := fmt.Sprintf("%s/validate.go", coreDir)
	coreValidateTemplate := template.CoreValidateTemplate()
	_ = f.Fx.WriteFile(coreValidatePath, coreValidateTemplate.Text())

	// jwt.go
	coreJwtPath := fmt.Sprintf("%s/jwt.go", coreDir)
	coreJwtTemplate := template.CoreJwtTemplate()
	_ = f.Fx.WriteFile(coreJwtPath, coreJwtTemplate.Text())

	// flag.go
	coreFlagPath := fmt.Sprintf("%s/flag.go", coreDir)
	coreFlagTemplate := template.CoreFlagTemplate()
	_ = f.Fx.WriteFile(coreFlagPath, coreFlagTemplate.Text())

	// header.go
	coreHeaderPath := fmt.Sprintf("%s/header.go", coreDir)
	coreHeaderTemplate := template.CoreHeaderTemplate()
	_ = f.Fx.WriteFile(coreHeaderPath, coreHeaderTemplate.Text())

	// Create policy
	casbinPolicyDir := fmt.Sprintf("%s/policy", currentDir)
	_ = f.Fx.EnsureDir(casbinPolicyDir)

	// model.conf
	casbinModelPath := fmt.Sprintf("%s/model.conf", casbinPolicyDir)
	casbinModelTemplate := template.CasbinModelTemplate()
	_ = f.Fx.WriteFile(casbinModelPath, casbinModelTemplate.Text())

	// policy.csv
	casbinPolicyPath := fmt.Sprintf("%s/policy.conf", casbinPolicyDir)
	casbinPolicyTemplate := template.CasbinPolicyTemplate()
	_ = f.Fx.WriteFile(casbinPolicyPath, casbinPolicyTemplate.Text())

	// Create Makefile
	makefilePath := fmt.Sprintf("%s/Makefile", currentDir)
	makefileTemplate := template.MakefileTemplate(opt.Project)
	_ = f.Fx.WriteFile(makefilePath, makefileTemplate.Text())

	// Create configuration
	configurationDir := fmt.Sprintf("%s/configuration", currentDir)
	_ = f.Fx.EnsureDir(configurationDir)

	// configuration.go
	configurationPath := fmt.Sprintf("%s/configuration.go", configurationDir)
	configurationTemplate := template.ConfigurationTemplate()
	_ = f.Fx.WriteFile(configurationPath, configurationTemplate.Text())

	// environment.go
	configurationEnvironmentPath := fmt.Sprintf("%s/environment.go", configurationDir)
	configurationEnvironmentTemplate := template.ConfigurationEnvironmentTemplate()
	_ = f.Fx.WriteFile(configurationEnvironmentPath, configurationEnvironmentTemplate.Text())

	// development.yml
	configurationDevelopmentPath := fmt.Sprintf("%s/development.yml", configurationDir)
	configurationDevelopmentTemplate := template.ConfigurationDevelopmentTemplate()
	_ = f.Fx.WriteFile(configurationDevelopmentPath, configurationDevelopmentTemplate.Text())

	// production.yml
	configurationProductionPath := fmt.Sprintf("%s/production.yml", configurationDir)
	configurationProductionTemplate := template.ConfigurationProductionTemplate()
	_ = f.Fx.WriteFile(configurationProductionPath, configurationProductionTemplate.Text())

	// Create internal/pkg
	internalPkgDir := fmt.Sprintf("%s/internal/pkg", currentDir)
	_ = f.Fx.EnsureDir(internalPkgDir)

	// Create internal/pkg/casbinx
	internalPkgCasbinxDir := fmt.Sprintf("%s/internal/pkg/casbinx", currentDir)
	_ = f.Fx.EnsureDir(internalPkgCasbinxDir)

	// casbinx.go
	internalPkgCasbinxPath := fmt.Sprintf("%s/casbinx.go", internalPkgCasbinxDir)
	internalPkgCasbinxTemplate := template.InternalPkgCasbinxTemplate()
	_ = f.Fx.WriteFile(internalPkgCasbinxPath, internalPkgCasbinxTemplate.Text())

	fmt.Println(fmt.Sprintf("Create project \"%s\" success", opt.Project))
}

func (f *generator) Templates(pkg Pkg) map[string]string {
	return map[string]string{
		"datasource.go":                f.DataSourceTemplate(pkg),
		"handler.go":                   f.HandlerTemplate(pkg),
		"provider.go":                  f.ProviderTemplate(pkg),
		"repository.go":                f.RepositoryTemplate(pkg),
		"router.go":                    f.RouterTemplate(pkg),
		"usecase.go":                   f.UseCaseTemplate(pkg),
		"validate.go":                  f.ValidateTemplate(pkg),
		fmt.Sprintf("%s.go", pkg.Name): f.ModelTemplate(pkg),
	}
}

func (f *generator) GenerateAll(opt option.Options) {
	fmt.Println("--> START")
	if opt.Project != "" && opt.Module != "" {
		f.NewProject(opt)
	} else if opt.Feature != "" {
		mod := f.GetModule()
		pkg := Pkg{
			Name:   opt.Feature,
			Module: mod,
		}
		for filename := range f.Templates(pkg) {
			f.Generate(pkg, filename)
		}
		f.AutoBinding(pkg)
	} else {
		fmt.Println("Not Supported")
	}
	fmt.Println("<-- END")
}

func (f *generator) AutoBinding(pkg Pkg) {
	pwd, _ := f.Fx.Getwd()
	routerPath := pwd + "/routers.go"
	wirePath := pwd + "/wire.go"

	routerB := f.Fx.ReadFile(routerPath)
	wireB := f.Fx.ReadFile(wirePath)
	routerText := routerB
	wireText := wireB

	// Binding wire
	wireImpPat := "//+fibergen:import wire:package"
	wireImp := fmt.Sprintf(
		`"%s/%s/api/%s"
	%s`, pkg.Module.Module, pkg.Module.AppPath, pkg.Name, wireImpPat,
	)
	wireText = strings.Replace(wireText, wireImpPat, wireImp, 1)

	wireBuildPat := "//+fibergen:func wire:build"
	wireBuild := fmt.Sprintf(
		`%s.ProviderSet,
		%s`, pkg.Name, wireBuildPat,
	)
	wireText = strings.Replace(wireText, wireBuildPat, wireBuild, 1)

	if err := f.Fx.WriteFile(wirePath, []byte(wireText)); err == nil {
		fmt.Println(fmt.Sprintf("Binding file %s success", "wire.go"))
	} else {
		fmt.Println(fmt.Sprintf("Binding file %s failure", "wire.go"))
	}

	// Binding routers
	routerImpPat := "//+fibergen:import routers:package"
	routerImp := fmt.Sprintf(
		`"%s/%s/api/%s"
	%s`, pkg.Module.Module, pkg.Module.AppPath, pkg.Name, routerImpPat,
	)
	routerText = strings.Replace(routerText, routerImpPat, routerImp, 1)

	routerStructPat := "//+fibergen:struct routers"
	routerStruct := fmt.Sprintf(
		`%sRoute %s.Router
	%s`, f.ModelName(pkg.Name), pkg.Name, routerStructPat,
	)
	routerText = strings.Replace(routerText, routerStructPat, routerStruct, 1)

	routerInitPat := "//+fibergen:func initials"
	routerInit := fmt.Sprintf(
		`r.%sRoute.Initial(app)
	%s`, f.ModelName(pkg.Name), routerInitPat,
	)
	routerText = strings.Replace(routerText, routerInitPat, routerInit, 1)

	routerNewPat := "//+fibergen:func new:routers"
	routerNew := fmt.Sprintf(
		`%sRoute %s.Router,
	%s`, pkg.Name, pkg.Name, routerNewPat,
	)
	routerText = strings.Replace(routerText, routerNewPat, routerNew, 1)

	routerBindPat := "//+fibergen:return &routers"
	routerBind := fmt.Sprintf(
		`%sRoute: %sRoute,
		%s`, f.ModelName(pkg.Name), pkg.Name, routerBindPat,
	)
	routerText = strings.Replace(routerText, routerBindPat, routerBind, 1)

	if err := f.Fx.WriteFile(routerPath, []byte(routerText)); err == nil {
		fmt.Println(fmt.Sprintf("Binding file %s success", "routers.go"))
	} else {
		fmt.Println(fmt.Sprintf("Binding file %s failure", "routers.go"))
	}
}

func (f *generator) Generate(pkg Pkg, filename string) {
	template := f.GetTemplate(pkg, filename)
	currentDir, err := f.Fx.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	currentDir = fmt.Sprintf("%s/%s", currentDir, pkg.Name)
	if f.Fx.EnsureDir(currentDir) != nil {
		fmt.Println("Create directory error")
		return
	}
	target := fmt.Sprintf("%s/%s", currentDir, filename)
	if err := f.Fx.WriteFile(target, []byte(template)); err != nil {
		fmt.Println("Generate file error", err)
	} else {
		fmt.Println(fmt.Sprintf("Generate file %s success", filename))
	}
}

func (f *generator) DataSourceTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s

import "%s/%s/database"

type DataSource interface {
}

type dataSource struct {
	Driver database.Drivers
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}`, pkg.Name, pkg.Module.Module, pkg.Module.AppPath)
}

func (f *generator) HandlerTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s

type Handler interface {
}

type handler struct {
	Uc UseCase
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}`, pkg.Name)
}

func (f *generator) ProviderTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDataSource,
	NewRepository,
	NewUseCase,
	NewHandler,
	NewRouter,
)`, pkg.Name)
}

func (f *generator) RepositoryTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s

type Repository interface {
}

type repository struct {
	Ds DataSource
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}`, pkg.Name)
}

func (f *generator) RouterTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
	"%s/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	Handle Handler
}

func (r *router) Initial(app *fiber.App) {
}

func NewRouter(handle Handler) Router {
	return &router{
		Handle: handle,
	}
}`, pkg.Name, pkg.Module.Module)
}

func (f *generator) UseCaseTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s
	
type UseCase interface {
}

type useCase struct {
	Repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}`, pkg.Name)
}

func (f *generator) ValidateTemplate(pkg Pkg) string {
	return fmt.Sprintf(`package %s

type Validate interface {
}

type validate struct {
}

func NewValidate() Validate {
	return &validate{}
}
`, pkg.Name)
}

func (f *generator) ModelTemplate(pkg Pkg) string {
	model := f.ModelName(pkg.Name)
	return fmt.Sprintf(`package %s
	
type %s struct  {
}`, pkg.Name, model)
}

func (f *generator) GetTemplate(pkg Pkg, filename string) string {
	return f.Templates(pkg)[filename]
}

func (f *generator) ModelName(feature string) string {
	first := strings.ToUpper(feature[:1])
	last := feature[1:]
	modelName := fmt.Sprintf("%s%s", first, last)
	return modelName
}

func (f *generator) GetModule() Mod {
	// Get current path
	pwd, _ := f.Fx.Getwd()

	// Get root project path
	changeToRootProject := "../../../"
	_ = f.Fx.Chdir(changeToRootProject)
	root, _ := f.Fx.Getwd()
	if bt := f.Fx.ReadFile(root + "/go.mod"); bt != "" {
		// Find module
		text := bt
		m := "module "
		s := strings.Index(text, m)
		e := strings.Index(text, "\n")
		if s < 0 && e < 0 {
			return Mod{}
		}
		mod := text[s+len(m) : e]

		// Find app path
		i := strings.LastIndex(mod, "/")
		if i < 0 {
			return Mod{}
		}
		pj := mod[i:]
		ign := "/api"
		c := strings.Index(pwd, pj)

		// Find internal/project-name
		ap := pwd[c+len(pj)+1 : len(pwd)-len(ign)]

		_ = f.Fx.Chdir("./" + ap + ign)

		return Mod{
			Module:  mod,
			AppPath: ap,
		}
	}
	return Mod{}
}

// NewGenerator is new instance with func
func NewGenerator(fx filex.FileX) Generator {
	return &generator{
		Fx: fx,
	}
}

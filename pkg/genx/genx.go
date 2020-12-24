package genx

import (
	"fmt"
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
	GenerateAll(feature string)
	Generate(pkg Pkg, filename string)
	AutoBinding(pkg Pkg)
	DataSourceTemplate(pkg Pkg) string
	HandlerTemplate(pkg Pkg) string
	ProviderTemplate(pkg Pkg) string
	RepositoryTemplate(pkg Pkg) string
	RouterTemplate(pkg Pkg) string
	UseCaseTemplate(pkg Pkg) string
	ModelTemplate(pkg Pkg) string
	GetTemplate(pkg Pkg, filename string) string
	ModelName(feature string) string
	Templates(pkg Pkg) map[string]string
	GetModule() Mod
}

type generator struct {
	Fx filex.FileX
}

func (f *generator) Templates(pkg Pkg) map[string]string {
	return map[string]string{
		"datasource.go":                f.DataSourceTemplate(pkg),
		"handler.go":                   f.HandlerTemplate(pkg),
		"provider.go":                  f.ProviderTemplate(pkg),
		"repository.go":                f.RepositoryTemplate(pkg),
		"router.go":                    f.RouterTemplate(pkg),
		"usecase.go":                   f.UseCaseTemplate(pkg),
		fmt.Sprintf("%s.go", pkg.Name): f.ModelTemplate(pkg),
	}
}

func (f *generator) GenerateAll(feature string) {
	mod := f.GetModule()
	pkg := Pkg{
		Name:   feature,
		Module: mod,
	}
	fmt.Println("--> START")
	for filename := range f.Templates(pkg) {
		f.Generate(pkg, filename)
	}
	f.AutoBinding(pkg)
	fmt.Println("<-- END")
}

func (f *generator) AutoBinding(pkg Pkg) {
	pwd, _ := f.Fx.Getwd()
	routerPath := pwd + "/routers.go"
	wirePath := pwd + "/wire.go"

	routerB := f.Fx.ReadFile(routerPath)
	wireB := f.Fx.ReadFile(wirePath)
	routerText := string(routerB)
	wireText := string(wireB)

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

import "github.com/gofiber/fiber/v2"

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
}`, pkg.Name)
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
	f.Fx.Chdir(changeToRootProject)
	root, _ := f.Fx.Getwd()
	if bt := f.Fx.ReadFile(root + "/go.mod"); bt != "" {
		// Find module
		text := string(bt)
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
		c := strings.LastIndex(pwd, pj)
		ap := pwd[c+len(pj)+1 : len(pwd)-len(ign)]

		f.Fx.Chdir("./" + ap + ign)

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

package genx

import (
	"fmt"
	"log"
	"strings"

	"github.com/prongbang/fibergen/pkg/filex"
)

// Generator is the interface
type Generator interface {
	GenerateAll(feature string)
	Generate(feature string, filename string)
	DataSourceTemplate(pkg string) string
	HandlerTemplate(pkg string) string
	ProviderTemplate(pkg string) string
	RepositoryTemplate(pkg string) string
	RouterTemplate(pkg string) string
	UseCaseTemplate(pkg string) string
	ModelTemplate(pkg string) string
	GetTemplate(pkg string, filename string) string
	ModelName(feature string) string
	Templates(pkg string) map[string]string
}

type generator struct {
	Fx filex.FileX
}

func (f *generator) Templates(pkg string) map[string]string {
	return map[string]string{
		"datasource.go":           f.DataSourceTemplate(pkg),
		"handler.go":              f.HandlerTemplate(pkg),
		"provider.go":             f.ProviderTemplate(pkg),
		"repository.go":           f.RepositoryTemplate(pkg),
		"router.go":               f.RouterTemplate(pkg),
		"usecase.go":              f.UseCaseTemplate(pkg),
		fmt.Sprintf("%s.go", pkg): f.ModelTemplate(pkg),
	}
}

func (f *generator) GenerateAll(feature string) {
	log.Println("--> START")
	for filename := range f.Templates(feature) {
		f.Generate(feature, filename)
	}
	log.Println("<-- END")
}

func (f *generator) Generate(feature string, filename string) {
	template := f.GetTemplate(feature, filename)
	currentDir, err := f.Fx.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	currentDir = fmt.Sprintf("%s/%s", currentDir, feature)
	if f.Fx.EnsureDir(currentDir) != nil {
		log.Println("Create directory error")
		return
	}
	target := fmt.Sprintf("%s/%s", currentDir, filename)
	if err := f.Fx.WriteFile(target, []byte(template)); err != nil {
		log.Println("Generate file error", err)
	} else {
		log.Println(fmt.Sprintf("Generate file %s success", filename))
	}
}

func (f *generator) DataSourceTemplate(pkg string) string {
	return fmt.Sprintf(`package %s
type DataSource interface {
}
type dataSource struct {
	DbSource database.DataSource
}
func NewDataSource(dbSource database.DataSource) DataSource {
	return &dataSource{
		DbSource: dbSource,
	}
}`, pkg)
}

func (f *generator) HandlerTemplate(pkg string) string {
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
}`, pkg)
}

func (f *generator) ProviderTemplate(pkg string) string {
	return fmt.Sprintf(`package %s
import "github.com/google/wire"
var ProviderSet = wire.NewSet(
	NewDataSource,
	NewRepository,
	NewUseCase,
	NewHandler,
	NewRouter,
)`, pkg)
}

func (f *generator) RepositoryTemplate(pkg string) string {
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
}`, pkg)
}

func (f *generator) RouterTemplate(pkg string) string {
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
	return &router{Handle: handle}
}`, pkg)
}

func (f *generator) UseCaseTemplate(pkg string) string {
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
}`, pkg)
}

func (f *generator) ModelTemplate(pkg string) string {
	model := f.ModelName(pkg)
	return fmt.Sprintf(`package %s
type %s struct  {
}`, pkg, model)
}

func (f *generator) GetTemplate(pkg string, filename string) string {
	return f.Templates(pkg)[filename]
}

func (f *generator) ModelName(feature string) string {
	first := strings.ToUpper(feature[:1])
	last := feature[1:]
	modelName := fmt.Sprintf("%s%s", first, last)
	return modelName
}

// NewGenerator is new instance with func
func NewGenerator(fx filex.FileX) Generator {
	return &generator{
		Fx: fx,
	}
}

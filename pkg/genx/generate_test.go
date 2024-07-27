package genx_test

import (
	"fmt"
	"testing"

	"github.com/prongbang/fibergen/pkg/option"
	"github.com/prongbang/fibergen/pkg/pkgs"
	"github.com/prongbang/fibergen/pkg/template"
	"github.com/prongbang/fibergen/pkg/tocase"
	"github.com/prongbang/fibergen/pkg/tools"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
	"github.com/prongbang/fibergen/pkg/mod"
)

var module string = "github.com/prongbang/fibergen"
var appPath string = "internal/app"
var pwd string = "/usr/go/src/github.com/prongbang/fibergen/internal/app/api"
var pwdPath string = "/usr/go/src/github.com/prongbang/fibergen/internal/app/api"
var rootPath string = "/usr/go/src/github.com/prongbang/fibergen"
var read string = ""

type fileXMock struct {
}

func (f *fileXMock) IsExist(filename string) bool {
	return true
}

func (f *fileXMock) EnsureDir(dir string) error {
	return nil
}

func (f *fileXMock) WriteFile(filename string, data []byte) error {
	return nil
}

func (f *fileXMock) Chdir(dir string) error {
	if dir == "../../../" {
		pwd = rootPath
		read = "module " + module + "\n"
	} else {
		pwd = pwdPath
		read = ""
	}
	return nil
}

func (f *fileXMock) ReadFile(filename string) string {
	return read
}

func (f *fileXMock) Getwd() (string, error) {
	return pwd, nil
}

func NewFileXMock() filex.FileX {
	return &fileXMock{}
}

type fileXMockError struct {
}

func (f *fileXMockError) IsExist(filename string) bool {
	return true
}

type installerMock struct {
}

// Install implements tools.Installer.
func (*installerMock) Install() error {
	return nil
}

func NewInstallerMock() tools.Installer {
	return &installerMock{}
}

type runnerMock struct{}

// Run implements tools.Runner.
func (*runnerMock) Run() error {
	return nil
}

func NewRunnerMock() tools.Runner {
	return &runnerMock{}
}

var writeFile error = nil
var changeDir error = nil
var getwdPath string = ""
var readFile string = ""

func (f *fileXMockError) EnsureDir(dir string) error {
	return fmt.Errorf("%s", "Error")
}

func (f *fileXMockError) WriteFile(filename string, data []byte) error {
	return writeFile
}

func (f *fileXMockError) Chdir(dir string) error {
	return changeDir
}

func (f *fileXMockError) ReadFile(filename string) string {
	return readFile
}

func (f *fileXMockError) Getwd() (string, error) {
	return getwdPath, fmt.Errorf("%s", "Error")
}

func NewFileXMockError() filex.FileX {
	return &fileXMockError{}
}

var genX genx.Generator
var genXError genx.Generator

func init() {
	tl := NewInstallerMock()
	runner := NewRunnerMock()
	fx := NewFileXMock()
	fxError := NewFileXMockError()
	genX = genx.NewGenerator(fx, tl, tl, runner)
	genXError = genx.NewGenerator(fxError, tl, tl, runner)
}

func TestGenerateWriteFileError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	tmpl := ""
	writeFile = fmt.Errorf("%s", "Error")
	pkg := pkgs.Pkg{
		Name: feature,
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genx.GenerateFeature(NewFileXMock(), pkg, filename, tmpl)
}

func TestGenerateEnsureDirError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	tmpl := ""
	pkg := pkgs.Pkg{
		Name: feature,
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genx.GenerateFeature(NewFileXMock(), pkg, filename, tmpl)
}

func TestGenerateGetwdError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	tmpl := ""
	pkg := pkgs.Pkg{
		Name: feature,
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genx.GenerateFeature(NewFileXMock(), pkg, filename, tmpl)
}

func TestTemplates(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "user",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	m := template.FeatureTemplates(pkg)
	if m["datasource.go"] == "" {
		t.Error("Error")
	}
	if m["handler.go"] == "" {
		t.Error("Error")
	}
	if m["provider.go"] == "" {
		t.Error("Error")
	}
	if m["repository.go"] == "" {
		t.Error("Error")
	}
	if m["router.go"] == "" {
		t.Error("Error")
	}
	if m["usecase.go"] == "" {
		t.Error("Error")
	}
	if m[fmt.Sprintf("%s.go", pkg.Name)] == "" {
		t.Error("Error")
	}
}

func TestGenerateAll(t *testing.T) {
	opt := option.Options{
		Project: "",
		Module:  "",
		Feature: "hello",
	}
	genX.Generate(opt)
}

func TestGenerate(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	tmpl := ""
	pkg := pkgs.Pkg{
		Name: feature,
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genx.GenerateFeature(NewFileXMock(), pkg, filename, tmpl)
}

func TestDataSourceTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.DataSource(pkg.Name, pkg.Module.Module, pkg.Module.AppPath) == "" {
		t.Error("Error")
	}
}

func TestHandlerTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.Handler(pkg.Name) == "" {
		t.Error("Error")
	}
}

func TestProviderTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.Provider(pkg.Name) == "" {
		t.Error("Error")
	}
}

func TestRepositoryTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.Repository(pkg.Name) == "" {
		t.Error("Error")
	}
}

func TestRouterTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.Router(pkg.Name, pkg.Module.Module) == "" {
		t.Error("Error")
	}
}

func TestUseCaseTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.UseCase(pkg.Name) == "" {
		t.Error("Error")
	}
}

func TestValidateTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.Validate(pkg.Name) == "" {
		t.Error("Error")
	}
}

func TestModelTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if template.Model(pkg.Name) == "" {
		t.Error("Error")
	}
}

func TestGetTemplate(t *testing.T) {
	pkg := pkgs.Pkg{
		Name: "hello",
		Module: mod.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	filename := "usecase.go"
	if template.FeatureTemplates(pkg)[filename] == "" {
		t.Error("Error")
	}
}

func TestUpperCamelName(t *testing.T) {
	feature := "hello_model"
	if tocase.UpperCamelName(feature) != "HelloModel" {
		t.Error("Error")
	}
}

func TestLowerCamelName(t *testing.T) {
	feature := "hello_model_test"
	actual := tocase.LowerCamelName(feature)
	if actual != "helloModelTest" {
		t.Error("Error", actual)
	}
}

func TestGetModuleName(t *testing.T) {
	m := mod.GetModule(NewFileXMock())
	if m.Module == "" {
		t.Error("Error")
	}
}

package genx_test

import (
	"fmt"
	"testing"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
)

var module string = "github.com/prongbang/fibergen"
var appPath string = "internal/app"
var pwd string = "/usr/go/src/github.com/prongbang/fibergen/internal/app/api"
var pwdPath string = "/usr/go/src/github.com/prongbang/fibergen/internal/app/api"
var rootPath string = "/usr/go/src/github.com/prongbang/fibergen"
var read string = ""

type fileXMock struct {
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

var ensureDir error = fmt.Errorf("%s", "Error")
var writeFile error = nil
var changeDir error = nil
var getwdPath string = ""
var readFile string = ""
var getedError error = fmt.Errorf("%s", "Error")

func (f *fileXMockError) EnsureDir(dir string) error {
	return ensureDir
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
	return getwdPath, getedError
}

func NewFileXMockError() filex.FileX {
	return &fileXMockError{}
}

var genX genx.Generator
var genXError genx.Generator

func init() {
	fx := NewFileXMock()
	fxError := NewFileXMockError()
	genX = genx.NewGenerator(fx)
	genXError = genx.NewGenerator(fxError)
}

func TestGenerateWriteFileError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	getedError = nil
	ensureDir = nil
	writeFile = fmt.Errorf("%s", "Error")
	pkg := genx.Pkg{
		Name: feature,
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genXError.Generate(pkg, filename)
}

func TestGenerateEnsureDirError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	getedError = nil
	ensureDir = fmt.Errorf("%s", "Error")
	pkg := genx.Pkg{
		Name: feature,
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genXError.Generate(pkg, filename)
}

func TestGenerateGetwdError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	getedError = fmt.Errorf("%s", "Error")
	pkg := genx.Pkg{
		Name: feature,
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genXError.Generate(pkg, filename)
}

func TestTemplates(t *testing.T) {
	pkg := genx.Pkg{
		Name: "user",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	m := genX.Templates(pkg)
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
	feature := "hello"
	genX.GenerateAll(feature)
}

func TestGenerate(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	pkg := genx.Pkg{
		Name: feature,
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	genX.Generate(pkg, filename)
}

func TestDataSourceTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.DataSourceTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestHandlerTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.HandlerTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestProviderTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.ProviderTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestRepositoryTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.RepositoryTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestRouterTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.RouterTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestUseCaseTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.UseCaseTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestModelTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	if genX.ModelTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestGetTemplate(t *testing.T) {
	pkg := genx.Pkg{
		Name: "hello",
		Module: genx.Mod{
			Module:  module,
			AppPath: appPath,
		},
	}
	filename := "usecase.go"
	if genX.GetTemplate(pkg, filename) == "" {
		t.Error("Error")
	}
}

func TestModelName(t *testing.T) {
	feature := "hello"
	if genX.ModelName(feature) != "Hello" {
		t.Error("Error")
	}
}

func TestGetModuleName(t *testing.T) {
	mod := genX.GetModule()
	if mod.Module == "" {
		t.Error("Error")
	}
}

package genx_test

import (
	"fmt"
	"testing"

	"github.com/prongbang/fibergen/pkg/filex"
	"github.com/prongbang/fibergen/pkg/genx"
)

type fileXMock struct {
}

func (f *fileXMock) EnsureDir(dir string) error {
	return nil
}

func (f *fileXMock) WriteFile(filename string, data []byte) error {
	return nil
}

func (f *fileXMock) Getwd() (string, error) {
	return "/mock/path", nil
}

func NewFileXMock() filex.FileX {
	return &fileXMock{}
}

type fileXMockError struct {
}

var ensureDir error = fmt.Errorf("%s", "Error")
var writeFile error = nil
var getwdPath string = ""
var getedError error = fmt.Errorf("%s", "Error")

func (f *fileXMockError) EnsureDir(dir string) error {
	return ensureDir
}

func (f *fileXMockError) WriteFile(filename string, data []byte) error {
	return writeFile
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
	genXError.Generate(feature, filename)
}

func TestGenerateEnsureDirError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	getedError = nil
	ensureDir = fmt.Errorf("%s", "Error")
	genXError.Generate(feature, filename)
}

func TestGenerateGetwdError(t *testing.T) {
	feature := "hello"
	filename := "usecase.go"
	getedError = fmt.Errorf("%s", "Error")
	genXError.Generate(feature, filename)
}

func TestTemplates(t *testing.T) {
	pkg := "user"
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
	if m[fmt.Sprintf("%s.go", pkg)] == "" {
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
	genX.Generate(feature, filename)
}

func TestDataSourceTemplate(t *testing.T) {
	pkg := "hello"
	if genX.DataSourceTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestHandlerTemplate(t *testing.T) {
	pkg := "hello"
	if genX.HandlerTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestProviderTemplate(t *testing.T) {
	pkg := "hello"
	if genX.ProviderTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestRepositoryTemplate(t *testing.T) {
	pkg := "hello"
	if genX.RepositoryTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestRouterTemplate(t *testing.T) {
	pkg := "hello"
	if genX.RouterTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestUseCaseTemplate(t *testing.T) {
	pkg := "hello"
	if genX.UseCaseTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestModelTemplate(t *testing.T) {
	pkg := "hello"
	if genX.ModelTemplate(pkg) == "" {
		t.Error("Error")
	}
}

func TestGetTemplate(t *testing.T) {
	pkg := "hello"
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

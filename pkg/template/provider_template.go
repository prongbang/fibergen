package template

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/tocase"
)

func Provider(name string) string {
	return fmt.Sprintf(`package %s

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDataSource,
	NewRepository,
	NewUseCase,
	NewHandler,
	NewRouter,
	NewValidate,
)`, tocase.ToLower(name))
}

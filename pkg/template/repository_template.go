package template

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/tocase"
)

func Repository(name string) string {
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
}`, tocase.ToLower(name))
}

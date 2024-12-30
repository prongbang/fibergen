package template

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/tocase"
)

func UseCase(name string) string {
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
}`, tocase.ToLower(name))
}

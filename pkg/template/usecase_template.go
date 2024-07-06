package template

import "fmt"

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
}`, name)
}

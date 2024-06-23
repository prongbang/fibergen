package template

import "fmt"

func Handler(name string) string {
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
}`, name)
}

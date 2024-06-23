package template

import (
	"strings"

	"github.com/ettle/strcase"
)

func UseCaseCrud(
	model string,
	pk string,
) string {
	tmpl := `package {name}
	
	type UseCase interface {
		Count(params Params) int64
		FindList(params Params) []{model}
		FindById(id {pk}) {model}
		Create(data *Create{model}) ({model}, error)
		Update(data *Update{model}) ({model}, error)
		Delete(id {pk}) error
	}
	
	type useCase struct {
		Repo Repository
	}

	func (u *useCase) Count(params Params) int64 {
		return u.Repo.Count(params)
	}
	
	func (u *useCase) FindList(params Params) []{model} {
		return u.Repo.FindList(params)
	}
	
	func (u *useCase) FindById(id {pk}) {model} {
		return u.Repo.FindById(id)
	}
	
	func (u *useCase) Create(data *Create{model}) ({model}, error) {
		err := u.Repo.Create(data)
		if err != nil {
			return {model}{}, err
		}
		return u.Repo.FindById(data.Id), nil
	}
	
	func (u *useCase) Update(data *Update{model}) ({model}, error) {
		err := u.Repo.Update(data)
		if err != nil {
			return {model}{}, err
		}
		return u.Repo.FindById(data.Id), nil
	}
	
	func (u *useCase) Delete(id {pk}) error {
		return u.Repo.Delete(id)
	}	

	func NewUseCase(repo Repository) UseCase {
		return &useCase{
			Repo: repo,
		}
	}`

	tmpl = strings.ReplaceAll(tmpl, "{pk}", pk)
	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(model))
	tmpl = strings.ReplaceAll(tmpl, "{name}", strings.ToLower(model))

	return tmpl
}

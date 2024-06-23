package template

import (
	"strings"

	"github.com/ettle/strcase"
)

func RepositoryCrud(
	model string,
	pk string,
) string {
	tmpl := `package {name}

	type Repository interface {
		Count(params Params) int64
		FindList(params Params) []{model}
		FindById(id {pk}) {model}
		Create(data *Create{model}) error
		Update(data *Update{model}) error
		Delete(id {pk}) error
	}
	
	type repository struct {
		Ds DataSource
	}
	
	func (r *repository) Count(params Params) int64 {
		return r.Ds.Count(params)
	}
	
	func (r *repository) FindList(params Params) []{model} {
		return r.Ds.FindList(params)
	}

	func (r *repository) FindById(id {pk}) {model} {
		return r.Ds.FindById(id)
	}

	func (r *repository) Create(data *Create{model}) error {
		return r.Ds.Create(data)
	}

	func (r *repository) Update(data *Update{model}) error {
		return r.Ds.Update(data)
	}

	func (r *repository) Delete(id {pk}) error {
		return r.Ds.Delete(id)
	}

	func NewRepository(ds DataSource) Repository {
		return &repository{
			Ds: ds,
		}
	}`

	tmpl = strings.ReplaceAll(tmpl, "{pk}", pk)
	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(model))
	tmpl = strings.ReplaceAll(tmpl, "{name}", strings.ToLower(model))

	return tmpl
}

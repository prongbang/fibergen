package option

import "github.com/prongbang/fibergen/pkg/template"

type Options struct {
	Project string
	Module  string
	Feature string
	Crud    string
	Spec    string
	Driver  string
}

type Spec struct {
	Driver          string
	QueryColumns    string
	Fields          []template.Field
	PrimaryField    template.PrimaryField
	Columns         []string
	InsertValues    string
	InsertFields    string
	InsertQuestions string
	UpdateSets      string
}

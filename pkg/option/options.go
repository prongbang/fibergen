package option

import "github.com/prongbang/fibergen/pkg/template"

type Options struct {
	Project string
	Module  string
	Feature string
	Shared  string
	Spec    string
	Driver  string
}

type Spec struct {
	Imports         []string
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

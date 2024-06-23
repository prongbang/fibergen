package option

type Options struct {
	Project string
	Module  string
	Feature string
	Crud    string
	Spec    string
	Driver  string
}

type Spec struct {
	Pk              string
	Driver          string
	QueryColumns    string
	Fields          string
	InsertValues    string
	InsertFields    string
	InsertQuestions string
	UpdateSets      string
}

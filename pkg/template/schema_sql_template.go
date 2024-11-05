package template

type schemaSqlTemplate struct {
}

func (c *schemaSqlTemplate) Text() []byte {
	return []byte(`package schema

var OrderBy = map[string]bool{
	"asc":  true,
	"desc": true,
}
`)
}

func SchemaSqlTemplate() Template {
	return &schemaSqlTemplate{}
}

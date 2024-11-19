package template

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
)

func ModelCrud(imports []string, module string, pk string, name string, fields string, columns []string) string {
	model := strcase.ToPascal(name)
	tmpl := `package {name}
	
import ({import}
	"{module}/pkg/core"
)

var columns = map[string]bool{
{columns}
}

type {model} struct {
{fields}
}

type {model}Lite struct {
{fields}
}

type Create{model} struct {
{fields}
}

type Update{model} struct {
{fields}
}

type Delete{model} struct {
	Id {pk} {jsonId}
}

type QueryOne struct {
	Id {pk} {jsonId}
}

type QueryMany struct {
	core.Params
}

type LiteQueryMany struct {
	core.Sorting
}

type Params struct {
	QueryOne
	QueryMany
}

type LiteParams struct {
	LiteQueryMany
}
`
	if len(imports) == 1 {
		tmpl = strings.ReplaceAll(tmpl, "{import}", fmt.Sprintf("\n\t%s", imports[0]))
	} else if len(imports) > 1 {
		tmpl = strings.ReplaceAll(tmpl, "{import}", strings.Join(imports, "\n\t"))
	} else {
		tmpl = strings.ReplaceAll(tmpl, "{import}", "")
	}

	validate := ""
	if strings.Contains(pk, "int") {
		validate += " validate:\"gt=0\""
	}

	tmpl = strings.ReplaceAll(tmpl, "{module}", module)
	tmpl = strings.ReplaceAll(tmpl, "{pk}", pk)
	tmpl = strings.ReplaceAll(tmpl, "{jsonId}", fmt.Sprintf("`json:\"id\"%s`", validate))
	tmpl = strings.ReplaceAll(tmpl, "{page}", "`json:\"page\"`")
	tmpl = strings.ReplaceAll(tmpl, "{limit}", "`json:\"limit\"`")
	tmpl = strings.ReplaceAll(tmpl, "{model}", model)
	tmpl = strings.ReplaceAll(tmpl, "{fields}", fields)
	tmpl = strings.ReplaceAll(tmpl, "{name}", name)

	var colBuilder strings.Builder
	for _, item := range columns {
		colBuilder.WriteString(fmt.Sprintf("\t\"%s\": true,\n", item))
	}
	tmpl = strings.ReplaceAll(tmpl, "{columns}", colBuilder.String())

	return tmpl
}

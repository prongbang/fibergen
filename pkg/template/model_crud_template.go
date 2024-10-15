package template

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
)

func ModelCrud(imports []string, module string, pk string, name string, fields string) string {
	model := strcase.ToPascal(name)
	tmpl := `package {name}
	
import ({import}
	"{module}/pkg/core"
)

type {model} struct {
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

type Params struct {
	QueryOne
	QueryMany
}
`
	if len(imports) == 1 {
		tmpl = strings.ReplaceAll(tmpl, "{import}", fmt.Sprintf("\n\t%s", imports[0]))
	} else if len(imports) > 1 {
		tmpl = strings.ReplaceAll(tmpl, "{import}", strings.Join(imports, "\n\t"))
	} else {
		tmpl = strings.ReplaceAll(tmpl, "{import}", "")
	}

	tmpl = strings.ReplaceAll(tmpl, "{module}", module)
	tmpl = strings.ReplaceAll(tmpl, "{pk}", pk)
	tmpl = strings.ReplaceAll(tmpl, "{jsonId}", "`json:\"id\"`")
	tmpl = strings.ReplaceAll(tmpl, "{page}", "`json:\"page\"`")
	tmpl = strings.ReplaceAll(tmpl, "{limit}", "`json:\"limit\"`")
	tmpl = strings.ReplaceAll(tmpl, "{model}", model)
	tmpl = strings.ReplaceAll(tmpl, "{fields}", fields)
	tmpl = strings.ReplaceAll(tmpl, "{name}", name)

	return tmpl
}

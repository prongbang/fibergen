package template

import (
	"fmt"

	"github.com/ettle/strcase"
)

func Model(name string) string {
	model := strcase.ToPascal(name)
	return fmt.Sprintf(`package %s
	
type %s struct {
}
`, name, model)
}

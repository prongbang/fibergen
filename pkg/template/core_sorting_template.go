package template

import (
	"strings"
)

type coreSortingTemplate struct {
}

func (c *coreSortingTemplate) Text() []byte {
	var builder strings.Builder
	builder.WriteString("package core\n")
	builder.WriteString("\n")
	builder.WriteString("type Sorting struct {\n")
	builder.WriteString("\tSort  string `json:\"sort\"`\n")
	builder.WriteString("\tOrder  string `json:\"order\"`\n")
	builder.WriteString("}\n")
	return []byte(builder.String())
}

func CoreSortingTemplate() Template {
	return &coreSortingTemplate{}
}

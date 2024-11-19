package template

import "strings"

type coreParamsTemplate struct {
}

func (c *coreParamsTemplate) Text() []byte {
	lines := []string{
		"package core",
		"",
		"type Params struct {",
		"	Offset 	int64  `json:\"offset\"`",
		"	Page   	int64  `json:\"page\" validate:\"gt=0\"`",
		"	Limit  	int64  `json:\"limit\" validate:\"gt=0,lte=100\"`",
		"	Sorting",
		"}",
	}
	return []byte(strings.Join(lines, "\n"))
}

func CoreParamsTemplate() Template {
	return &coreParamsTemplate{}
}

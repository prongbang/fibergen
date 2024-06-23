package template

import "strings"

type coreParamsTemplate struct {
}

func (c *coreParamsTemplate) Text() []byte {
	lines := []string{
		"package core",
		"",
		"type Params struct {",
		"	Offset int64 `json:\"offset\"`",
		"	Page   int64 `json:\"page\"`",
		"	Limit  int64 `json:\"limit\"`",
		"}",
	}
	return []byte(strings.Join(lines, "\n"))
}

func CoreParamsTemplate() Template {
	return &coreParamsTemplate{}
}

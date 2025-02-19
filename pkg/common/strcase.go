package common

import (
	"fmt"
	"github.com/ettle/strcase"
	"strings"
)

func UpperCamelName(name string) string {
	names := strings.Split(name, "_")
	modelName := ""
	for _, v := range names {
		first := strings.ToUpper(v[:1])
		last := v[1:]
		modelName += fmt.Sprintf("%s%s", first, last)
	}
	return modelName
}

func LowerCamelName(name string) string {
	names := strings.Split(name, "_")
	modelName := ""
	first := ""
	last := ""
	for i, v := range names {
		if i == 0 {
			first = strings.ToLower(v[:1])
		} else {
			first = strings.ToUpper(v[:1])
		}
		last = v[1:]
		modelName += fmt.Sprintf("%s%s", first, last)
	}
	return modelName
}

func ToLower(name string) string {
	return strings.ReplaceAll(strcase.ToKebab(name), "-", "")
}

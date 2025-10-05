package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ettle/strcase"
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

func Abbrev(s string) string {
	// Step 1: Insert space before uppercase letters (to handle camelCase)
	s = regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllString(s, `${1} ${2}`)

	// Step 2: Normalize all separators (space, underscore, hyphen)
	s = strings.ToLower(s)
	parts := regexp.MustCompile(`[\s\-_]+`).Split(s, -1)

	// Step 3: Collect first letters
	var b strings.Builder
	for _, p := range parts {
		if len(p) > 0 {
			b.WriteByte(p[0])
		}
	}
	return b.String()
}

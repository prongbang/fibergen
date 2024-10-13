package typer

import (
	"fmt"
	"strings"
)

func Get(value any) string {
	if value == nil {
		return "any"
	}

	switch v := value.(type) {
	case string:
		return "string"
	case int, int8, int16, int32, int64:
		return "int64"
	case bool:
		return "bool"
	case float32, float64:
		if strings.Index(fmt.Sprintf("%v", v), ".") > -1 {
			return "float64"
		}
		return "int64"
	default:
		return "any"
	}
}

func Value(typ string) string {
	if typ == "any" {
		return "nil"
	} else if typ == "string" {
		return `""`
	} else if typ == "int64" {
		return "0"
	} else if typ == "float64" {
		return "0.0"
	}
	return ""
}

func Operate(typ string) string {
	if typ == "any" || typ == "string" {
		return `!=`
	} else if typ == "int64" || typ == "float64" {
		return ">"
	}
	return "!="
}

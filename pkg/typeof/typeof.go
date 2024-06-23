package typeof

import (
	"fmt"
	"reflect"
)

func Type(variable interface{}) string {
	switch variable.(type) {
	case int:
	case int16:
	case int32:
	case int64:
		return "int64"
	case float32:
	case float64:
		return "float64"
	case bool:
		return "boolean"
	case string:
		return "string"
	}
	return "interface{}"
}

func SprintOf(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func TypeOf(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}

func ValueOf(variable interface{}) string {
	return reflect.ValueOf(variable).Kind().String()
}

package template

func StructXStruct() string {
	tmpl := `package structx

import (
	"reflect"
)

func Name(obj interface{}) string {
	name := reflect.TypeOf(obj).Name()
	return name
}
`
	return tmpl
}

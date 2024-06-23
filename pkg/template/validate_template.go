package template

import "fmt"

func Validate(name string) string {
	return fmt.Sprintf(`package %s

	type Validate interface {
	}
	
	type validate struct {
	}
	
	func NewValidate() Validate {
		return &validate{}
	}
	`, name)
}

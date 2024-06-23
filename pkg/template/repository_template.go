package template

import "fmt"

func Repository(name string) string {
	return fmt.Sprintf(`package %s

	type Repository interface {
	}
	
	type repository struct {
		Ds DataSource
	}
	
	func NewRepository(ds DataSource) Repository {
		return &repository{
			Ds: ds,
		}
	}`, name)
}

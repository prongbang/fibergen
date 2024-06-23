package template

import "fmt"

func Provider(name string) string {
	return fmt.Sprintf(`package %s

	import "github.com/google/wire"
	
	var ProviderSet = wire.NewSet(
		NewDataSource,
		NewRepository,
		NewUseCase,
		NewHandler,
		NewRouter,
		NewValidate,
	)`, name)
}

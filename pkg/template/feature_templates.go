package template

import (
	"github.com/prongbang/fibergen/pkg/option"
)

func FeatureTemplates(pkg option.Package) map[string]string {
	return map[string]string{
		//"datasource.go":                DataSource(pkg.Name, pkg.Module.Module, pkg.Module.AppPath),
		//"handler.go":                   Handler(pkg.Name),
		//"provider.go":                  Provider(pkg.Name),
		//"repository.go":                Repository(pkg.Name),
		//"router.go":                    Router(pkg.Name, pkg.Module.Module),
		//"usecase.go":                   UseCase(pkg.Name),
		//fmt.Sprintf("%s.go", pkg.Name): Model(pkg.Name),
	}
}

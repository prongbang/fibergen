package template

import (
	"fmt"

	"github.com/prongbang/fibergen/pkg/pkgs"
)

func FeatureCrudTemplates(pkg pkgs.Pkg) map[string]string {
	return map[string]string{
		"datasource.go": DataSourceCrud(
			pkg.Name,
			pkg.Module.Module,
			pkg.Module.AppPath,
			pkg.Spec.Pk,
			pkg.Spec.Driver,
			pkg.Spec.QueryColumns,
			pkg.Spec.InsertValues,
			pkg.Spec.InsertFields,
			pkg.Spec.InsertQuestions,
			pkg.Spec.UpdateSets,
		),
		"handler.go": HandlerCrud(
			pkg.Name,
			pkg.Module.Module,
			pkg.Spec.Pk,
		),
		"provider.go": Provider(pkg.Name),
		"repository.go": RepositoryCrud(
			pkg.Name,
			pkg.Spec.Pk,
		),
		"router.go": RouterCrud(
			pkg.Name,
			pkg.Module.Module,
		),
		"usecase.go": UseCaseCrud(
			pkg.Name,
			pkg.Spec.Pk,
		),
		"validate.go": ValidateCrud(
			pkg.Name,
		),
		fmt.Sprintf("%s.go", pkg.Name): ModelCrud(
			pkg.Imports,
			pkg.Module.Module,
			pkg.Spec.Pk,
			pkg.Name,
			pkg.Spec.Fields,
		),
	}
}

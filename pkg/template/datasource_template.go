package template

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/tocase"
)

func DataSource(pkgName string, moduleName string, modulePath string) string {
	return fmt.Sprintf(`package %s

import "%s/%s/database"

type DataSource interface {
}

type dataSource struct {
	Driver database.Drivers
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}`, tocase.ToLower(pkgName), moduleName, modulePath)
}

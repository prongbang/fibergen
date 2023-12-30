package forgot

import "github.com/prongbang/mvp/test_test/internal/test_test/database"

type DataSource interface {
}

type dataSource struct {
	Driver database.Drivers
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}
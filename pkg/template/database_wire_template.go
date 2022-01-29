package template

type databaseWireTemplate struct {
}

func (w *databaseWireTemplate) Text() []byte {
	return []byte(`//+build wireinject

package database

import (
	"github.com/google/wire"
)

func NewDatabaseDriver() Drivers {
	wire.Build(NewMongoDbDriver, NewDrivers)
	return nil
}
`)
}

func DatabaseWireTemplate() Template {
	return &databaseWireTemplate{}
}

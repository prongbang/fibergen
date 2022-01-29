package template

type databaseDriversTemplate struct {
}

func (d *databaseDriversTemplate) Text() []byte {
	return []byte(`package database

import (
	"github.com/innotechdevops/mgo-driver/pkg/mgodriver"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Drivers interface {
	GetMongoDB() *mongo.Database
}

type drivers struct {
	MongoDB     *mongo.Database
	MariaDB     *sqlx.DB
	MongoDriver mgodriver.MongoDriver
}

func (d *drivers) GetMongoDB() *mongo.Database {
	return d.MongoDB
}

func NewDrivers(
	mongoDB mgodriver.MongoDriver,
) Drivers {
	return &drivers{
		MongoDB:     mongoDB.Connect(),
		MongoDriver: mongoDB,
	}
}`)
}

func DatabaseDriversTemplate() Template {
	return &databaseDriversTemplate{}
}

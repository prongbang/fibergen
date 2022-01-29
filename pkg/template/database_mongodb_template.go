package template

type databaseMongodbTemplate struct {
}

func (m *databaseMongodbTemplate) Text() []byte {
	return []byte(`package database

import (
	"github.com/innotechdevops/mgo-driver/pkg/mgodriver"
	"github.com/spf13/viper"
)

/*
NewMongoDbDriver a new instance
MongoDB:
(>) greater than - $gt

(<) less than - $lt

(>=) greater than equal to - $gte

(<= ) less than equal to - $lte
*/
func NewMongoDbDriver() mgodriver.MongoDriver {
	return mgodriver.New(mgodriver.Config{
		User:         viper.GetString("mongodb.user"),
		Pass:         viper.GetString("mongodb.pass"),
		Host:         viper.GetString("mongodb.host"),
		DatabaseName: viper.GetString("mongodb.database"),
		Port:         viper.GetString("mongodb.port"),
	})
}
`)
}

func DatabaseMongodbTemplate() Template {
	return &databaseMongodbTemplate{}
}

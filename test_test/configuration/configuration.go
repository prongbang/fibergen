package configuration

import (
	"github.com/spf13/viper"
	"log"
)

func Load(env string) {
	viper.SetConfigName(env)
	viper.AddConfigPath("configuration")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

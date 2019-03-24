package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string, port string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config")
	config.Set("env", env)
	config.Set("port", port)
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing ", env, " configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
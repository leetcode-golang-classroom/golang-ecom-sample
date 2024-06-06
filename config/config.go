package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT string `mapstructure:"PORT"`

	MYSQL_DATABASE string `mapstructure:"MYSQL_DATABASE"`
	MYSQL_USER     string `mapstructure:"MYSQL_USER"`
	MYSQL_PASSWORD string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_ADDR     string `mapstructure:"MYSQL_ADDR"`
}

var C *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		failOnError(err, "Failed to read config")
	}
	v.AutomaticEnv()

	err = v.Unmarshal(&C)
	if err != nil {
		failOnError(err, "Failed to read enivroment")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

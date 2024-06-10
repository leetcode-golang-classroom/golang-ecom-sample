package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`

	MysqlDatabase string `mapstructure:"MYSQL_DATABASE"`
	MysqlUser     string `mapstructure:"MYSQL_USER"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlAddr     string `mapstructure:"MYSQL_ADDR"`

	JWTSecret              string `mapstructure:"JWT_SECRET"`
	JWTExpirationInSeconds int64  `mapstructure:"JWT_EXPIRATION_IN_SECONDS"`
}

var C *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	failOnError(v.BindEnv("PORT"), "fail on Bind PORT")
	failOnError(v.BindEnv("MYSQL_DATABASE"), "fail on Bind MYSQL_DATABASE")
	failOnError(v.BindEnv("MYSQL_USER"), "fail on Bind MYSQL_USER")
	failOnError(v.BindEnv("MYSQL_PASSWORD"), "fail on Bind MYSQL_PASSWORD")
	failOnError(v.BindEnv("MYSQL_ADDR"), "fail on Bind MYSQL_ADDR")
	failOnError(v.BindEnv("JWT_SECRET"), "fail on Bind JWT_SECRET")
	failOnError(v.BindEnv("JWT_EXPIRATION_IN_SECONDS"), "fail on Bind JWT_EXPIRATION_IN_SECONDS")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("load from environment variable")
	}
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

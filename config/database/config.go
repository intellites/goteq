package database

import (
	"log"

	"github.com/spf13/viper"
)

var env Config

type Config struct {
	DB_Host     string
	DB_Port     int
	DB_SSL_Mode string
	DB_Database string
	DB_Username string
	DB_Password string
}

func init() {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
	var err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatalln("Error reading config file:", err)
	}

	log.Println("[INIT] Configuration loaded")
}

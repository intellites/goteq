package routers

import (
	"log"

	"github.com/spf13/viper"
)

var env Config

type Config struct {
	JWT_SECRET string
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

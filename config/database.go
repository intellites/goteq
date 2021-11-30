package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	if DB != nil {
		return
	}

	// Initialize DB
	log.Println("[INIT] Connecting to PostgreSQL")
	DB, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			env.DB_Host, env.DB_Port, env.DB_Database,
			env.DB_Username, env.DB_Password, env.DB_SSL_Mode),
	), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
	}
	log.Println("[INIT] connected to PostgreSQL")
}

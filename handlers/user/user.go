package user

import (
	"github.com/intellites/goteq/config/database"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func init() {
	database.DB.AutoMigrate(&User{})
}

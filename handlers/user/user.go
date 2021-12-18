package user

import (
	"github.com/intellites/goteq/config/database"
)

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

func init() {
	database.DB.Statement.AutoMigrate(&User{})
}

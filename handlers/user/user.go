package user

import (
	"github.com/intellites/goteq/config/database"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;unique" validate:"required"`
	Description string `json:"description" gorm:"not null" validate:"required"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique" validate:"required"`
	Email    string `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles"`
}

func init() {
	// Create role tabale and seed the data
	if err := database.DB.Statement.AutoMigrate(&Role{}); err == nil && database.DB.Migrator().HasTable(&Role{}) {
		// Seed the roles into the table
		if database.DB.First(&Role{}).Error == gorm.ErrRecordNotFound {
			database.DB.Create(&Role{
				Name:        "admin",
				Description: "Administrator",
			})
			database.DB.Create(&Role{
				Name:        "user",
				Description: "User",
			})
		}
	}

	database.DB.Statement.AutoMigrate(&User{})
}

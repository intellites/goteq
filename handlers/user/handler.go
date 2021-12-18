package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/intellites/goteq/config/database"
	"github.com/intellites/goteq/helpers/provider"
	"github.com/intellites/goteq/util"
)

// User for create the users in the system. This is use by the
// middleware to validate the user.
func CreateAuthUser(w http.ResponseWriter, r *http.Request) {
	// Get the Email and password from the request
	user := User{}

	// Decode the request body into the struct
	json.NewDecoder(r.Body).Decode(&user)

	// Validate the request data
	if util.Validate(w, user) {
		return
	}

	// Get role for the user
	role := Role{}
	if err := database.DB.Where("name = ?", "admin").First(&role).Error; err == nil {
		user.Roles = []Role{role}
	}

	// Encrypt the password
	user.Password = util.HashAndSalt([]byte(user.Password))

	// Create the user
	if err := database.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			provider.Error(w, http.StatusConflict, "User already exists!")
			return
		}
	}

	provider.Success(w, http.StatusCreated, user)
}

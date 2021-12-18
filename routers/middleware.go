package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/intellites/goteq/config/database"
	"github.com/intellites/goteq/handlers/user"
	"github.com/intellites/goteq/helpers/provider"
	"github.com/intellites/goteq/util"
)

var MiddlewareRoutes = RoutePrefix{
	"/api/auth",
	[]Route{
		{"Middlware Autentication", "POST", "/token", Authenticate, false},
		{"Middlware Autentication", "POST", "/users", user.CreateAuthUser, false},
	},
}

// Get token from the header
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			provider.Error(w, http.StatusUnauthorized, "Missing Authorization Header")
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := VerifyToken(tokenString)

		if err != nil {
			provider.Error(w, http.StatusUnauthorized, "Error verifying JWT token: "+err.Error())
			return
		}

		email := claims.(jwt.MapClaims)["email"].(string)

		// Get the user from the database
		user := GetUserByEmail(email)

		// Validate the user
		if len(user.Email) == 0 {
			provider.Error(w, http.StatusUnauthorized, "User not found.")
			return
		}

		// Marshal the user roles to json
		roles, _ := json.Marshal(GetUserRoles(user))

		// Set the user to the request context
		r.Header.Set("email", email)
		r.Header.Set("roles", string(roles))

		next.ServeHTTP(w, r)
	})
}

// AuthHandler is the handler for the auth route
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// Get the Email and password from the request
	var request user.User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		provider.Error(w, http.StatusUnauthorized, "Error decoding request body.")
		return
	}

	if len(request.Email) == 0 || len(request.Password) == 0 {
		provider.Error(w, http.StatusUnauthorized, "Please provide email and password to obtain the token.")
		return
	}

	// Get the user from the database
	user := GetUserByEmail(request.Email)

	if len(user.Email) == 0 {
		provider.Error(w, http.StatusUnauthorized, "User not found.")
		return
	}

	if user.Email == request.Email && util.ComparePassword(user.Password, []byte(request.Password)) {
		token, err := GetToken(user)
		if err != nil {
			log.Fatalln(err.Error())
			provider.Error(w, http.StatusInternalServerError, "Error generating JWT token")
		} else {
			provider.Success(w, http.StatusOK, map[string]string{"token": token})
		}
	} else {
		provider.Error(w, http.StatusUnauthorized, "Email or password do not match!")
	}
}

// verifyToken checks the token and returns the claims
func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(env.JWT_SECRET)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, err
}

// Get Token generates a token
func GetToken(user user.User) (string, error) {
	signingKey := []byte(env.JWT_SECRET)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"roles": GetUserRoles(user), // Get user roles
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

// Get user roles for token
func GetUserRoles(user user.User) []string {
	tokenRoles := make([]string, 0)
	roles := user.Roles

	if len(roles) > 0 {
		for _, role := range roles {
			tokenRoles = append(tokenRoles, strings.ToLower(role.Name))
		}
	}

	return tokenRoles
}

// Get user email for authentication
func GetUserByEmail(email string) user.User {
	user := user.User{}

	// Find user by email
	database.DB.Preload("Roles").Where("email = ?", email).First(&user)

	return user
}

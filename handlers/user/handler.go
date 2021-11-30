package user

import (
	"fmt"
	"net/http"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	var user User
	fmt.Fprintln(w, user)
}

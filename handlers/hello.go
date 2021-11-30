package handlers

import (
	"fmt"
	"net/http"

	"github.com/intellites/goteq/models"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	var data models.Hello
	data.Name = "Nadeen Gamage"
	fmt.Fprintln(w, data)
}

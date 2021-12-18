package util

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/intellites/goteq/helpers/provider"
)

// Validate the request data with model
func Validate(w http.ResponseWriter, model interface{}) bool {
	// Validate the model with rules
	if err := validator.New().Struct(model); err != nil {
		var response []string

		for _, e := range err.(validator.ValidationErrors) {
			response = append(response, e.Error())
		}

		// If there is an error, return an error
		provider.Success(w, http.StatusBadRequest, response)
		return true
	}

	return false
}

package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

var AppRoutes []RoutePrefix

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

func NewRouter() *mux.Router {
	// Init router
	router := mux.NewRouter()

	// Append handlers
	AppRoutes = append(AppRoutes, Routes)

	for _, route := range AppRoutes {

		// Create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		// Loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			// Check to see if route should be protected with jwt
			if r.Protected {
				// handler = JSONContentTypeMiddleware(r.HandlerFunc)
			}

			// Attach sub route
			routePrefix.Path(r.Pattern).Handler(handler).Methods(r.Method).Name(r.Name)
		}
	}
	return router
}

package routers

import (
	"github.com/intellites/goteq/handlers"
)

var Routes = RoutePrefix{
	"/api/v1",
	[]Route{
		{"UsersIndex", "GET", "/hello", handlers.Hello, false},
	},
}

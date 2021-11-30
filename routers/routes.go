package routers

import (
	"github.com/intellites/goteq/handlers/user"
)

var Routes = RoutePrefix{
	"/api/v1",
	[]Route{
		{"UsersIndex", "GET", "/users", user.UsersIndex, false},
	},
}

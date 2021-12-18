package routers

var Routes = RoutePrefix{
	"/api/v1",
	[]Route{
		{"Authentication", "POST", "/auth/token", Authenticate, false},
	},
}

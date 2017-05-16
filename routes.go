package usermgmt

import (
	"net/http"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

func (env *Env) userRoutes() Routes {
	var userRoutes = Routes{
		Route{
			"RegisterHandler",
			"POST",
			"/register",
			env.RegisterHandler,
		},
		Route{
			"LoginHandler",
			"POST",
			"/login",
			env.LoginHandler,
		},
		Route{
			"LoginHandler",
			"POST",
			"/logout",
			env.LogoutHandler,
		},
	}
	return userRoutes
}

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
			env.AddUserHandler,
		},
		Route{
			"LoginHandler",
			"POST",
			"/login",
			env.LoginHandler,
		},
		Route{
			"LogoutHandler",
			"GET",
			"/logout",
			env.LogoutHandler,
		},
	}
	return userRoutes
}

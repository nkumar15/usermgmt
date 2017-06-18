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
			"/user",
			env.AddUserHandler,
		},
		Route{
			"GetUserHandler",
			"GET",
			"/user/{Id}",
			env.GetUserHandler,
		},
		Route{
			"GetUsersHandler",
			"GET",
			"/user",
			env.GetUsersHandler,
		},
		Route{
			"DeleteUserHandler",
			"DELETE",
			"/user/{Id}",
			env.DeleteUserHandler,
		},
		Route{
			"UpdateUserHandler",
			"PUT",
			"/user/{Id}",
			env.UpdateUserHandler,
		},
	}
	return userRoutes
}

func (env *Env) authRoutes() Routes {
	var authRoutes = Routes{
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

	return authRoutes
}

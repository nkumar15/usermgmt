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

func (conf *Configuration) userRoutes() Routes {

	var userRoutes = Routes{
		Route{
			"AddUserHandler",
			"POST",
			"/user",
			conf.AddUserHandler,
		},
		Route{
			"GetUserHandler",
			"GET",
			"/user/{Id}",
			conf.GetUserHandler,
		},
		Route{
			"GetUsersHandler",
			"GET",
			"/user",
			conf.GetUsersHandler,
		},
		Route{
			"DeleteUserHandler",
			"DELETE",
			"/user/{Id}",
			conf.DeleteUserHandler,
		},
		Route{
			"UpdateUserHandler",
			"PUT",
			"/user/{Id}",
			conf.UpdateUserHandler,
		},
	}
	return userRoutes
}

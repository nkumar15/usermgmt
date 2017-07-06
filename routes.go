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

const (
	AddUserRoute    = "/user"
	GetUserRoute    = "/user/{Id}"
	GetUsersRoute   = "/user"
	DeleteUserRoute = "/user/{Id}"
	UpdateUserRoute = "/user/{Id}"
)

// var userRoutes = Routes{
// 	Route{
// 		"AddUserHandler",
// 		"POST",
// 		"/user",
// 		AddUserHandler,
// 	},
// 	Route{
// 		"GetUserHandler",
// 		"GET",
// 		"/user/{Id}",
// 		GetUserHandler,
// 	},
// 	Route{
// 		"GetUsersHandler",
// 		"GET",
// 		"/user",
// 		GetUsersHandler,
// 	},
// 	Route{
// 		"DeleteUserHandler",
// 		"DELETE",
// 		"/user/{Id}",
// 		DeleteUserHandler,
// 	},
// 	Route{
// 		"UpdateUserHandler",
// 		"PUT",
// 		"/user/{Id}",
// 		UpdateUserHandler,
// 	},
// }

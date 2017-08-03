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

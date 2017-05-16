package usermgmt

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter ...
func (env *Env) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range env.userRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		//handler = handler
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

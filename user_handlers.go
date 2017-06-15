package usermgmt

import (
	"net/http"

	"github.com/gorilla/schema"
)

// AddUserHandler ...
func (env *Env) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	user := new(User)
	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.Form)
	if err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	if err := env.Database.addUser(user); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	renderJSON(w, user)
}

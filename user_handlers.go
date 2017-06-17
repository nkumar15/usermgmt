package usermgmt

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if err := env.userDb.AddUser(user); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	renderJSON(w, user)
}

// GetUserHandler ...
func (env *Env) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["Id"], 10, 64)
	if err != nil {
		httpStatusBadRequest(w, err)
		return
	}

	user, err := env.userDb.GetUserById(id)
	if err != nil {
		if err.Error() == "ErrNoMoreRows" {
			httpStatusNotFound(w, r, err)
		} else {
			httpStatusInternalServerError(w, err)
		}
		return
	}

	renderJSON(w, user)
}

// GetUsersHandler ...
func (env *Env) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := env.userDb.GetUsers()
	if err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	renderJSON(w, users)
}

// DeleteUserHandler ...
func (env *Env) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["Id"], 10, 64)
	if err != nil {
		httpStatusBadRequest(w, err)
		return
	}

	if err = env.userDb.DeleteUserById(id); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	httpStatusNoContent(w, r)
}

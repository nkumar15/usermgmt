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

	if err := user.addUser(env.DB); err != nil {
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

	user := new(User)
	user.ID = id
	if err = user.getUser(env.DB); err != nil {
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

	user := new(User)
	users, err := user.getUsers(env.DB)

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

	user := new(User)
	user.ID = id
	if err = user.deleteUser(env.DB); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}
	httpStatusNoContent(w, r)
}

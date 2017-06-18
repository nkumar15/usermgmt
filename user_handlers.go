package usermgmt

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func getIDFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["Id"], 10, 64)
	return id, err
}

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

	id, err := getIDFromRequest(r)
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

	id, err := getIDFromRequest(r)
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

// UpdateUserHandler ...
func (env *Env) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getIDFromRequest(r)
	if err != nil {
		httpStatusBadRequest(w, err)
		return
	}

	err = r.ParseForm()
	user := new(User)

	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.Form)
	if err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	user.ID = id
	if err := env.userDb.UpdateUser(user); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}

	renderJSON(w, user)
}

package usermgmt

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func getUserIDFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["Id"], 10, 64)
	return id, err
}

// AddUserHandler ...
func (conf *Configuration) AddUserHandler(w http.ResponseWriter, r *http.Request) {

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

	if err := conf.userDb.AddUser(user); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}
	renderJSON(w, user)
}

// GetUserHandler ...
func (conf *Configuration) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		httpStatusBadRequest(w, err)
		return
	}

	user, err := conf.userDb.GetUserById(id)
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
func (conf *Configuration) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := conf.userDb.GetUsers()
	if err != nil {
		httpStatusInternalServerError(w, err)
		return
	}
	renderJSON(w, users)
}

// DeleteUserHandler ...
func (conf *Configuration) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		httpStatusBadRequest(w, err)
		return
	}

	if err = conf.userDb.DeleteUserById(id); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}
	httpStatusNoContent(w, r)
}

// UpdateUserHandler ...
func (conf *Configuration) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getUserIDFromRequest(r)
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
	if err := conf.userDb.UpdateUser(user); err != nil {
		httpStatusInternalServerError(w, err)
		return
	}
	renderJSON(w, user)
}

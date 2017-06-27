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
func AddUserHandler(conf *Configuration, w http.ResponseWriter, r *http.Request) (int, error) {

	err := r.ParseForm()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user := new(User)
	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.Form)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err := conf.userDb.AddUser(user); err != nil {
		return http.StatusInternalServerError, err
	}
	renderJSON(w, user)
	return 200, nil
}

// GetUserHandler ...
func GetUserHandler(conf *Configuration, w http.ResponseWriter, r *http.Request) (int, error) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user, err := conf.userDb.GetUserByID(id)
	if err != nil {
		if err.Error() == "ErrNoMoreRows" {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}
	renderJSON(w, user)
	return 200, nil
}

// GetUsersHandler ...
func GetUsersHandler(conf *Configuration, w http.ResponseWriter, r *http.Request) (int, error) {

	users, err := conf.userDb.GetUsers()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	renderJSON(w, users)
	return 200, nil
}

// DeleteUserHandler ...
func DeleteUserHandler(conf *Configuration, w http.ResponseWriter, r *http.Request) (int, error) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if err = conf.userDb.DeleteUserByID(id); err != nil {
		return http.StatusInternalServerError, err
	}
	httpStatusNoContent(w, r)
	return 200, nil
}

// UpdateUserHandler ...
func UpdateUserHandler(conf *Configuration, w http.ResponseWriter, r *http.Request) (int, error) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = r.ParseForm()
	user := new(User)

	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.Form)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.ID = id
	if err := conf.userDb.UpdateUser(user); err != nil {
		return http.StatusInternalServerError, err
	}
	renderJSON(w, user)
	return 200, nil
}

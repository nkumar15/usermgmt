package usermgmt

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// userHandler ...
type userHandler struct {
	Conf *Configuration
}

func newUserHandler(conf *Configuration) *userHandler {
	return &userHandler{Conf: conf}
}

// RegisterHandlers ...
func RegisterHandlers(router *mux.Router, conf *Configuration) {

	uh := newUserHandler(conf)

	router.HandleFunc(AddUserRoute, uh.addUser).Methods("POST")
	router.HandleFunc(GetUserRoute, uh.getUser).Methods("GET")
	router.HandleFunc(GetUsersRoute, uh.getUsers).Methods("GET")
	router.HandleFunc(DeleteUserRoute, uh.deleteUser).Methods("DELETE")
	router.HandleFunc(UpdateUserRoute, uh.updateUser).Methods("PUT")
}

func getUserIDFromRequest(r *http.Request) (int64, error) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["Id"], 10, 64)
	return id, err
}

func (uh *userHandler) addUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	conf := uh.Conf

	if err != nil {
		httpStatusInternalServerError(w, err)
	}

	user := new(User)
	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.Form)
	if err != nil {
		httpStatusInternalServerError(w, err)
	}

	if err := conf.db.AddUser(user); err != nil {
		httpStatusInternalServerError(w, err)
	}

	renderJSON(w, user)
}

// GetUserHandler ...
func (uh *userHandler) getUser(w http.ResponseWriter, r *http.Request) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		httpStatusInternalServerError(w, err)
	}

	conf := uh.Conf
	user, err := conf.db.GetUserByID(id)
	if err != nil {
		if err.Error() == "ErrNoMoreRows" {
			httpStatusNotFound(w, r, err)
		}
		httpStatusInternalServerError(w, err)
	}

	if conf.serveAsAPI == true {
		renderJSON(w, user)
	} else {
		renderUser(w, user)
	}
}

// GetUsersHandler ...
func (uh *userHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	conf := uh.Conf
	users, err := conf.db.GetUsers()
	if err != nil {
		httpStatusInternalServerError(w, err)
	}

	if conf.serveAsAPI == true {
		renderJSON(w, users)
	} else {
		renderUsers(w, users)
	}

}

// DeleteUserHandler ...
func (uh *userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		httpStatusBadRequest(w, err)
	}

	conf := uh.Conf
	if err = conf.db.DeleteUserByID(id); err != nil {
		httpStatusInternalServerError(w, err)
	}
	httpStatusNoContent(w, r)
}

// UpdateUserHandler ...
func (uh *userHandler) updateUser(w http.ResponseWriter, r *http.Request) {

	id, err := getUserIDFromRequest(r)
	if err != nil {
		httpStatusBadRequest(w, err)
	}

	err = r.ParseForm()
	user := new(User)

	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.Form)
	if err != nil {
		httpStatusInternalServerError(w, err)
	}

	user.ID = id
	conf := uh.Conf
	if err := conf.db.UpdateUser(user); err != nil {
		httpStatusInternalServerError(w, err)
	}
	renderJSON(w, user)
}

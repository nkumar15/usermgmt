package usermgmt

import (
	"encoding/json"
	"log"
	"net/http"
)

func fillUser(user *User) {
	user.Name = "Neeraj"
	user.Password = "pwd"
	user.Email = "email"
}

func httpGenericError(w http.ResponseWriter) {

	http.Error(w, "Something went wrong, check the logs", http.StatusInternalServerError)
}

func httpStatusInternalServerError(w http.ResponseWriter, err error) {

	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func httpStatusBadRequest(w http.ResponseWriter, err error) {

	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func httpStatusNotFound(w http.ResponseWriter, r *http.Request, err error) {

	log.Println(err.Error())
	http.NotFound(w, r)
}

func httpStatusNoContent(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNoContent)
}

func renderJSON(w http.ResponseWriter, data interface{}) {

	var j []byte
	var err error

	j, err = json.Marshal(data)
	if err != nil {
		log.Println(err)
		httpGenericError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

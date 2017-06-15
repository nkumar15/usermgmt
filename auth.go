package usermgmt

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Session struct {
	ID              string
	Authenticated   bool
	Unauthenticated bool
	User            User
}

// LoginHandler ...
func (env *Env) LoginHandler(w http.ResponseWriter, r *http.Request) {

	validateSession(w, r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(422) // unprocessable entity
		return
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(user.Name, user.Password)
	if Authenticate(user) {
		//http.Redirect(w, r, "http://localhost:5000/success", http.StatusTemporaryRedirect)
		io.WriteString(w, "Success")
	} else {
		//http.Redirect(w, r, "http://localhost:5000/login", http.StatusUnauthorized)
		io.WriteString(w, "fail")
	}
}

// LogoutHandler ...
func (env *Env) LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

// Authenticate ...
func Authenticate(user User) bool {
	return user.Name == "admin" && user.Password == "pwd"
}

func getSessionUID(sid string) int {
	//user := User{}
	//some logic here
	return 1
}

func updateSession(sid string, uid int) {

}

func generateSessionID() string {
	//sid := make([]byte, 24)
	return "a"
}

func validateSession(w http.ResponseWriter, r *http.Request) {

}

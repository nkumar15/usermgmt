package usermgmt

import (
	"encoding/json"
	"log"
	"net/http"
)

// User ...
type User struct {
	Id       int
	Name     string
	Password string
}

// RegisterHandler ...
func (env *Env) RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

// LoginHandler ...
func (env *Env) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	log.Println("LoginHandler called", user.Name, user.Password)
	if Authenticate(user) {
		http.Redirect(w, r, "http://localhost:5000/success", http.StatusPermanentRedirect)
	} else {
		http.Redirect(w, r, "http://localhost:5000/login", http.StatusUnauthorized)
	}
}

// LogoutHandler ...
func (env *Env) LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

// Authenticate ...
func Authenticate(user User) bool {
	return user.Name == "admin" && user.Password == "pwd"
}

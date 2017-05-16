package usermgmt

import (
	"encoding/json"
	"net/http"
)

// User ...
type User struct {
	Id       int
	Name     string
	Password string
}

// RegisterHandler ...
func (env *Env) RegisterHandler(w http.ResponseWriter, r *http.Request) bool {
	return true
}

// LoginHandler ...
func (env *Env) LoginHandler(w http.ResponseWriter, r *http.Request) bool {
	var user User

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return false
	}

	return Authenticate(user)
}

// LogoutHandler ...
func (env *Env) LogoutHandler(w http.ResponseWriter, r *http.Request) bool {
	return true
}

// Authenticate ...
func Authenticate(user User) bool {
	return user.Name == "admin" && user.Password == "pwd"
}

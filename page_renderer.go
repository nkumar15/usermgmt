package usermgmt

import (
	"html/template"
	"net/http"
)

func renderUsers(w http.ResponseWriter, users *[]User) {

	t, err := template.ParseFiles("../templates/users.html")

	if err != nil {
		httpStatusInternalServerError(w, err)
	} else {
		t.Execute(w, users)
	}
}

func renderUser(w http.ResponseWriter, user *User) {

	t, err := template.ParseFiles("../templates/user.html")

	if err != nil {
		httpStatusInternalServerError(w, err)
	} else {
		t.Execute(w, user)
	}
}

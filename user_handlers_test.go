package usermgmt

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func openFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
}

func setupLogger() *logrus.Logger {
	logger := logrus.New()

	logger.Formatter = new(logrus.TextFormatter)
	logger.Level = logrus.DebugLevel

	file, err := openFile("usermgmt.log")
	if err != nil {
		log.Fatal("Cannot open log file.")
		os.Exit(1)
	}
	logger.Out = file
	return logger
}

var users = []User{
	{
		ID:       0,
		Name:     "Neeraj",
		Email:    "abc@xyz.com",
		Password: "abc",
	},
	{
		ID:       1,
		Name:     "Dhiraj",
		Email:    "bcd@xyz.com",
		Password: "bcd",
	},
}

var newUserStr = []byte(`{
	ID:       2,
	Name:     "Neha",
	Email:    "pqr@xyz.com",
	Password: "pqr",
}`)

var updateUserStr = []byte(`{
	ID:       0,
	Name:     "Neeraj Kumar",
	Email:    "pqr@xyz.com",
	Password: "pqr",
}`)

var logger = appLogger{logger: setupLogger()}

type mockDB struct{}

var mockdb = &mockDB{}

var conf = &Configuration{mockdb, logger, true}
var hndlrs = &userHandler{Conf: conf}

func (mockdb *mockDB) GetUsers() (*[]User, error) {
	return &users, nil
}

func (mockdb *mockDB) GetUserByID(id int64) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (mockdb *mockDB) AddUser(user *User) error {
	users = append(users, *user)
	return nil
}

func (mockdb *mockDB) DeleteUserByID(id int64) error {
	for idx, user := range users {
		if user.ID == id {
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (mockdb *mockDB) UpdateUser(usr *User) error {
	for _, user := range users {
		if user.ID == usr.ID {
			user.Email = usr.Email
			user.Name = usr.Name
			user.Password = usr.Password
			return nil
		}
	}
	return errors.New("user not found")
}

func Test_getUsers(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", GetUsersRoute, nil)

	http.HandlerFunc(hndlrs.getUsers).ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("get users failed with response code %v", rec.Code)
	}
}

func Test_getUser(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc(GetUserRoute, hndlrs.getUser)

	req, _ := http.NewRequest("GET", "/user/0", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("get user failed with response code %v", rec.Code)
	}

	idx := 0
	expected := users[idx]
	recieved := new(User)
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &recieved)
	if err != nil {
		panic(err.Error())
	}

	if diff := cmp.Diff(expected, *recieved); diff != "" {
		t.Errorf("GetUser returned User id %v but expected %v", recieved.ID, expected.ID)
	}
}

func Test_addUser(t *testing.T) {

	req, _ := http.NewRequest("POST", AddUserRoute, bytes.NewBuffer(newUserStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	http.HandlerFunc(hndlrs.addUser).ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Add user failed with response code %v", rec.Code)
	}
}

func Test_updateUser(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc(UpdateUserRoute, hndlrs.updateUser)

	req, _ := http.NewRequest("PUT", "/user/0", bytes.NewBuffer(updateUserStr))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("update user failed with response code %v", rec.Code)
	}
}

func Test_deleteUser(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc(UpdateUserRoute, hndlrs.updateUser)

	req, _ := http.NewRequest("DELETE", "/user/0", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("update user failed with response code %v", rec.Code)
	}
}

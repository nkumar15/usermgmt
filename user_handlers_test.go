package usermgmt

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

type mockDB struct{}

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
}

func (db *mockDB) GetUsers() (*[]User, error) {
	return &users, nil
}

func (db *mockDB) GetUserByID(id int64) (*User, error) {
	return nil, nil
}

func (db *mockDB) AddUser(*User) error {
	return nil
}

func (db *mockDB) DeleteUserByID(id int64) error {
	return nil
}

func (db *mockDB) UpdateUser(user *User) error {
	return nil
}

func Test_getUsers(t *testing.T) {

	logger := appLogger{logger: setupLogger()}
	db := &mockDB{}
	conf := &Configuration{db, logger}
	hndlrs := &userHandler{Conf: conf}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", GetUserRoute, nil)

	http.HandlerFunc(hndlrs.getUser).ServeHTTP(rec, req)

}

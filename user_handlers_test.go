package usermgmt

import (
	"errors"
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
	{
		ID:       1,
		Name:     "Dhiraj",
		Email:    "bcd@xyz.com",
		Password: "bcd",
	},
}

func (db *mockDB) GetUsers() (*[]User, error) {
	return &users, nil
}

func (db *mockDB) GetUserByID(id int64) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (db *mockDB) AddUser(user *User) error {
	users = append(users, *user)
	return nil
}

func (db *mockDB) DeleteUserByID(id int64) error {
	for idx, user := range users {
		if user.ID == id {
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (db *mockDB) UpdateUser(usr *User) error {
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

	logger := appLogger{logger: setupLogger()}
	db := &mockDB{}
	conf := &Configuration{db, logger}
	hndlrs := &userHandler{Conf: conf}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", GetUsersRoute, nil)

	http.HandlerFunc(hndlrs.getUsers).ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("get users failed with response code %v", rec.Code)
	}
}

// func Test_getUser(t *testing.T) {

// 	logger := appLogger{logger: setupLogger()}
// 	db := &mockDB{}
// 	conf := &Configuration{db, logger}
// 	hndlrs := &userHandler{Conf: conf}

// 	rec := httptest.NewRecorder()

// 	r := mux.NewRouter()
// 	ts := httptest.NewServer(r)
// 	defer ts.Close()

// 	// Forced to use this approach, Wish setVars in gorilla context gets public
// 	// But still not complete
// 	r.HandleFunc("/user/1", http.HandlerFunc(hndlrs.getUser))

// 	if status := rec.Code; status != http.StatusOK {
// 		t.Errorf("get user failed with response code %v", rec.Code)
// 	}

// 	//idx := 0
// 	//expected := users[idx]
// 	//recieved := new(User)
// 	body, err := ioutil.ReadAll(rec.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	bodyString := string(body)
// 	println("Body :", bodyString)

// 	// err = json.Unmarshal(body, &recieved)
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// if reflect.DeepEqual(recieved, expected) != true {
// 	// 	t.Errorf("Get user returned invaid user %v", expected.ID)
// 	// }

// }

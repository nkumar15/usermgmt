package main

import (
	"errors"
	"log"
	"net/http"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"

	"os"

	"github.com/gorilla/mux"
	um "github.com/nkumar15/usermgmt"
	"github.com/pilu/xrequestid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
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

func connectDB() (sqlbuilder.Database, error) {

	var sqliteSettings = sqlite.ConnectionURL{
		Database: `D:\codes\database\database.sqlite`,
	}

	var db sqlbuilder.Database
	var err error
	db, err = sqlite.Open(sqliteSettings)

	if err != nil {
		return db, errors.New("Couldn't connect database")
	}
	err = db.Ping()
	if err != nil {
		return db, errors.New("Couldn't ping database")
	}
	return db, nil
}

func serveWeb() {

	db, err := connectDB()
	if err != nil {
		log.Fatal("Not able to connect database.", err)
	}

	logger := setupLogger()
	conf := um.NewConfiguration(db, logger)
	router := mux.NewRouter()

	addUserHandler := &um.AppHandler{Conf: conf, Handler: um.AddUserHandler}
	router.Handle(um.AddUserRoute, addUserHandler).Methods("POST")

	getUserHandler := &um.AppHandler{Conf: conf, Handler: um.GetUserHandler}
	router.Handle(um.GetUserRoute, getUserHandler).Methods("GET")

	getUsersHandler := &um.AppHandler{Conf: conf, Handler: um.GetUsersHandler}
	router.Handle(um.GetUsersRoute, getUsersHandler).Methods("GET")

	deleteUserHandler := &um.AppHandler{Conf: conf, Handler: um.DeleteUserHandler}
	router.Handle(um.DeleteUserRoute, deleteUserHandler).Methods("DELETE")

	updateUserHandler := &um.AppHandler{Conf: conf, Handler: um.UpdateUserHandler}
	router.Handle(um.UpdateUserRoute, updateUserHandler).Methods("DELETE")

	n := negroni.New()
	n.Use(xrequestid.New(16))
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":5000", n))
}

func main() {
	serveWeb()
}

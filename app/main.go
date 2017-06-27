package main

import (
	"errors"
	"log"
	"net/http"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/sqlite"

	"os"

	"github.com/gorilla/mux"
	"github.com/nkumar15/usermgmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func openFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
}

func setupLogger(mode string) *logrus.Logger {
	logger := logrus.New()

	if mode == "production" {
		logger.Formatter = new(logrus.JSONFormatter)
		logger.Level = logrus.DebugLevel
	} else {
		logger.Formatter = new(logrus.TextFormatter)
		logger.Level = logrus.InfoLevel
	}

	file, err := openFile("logrus.log")
	if err != nil {
		log.Fatal("cannot open log file.")
		os.Exit(1)
	}
	logger.Out = file
	return logger
}

func connectDB() (sqlbuilder.Database, error) {

	var sqliteSettings = sqlite.ConnectionURL{
		Database: `D:\codes\database\database.sqlite`,
	}

	var pgSettings = postgresql.ConnectionURL{
		Host:     "localhost", // PostgreSQL server IP or name.
		Database: "test",      // Database name.
		User:     "postgres",  // Optional user name.
		Password: "abc123",    // Optional user password.
	}

	var db sqlbuilder.Database
	var err error

	var Server = "sqlite3"

	if Server == "sqlite3" {
		db, err = sqlite.Open(sqliteSettings)
	} else if Server == "pg" {
		db, err = postgresql.Open(pgSettings)
	} else {
		log.Fatal("Invalid database")
	}

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

	logger := setupLogger("development")
	conf := usermgmt.NewConfiguration(db, logger)
	router := mux.NewRouter()

	addUserHandler := &usermgmt.AppHandler{Conf: conf, H: usermgmt.AddUserHandler}
	router.Handle("/user", addUserHandler).Methods("POST")

	getUserHandler := &usermgmt.AppHandler{Conf: conf, H: usermgmt.GetUserHandler}
	router.Handle("/user/{id}", getUserHandler).Methods("GET")

	getUsersHandler := &usermgmt.AppHandler{Conf: conf, H: usermgmt.GetUsersHandler}
	router.Handle("/user", getUsersHandler).Methods("GET")

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":5000", n))
}

func main() {
	serveWeb()
}

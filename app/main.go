package main

import (
	"log"
	"net/http"

	usermgt "github.com/nkumar15/usermgmt"
	"github.com/urfave/negroni"
)

func serveWeb() {
	env := usermgt.Env{}
	router := env.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":5000", n))
}

func main() {
	serveWeb()
}

package main

import (
	"log"
	"net/http"

	baraddur "git.enova.com/go/barad-dur"
)

const (
	// AppName is the application name
	AppName = "velveteen"
)

var (
	// AppRevision is injected by the build pipeline
	AppRevision = "development"
	// AppRevisionTag is injected by the build pipeline
	AppRevisionTag = ""
)

func main() {
	err := baraddur.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	app, err := baraddur.Configure(AppName, AppRevision, AppRevisionTag)
	if err != nil {
		log.Fatal(err)
	}

	app.Log.Infof("Starting %s server on :%s", app.Name(), app.Addr())

	err = app.Serve(setUpRouter())
	if err != nil {
		app.Log.Fatal(err)
	}
}

func setUpRouter() *http.ServeMux {
	router := http.NewServeMux()

	// TODO: add routes

	return router
}

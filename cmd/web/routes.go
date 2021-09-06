package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/calendar/{id:[0-9]+}", app.showCalendar)
	router.HandleFunc("/calendar/{calendar:[0-9]+}/{month:[0-9]+}", app.showMonth)
	http.Handle("/", router)
	return router
}

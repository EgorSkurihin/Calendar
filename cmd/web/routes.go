package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/calendar", app.calendarCreateForm)
	router.HandleFunc("/calendar/{id:[0-9]+}", app.showCalendar)
	router.HandleFunc("/calendar/create", app.createCalendar).Methods("POST")
	router.HandleFunc("/month/{month:[0-9]+}", app.showMonth)

	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	http.Handle("/", router)
	return router
}

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Главная страница</h1>")
}

func (app *Application) showCalendar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.pageNotFound(w)
	}
	response := fmt.Sprintf("<h1>Кадендарь с id = %d</h1>", id)
	fmt.Fprint(w, response)
}

func (app *Application) showMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	calId, err := strconv.Atoi(vars["calendar"])
	if err != nil {
		app.pageNotFound(w)
		return
	}
	monthID, err := strconv.Atoi(vars["month"])
	if err != nil {
		app.pageNotFound(w)
		return
	}
	if monthID < 1 || monthID > 12 {
		app.pageNotFound(w)
		return
	}
	fmt.Fprintf(w, "<h1>Календарь с id %d</h1><h1>Месяц номер %d</h1>", calId, monthID)
}

func (app *Application) pageNotFound(w http.ResponseWriter) {
	http.Error(w, "Страница не найдена - 404", http.StatusNotFound)
}

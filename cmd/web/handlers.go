package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EgorSkurihin/Calendar/pkg/models"
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
	cal, err := app.calendar.GetCalendar(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			fmt.Fprint(w, "<h1>Такого календаря нет!</h1>")
			return
		}
		fmt.Fprint(w, "<h1>Внутрення ошибка сервера!</h1>")
		return
	}
	response := fmt.Sprintf("<h1>Кадендарь: %s</h1><h1>Год: %d</h1>", cal.Name, cal.Year)
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

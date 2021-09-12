package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/EgorSkurihin/Calendar/pkg/models"
	"github.com/gorilla/mux"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	cals, err := app.calendar.LatestCalednars()
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}

	files := []string{
		"./ui/html/home.html",
		"./ui/html/base.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}
	err = ts.Execute(w,
		struct {
			Cals []*models.Calendar
		}{
			cals})
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}
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
			app.pageNotFound(w)
			return
		}
		app.serverError(w, err)
		return
	}
	files := []string{
		"./ui/html/calendar.html",
		"./ui/html/base.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, cal)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *Application) calendarCreateForm(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/cal_create.html",
		"./ui/html/base.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *Application) createCalendar(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}
	title := r.FormValue("title")
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		http.Redirect(w, r, "/calendar", http.StatusTemporaryRedirect)
		return
	}
	id, err := app.calendar.InsertCalendar(title, year)
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}
	app.calendar.FillCalendarByMonths(id)
	url := fmt.Sprintf("/calendar/%d", id)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (app *Application) showMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthID, err := strconv.Atoi(vars["month"])
	if err != nil {
		app.pageNotFound(w)
		return
	}

	month, err := app.calendar.GetMonth(monthID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			app.pageNotFound(w)
			return
		}
		app.serverError(w, err)
		return
	}

	events, err := app.calendar.GetEventsByMonthId(monthID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			app.pageNotFound(w)
			return
		}
		app.errorLog.Fatal(err)
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		description := r.FormValue("description")
		day, err := strconv.Atoi(r.FormValue("day"))
		if err != nil {
			return
		}
		_, err = app.calendar.InsertEvent(title, description, monthID, day)
		if err != nil {
			app.errorLog.Fatal(err)
		}
	}

	files := []string{
		"./ui/html/month.html",
		"./ui/html/base.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}
	err = ts.Execute(w,
		struct {
			Month  models.Month
			Events []*models.Event
		}{
			*month, events})
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

package models

import "errors"

var ErrNoRecords = errors.New("No records in table")

type Calendar struct {
	Id     int
	Name   string
	Year   int
	Months []*Month
}

type Month struct {
	Id          int
	Name        string
	CalendarId  int
	IsAnyEvents bool
}

type Event struct {
	Id          int
	MonthId     int
	Title       string
	Description string
	Day         int
}

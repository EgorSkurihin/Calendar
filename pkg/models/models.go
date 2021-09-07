package models

import "errors"

var ErrNoRecords = errors.New("No records in table")

type Calendar struct {
	Id   int
	Name string
	Year int
}

type Month struct {
	Id         int
	Name       string
	CalendarId int
}

type Event struct {
	Id          int
	MonthId     int
	Title       string
	Description string
	Day         int
}

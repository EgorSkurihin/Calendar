package mysql

import (
	"database/sql"
	"errors"
	"sort"

	"github.com/EgorSkurihin/Calendar/pkg/models"
)

type CalendarModel struct {
	DB *sql.DB
}

func (m *CalendarModel) InsertCalendar(name string, year int) (int, error) {
	stmt := "INSERT INTO Calendar (Name, Year) VALUES (?, ?)"
	result, err := m.DB.Exec(stmt, name, year)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *CalendarModel) GetCalendar(id int) (*models.Calendar, error) {
	stmt := "SELECT * FROM Calendar WHERE Id = ?"
	row := m.DB.QueryRow(stmt, id)
	cal := &models.Calendar{}
	err := row.Scan(&cal.Id, &cal.Name, &cal.Year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecords
		}
		return nil, err
	}
	months, err := m.GetMonthByCalendarID(id)
	if err != nil {
		return nil, err
	}
	cal.Months = months
	return cal, nil
}

func (m *CalendarModel) LatestCalednars() ([]*models.Calendar, error) {
	// Пишем SQL запрос, который мы хотим выполнить.
	stmt := `SELECT Id, Name, Year FROM Calendar ORDER BY Id DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var calendars []*models.Calendar

	for rows.Next() {
		c := &models.Calendar{}
		err = rows.Scan(&c.Id, &c.Name, &c.Year)
		if err != nil {
			return nil, err
		}
		calendars = append(calendars, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return calendars, nil
}

func (m *CalendarModel) InsertMonth(name string, calendarId int) (int, error) {
	stmt := "INSERT INTO Month (Name, calendarId) VALUES (?, ?)"
	result, err := m.DB.Exec(stmt, name, calendarId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *CalendarModel) GetMonth(id int) (*models.Month, error) {
	stmt := "SELECT * FROM Month WHERE id = ?"
	row := m.DB.QueryRow(stmt, id)
	mon := &models.Month{}
	err := row.Scan(&mon.Id, &mon.Name, &mon.CalendarId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecords
		}
		return nil, err
	}
	return mon, nil
}

func (m *CalendarModel) FillCalendarByMonths(calId int) error {
	var months = []string{
		"Январь", "Февраль", "Март", "Апрель",
		"Май", "Июнь", "Июль", "Август",
		"Сентябрь", "Октябрь", "Ноябрь", "Декабрь"}
	for _, month := range months {
		_, err := m.InsertMonth(month, calId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *CalendarModel) GetMonthByCalendarID(calendarId int) ([]*models.Month, error) {
	stmt := "SELECT * FROM Month WHERE calendarId = ?"
	rows, err := m.DB.Query(stmt, calendarId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var months []*models.Month
	for rows.Next() {
		m := &models.Month{}
		err := rows.Scan(&m.Id, &m.Name, &m.CalendarId)
		if err != nil {
			return nil, err
		}
		months = append(months, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	stmt = "SELECT EXISTS(SELECT * FROM Event WHERE MonthId = ?)"
	for _, mth := range months {
		row := m.DB.QueryRow(stmt, mth.Id)
		err = row.Scan(&mth.IsAnyEvents)
		if err != nil {
			return nil, err
		}
	}
	return months, nil
}

func (m *CalendarModel) InsertEvent(title, description string, monthId, day int) (int, error) {
	stmt := "INSERT INTO Event (Title, Description, MonthId, Day) VALUES (?, ?, ?, ?)"
	result, err := m.DB.Exec(stmt, title, description, monthId, day)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *CalendarModel) GetEvent(id int) (*models.Event, error) {
	stmt := "SELECT * FROM Event WHERE Id = ?"
	row := m.DB.QueryRow(stmt, id)
	ev := &models.Event{}
	err := row.Scan(&ev.Id, &ev.Title, &ev.Day, &ev.Description, &ev.MonthId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecords
		}
		return nil, err
	}
	return ev, nil
}

func (m *CalendarModel) GetEventsByMonthId(monthId int) ([]*models.Event, error) {
	stmt := "SELECT * FROM Event WHERE MonthId = ?"
	rows, err := m.DB.Query(stmt, monthId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*models.Event
	for rows.Next() {
		e := &models.Event{}
		err := rows.Scan(&e.Id, &e.Title, &e.Description, &e.MonthId, &e.Day)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	sort.Slice(events, func(i, j int) bool { return events[i].Day < events[j].Day })
	return events, nil
}

package mysql

import (
	"database/sql"

	"github.com/EgorSkurihin/Calendar/pkg/models"
)

type CalendarModel struct {
	DB *sql.DB
}

func (m *CalendarModel) InsertCalendar(name string, year int) (int, error) {
	return 0, nil
}

func (m *CalendarModel) GetCalendar(id int) (*models.Calendar, error) {
	return nil, nil
}

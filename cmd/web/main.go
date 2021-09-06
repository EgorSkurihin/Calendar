package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/EgorSkurihin/Calendar/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	calendar *mysql.CalendarModel
}

func main() {
	addr := flag.String("addr", ":8080", "Сетевой адрес HTTP")
	dsn := flag.String("dsn", "calendar_web:232323@/calendar", "Название MySQL источника данных")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
		calendar: &mysql.CalendarModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.router(),
	}

	app.infoLog.Println("Server is listening...")
	err = srv.ListenAndServe()
	app.errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

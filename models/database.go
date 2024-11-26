package models

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"

)

var db *sql.DB

func InitDB () error {
	
	var dbErr error
	db, dbErr = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		return errors.New("Ошибка при подключении к базе данных: " + dbErr.Error())
	}

		dbErr = db.Ping()
		if dbErr != nil {
			return errors.New("Ошибка при проверке соединения с базой данных:" + dbErr.Error())
			}
			return nil
}
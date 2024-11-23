package main

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var db *sql.DB

func initDB () error {
	
	err := godotenv.Load()
	if err != nil {
		return errors.New("Ошибка при загрузке .env файла: " + err.Error())
	}

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

func saveUser (name, email string) error {
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)" )
	if err != nil {
		return errors.New("Ошибка при подготовке запроса: " + err.Error())
		}
		defer stmt.Close()
		
		_, err = stmt.Exec(name, email)
		if err != nil {
			return errors.New("Ошибка при выполнении запроса: " + err.Error())
		}
		return nil
	}
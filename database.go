package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB () {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=root dbname=Go sslmode=disable")
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:",err)
		}

		sqlStmt := `
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255)
		);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Fatalf("%q: %s\n", err, sqlStmt)
			}
}

func saveUser (name, email string) error {
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)" )
	if err != nil {
		return err
		}
		defer stmt.Close()
		
		_, err = stmt.Exec(name, email)
		if err != nil {
			return err
		}
		return nil
	}
package models

import (
	"errors"

)


func SaveUser (name, email string) error {
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

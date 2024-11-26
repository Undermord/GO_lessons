package handlers

import (
	"fmt"
	"net/http"

	"GO_lessons/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Приветствую")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Это простой пет-проект на Go")
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, "Имя и email не могут быть пустыми", http.StatusBadRequest)
		return
	}

	err := models.SaveUser(name, email)
	if err != nil {
		http.Error(w, "Ошибка сохранения данных", http.StatusInternalServerError)
		return
	}	

	fmt.Fprintf(w, "Имя: %s\nEmail: %s", name, email)

}


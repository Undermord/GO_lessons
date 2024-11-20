package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	initDB()
	
	// Обработчик для корня
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Приветствую")
	})
	// Обработчик для другой страницы
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Это простой пет-проект на Go")

	})

	// Для статистических файлов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Для формы
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
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

		err := saveUser(name, email)
		if err != nil {
			http.Error(w, "Ошибка сохранения данных", http.StatusInternalServerError)
			return
		}	

		fmt.Fprintf(w, "Имя: %s\nEmail: %s", name, email)
	})

	fmt.Println("Сервер запущен на порте 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

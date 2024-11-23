package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	err = initDB()
	if err != nil {
		log.Fatal(err)
	}

	
	r := mux.NewRouter()
	

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Приветствую")
	})
	
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Это простой пет-проект на Go")

	})


	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))



	r.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
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
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

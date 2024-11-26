package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"GO_lessons/config"
	"GO_lessons/handlers"
	"GO_lessons/models"
	
)

func main() {

	config.LoadEnv()

	err :=models.InitDB()
	if err != nil {
		log.Fatal(err)
		}

	r := mux.NewRouter()
	
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/about", handlers.AboutHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/submit", handlers.SubmitHandler)




	fmt.Println("Сервер запущен на порте 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Приветствую")
		})

fmt.Println("Сервер запущен на порте 8080")
http.ListenAndServe(":8080", nil)
}
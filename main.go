package main

import (
	"fmt"
	"log"
	"net/http"
	"golang-todo-api/db"
	"golang-todo-api/routes"
)

func main() {
	db.ConnectDB()

	router := routes.SetupRouter()

	port := ":8080"
	fmt.Println("🔥 Servidor rodando em http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

package main

import (
	"log"
	"net/http"

	_ "github.com/swaggo/http-swagger"

	_ "CsvManipulator/docs"
	"CsvManipulator/handlers"
)

// @title CSV Manipulator API
// @description This is a sample server CSV Manipulator server.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	router := handlers.NewRouter()
	log.Println("Server is listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

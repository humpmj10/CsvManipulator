package main

import (
	"CsvManipulator/services"
	"log"
	"net/http"

	_ "github.com/swaggo/http-swagger" // docs are generated by Swag CLI, you have to import it.

	_ "CsvManipulator/docs" // docs are generated by Swag CLI, you have to import it.
	"CsvManipulator/handlers"
)

// @title CSV Manipulator API
// @description This is a sample server CSV Manipulator server.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	api := handlers.API{CsvService: services.CsvService{}}
	router := api.NewRouter()
	log.Println("Server is listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

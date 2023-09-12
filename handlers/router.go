package handlers

import (
	"github.com/gorilla/mux"
	"net/http"

	_ "CsvManipulator/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/upload", uploadHandler).Methods(http.MethodPost)

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}

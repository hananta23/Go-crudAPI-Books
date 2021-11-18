package router

import (
	"go-crudAPI/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/crudAPI/books", controller.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/crudAPI/books/{id}", controller.GetBook).Methods("GET", "OPTIONS")
	router.HandleFunc("/crudAPI/books", controller.AddBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/crudAPI/books/{id}", controller.UpdateBooks).Methods("PUT", "OPTIONS")
	router.HandleFunc("/crudAPI/books/{id}", controller.DeleteBooks).Methods("DELETE", "OPTIONS")

	return router
}

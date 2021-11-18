package main

import (
	"fmt"
	"go-crudAPI/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Server running on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

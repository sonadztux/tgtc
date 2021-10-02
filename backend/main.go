package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/radityaqb/tgtc/backend/handlers"
	"github.com/radityaqb/tgtc/backend/server"
)

func main() {

	// Init database connection
	// database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	// construct your own API endpoints
	// endpoint : /add-product
	router.HandleFunc("/add-product", handlers.AddProduct).Methods(http.MethodPost)

	// endpoint : /get-product?id=
	router.HandleFunc("/get-product", handlers.GetProduct).Methods(http.MethodGet)

	// endpoint : /update-product
	router.HandleFunc("/update-product", handlers.UpdateProduct).Methods(http.MethodPatch)

	// endpoint : /delete-product
	router.HandleFunc("/delete-product", handlers.DeleteProduct).Methods(http.MethodDelete)
	
	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}

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
	// endpoint : /get-product?id=
	// endpoint : /update-product
	// endpoint : /delete-product

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}

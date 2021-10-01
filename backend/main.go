package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/handlers"
	"github.com/radityaqb/tgtc/backend/server"
)

func ping(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "pong\n")
}

func main() {

	// Init database connection
	database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	router.HandleFunc("/ping", ping).Methods(http.MethodGet)
	router.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", handlers.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/product/insert", handlers.InsertProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/update", handlers.UpdateProduct).Methods(http.MethodPut)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}

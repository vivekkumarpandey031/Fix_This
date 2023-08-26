package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT string

func main() {
	flag.StringVar(&PORT, "port", "8080", "--port 8080 or -port 8080")
	flag.Parse()

	router := mux.NewRouter() //Returns a  new router instance

	//Creating a custum server for more control
	srv := http.Server{
		Addr:         ":" + PORT, //Pass the custum port
		Handler:      router,     //Pass the instance of gorilla mux
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	//Add handlers

	fmt.Println("Server started and listening on port->", PORT)
	srv.ListenAndServe() //starting server

}

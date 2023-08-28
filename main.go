// abhinav OBLLnnzk3BGqwlgl
//"mongodb+srv://abhinav:OBLLnnzk3BGqwlgl@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority"

package main

import (
	"context"
	"golang-project/database"
	//"demo/handlers"
	"flag"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

var (
	dsn="mongodb+srv://vscoproject:victoriasecret@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority"
	PORT   string
)

func main() {
	// Create a new client and connect to the server
	client, err := database.GetConnection(dsn)
	if err != nil {
		panic(err)
	}

	//Disconnect function to disconnect connection after the work is done
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	
	err=database.Ping(client,"demo")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	}

	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")
	flag.Parse()

	router := mux.NewRouter()

	srv := http.Server{
		Addr:         ":" + PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}
	
	srv.ListenAndServe()
}
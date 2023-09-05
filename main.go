package main

import (
	"context"
	"flag"
	"fmt"
	"golang-project/database"
	"golang-project/handlers"
	"golang-project/middleware"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	dsn  = "mongodb://localhost:27017"
	PORT string
)

func init() {
	myDir, _ := os.Getwd()
	flag.Set("logtostderr", "false")
	flag.Set("log_dir", myDir+"/log/dir")
	flag.Parse()
	//Intialize session
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	//Configuring the session
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
	}

}

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

	err = database.Ping(client, "demo")
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

	userDB := new(database.User)
	userDB.Client = client
	userDB.Dbname = "userdb"
	userDB.Collection = "users"
	userHandler := new(handlers.UserHandler)
	userHandler.DB = userDB
	mobileHandler := new(handlers.Mobilehandler)
	mobileDb := new(database.Mobile)
	mobileDb.Client = client
	mobileDb.Dbname = "mobiledb"
	mobileDb.Collection = "mobiles"
	mobileHandler.DB = mobileDb
	laptopHandler := new(handlers.Laptophandler)
	laptopDb := new(database.Laptop)
	laptopDb.Client = client
	laptopDb.Dbname = "laptopdb"
	laptopDb.Collection = "laptops"
	laptopHandler.DB = laptopDb
	serviceHandler := new(handlers.ServiceHandler)
	serviceDB := new(database.Service)
	serviceDB.Client = client
	serviceHandler.DB = serviceDB

	router.Use(middleware.WriteToConsole)
	//Add endpoints
	router.HandleFunc("/user/register", userHandler.Register)
	router.HandleFunc("/user/login", userHandler.Login)
	router.HandleFunc("/user/me", userHandler.Me)
	router.HandleFunc("/user/mobile", mobileHandler.GetAll)
	router.HandleFunc("/user/mobile/addmobileproblem", mobileHandler.AddMobileProblem)
	router.HandleFunc("/user/laptop/addlaptopproblem", laptopHandler.AddLaptopProblem)
	router.HandleFunc("/user/laptop", laptopHandler.GetAll)
	router.HandleFunc("/service/receipt/{id}", serviceHandler.Receipt)
	fmt.Printf("Listening on Port: %s",PORT)
	glog.Infoln(fmt.Sprintf("Listening on Port: %s",PORT))

	srv.ListenAndServe()
}

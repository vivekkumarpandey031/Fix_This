package main

import (
	"fmt"
	"net/http"
	//"github.com/vivekkumarpandey031/mongo-connection"

	"github.com/julienschmidt/httprouter"
	"github.com/vivekkumarpandey031/controllers"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

func main(){

	//se,err:=options.Client().ApplyURI("mongodb")
     
	 r := httprouter.New()
	 uc := controllers.NewUserController(getSession())
	 r.GET("/user/:id",uc.GetUSer)
	 r.POST("/user",uc.CreateUser)
	 r.DELETE("/user/:id",uc.DeleteUser)
	 http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session{

	s,err := mgo.Dial("mongodb://localhost:27017")
    if err !=nil{
		fmt.Println(err)
		//panic(err)
	
	}
	return s
}
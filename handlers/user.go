package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-project/database"
	"golang-project/models"
	"net/http"

	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)



type UserHandler struct {
	DB *database.User
}

/* var (
	Client *mongo.Client
	DbName string
	Collection string
)

func init(){
	dsn:="mongodb+srv://vscoproject:victoriasecret@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1) //which server ur using
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	//Database connection
	a,err:=mongo.Connect(context.TODO(), opts) //mongo.Connect to connect to the database4
	if err!=nil{
		glog.Errorln("Could not connect to database")
		return
	}
	Client=a
	DbName="userdb"
	Collection="users"
} */

func(u *UserHandler)Regsiter(w http.ResponseWriter,r *http.Request){
	if r.Method!="POST"{
		glog.Errorln("Method not implemented")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Method not implemented"))
		return
	}
	user:= new(models.User)
	err:=json.NewDecoder(r.Body).Decode(user)
	
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Data.Please contact admin"))
		return
	}

	err=user.Validate()
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	
	}

	//Crypting the password
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.MinCost)
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot store hashed password"))
	}
	user.Password=string(hashedPassword)
	
	//Inserting into the database
	result,err:=u.DB.Insert(context.TODO(),user)
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result.(primitive.ObjectID).String()))
}

func (u * UserHandler)Login(w http.ResponseWriter,r * http.Request ){
	/* 	//Database connection
	dsn:="mongodb+srv://vscoproject:victoriasecret@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1) //which server ur using
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	client,err:=mongo.Connect(context.TODO(), opts) //mongo.Connect to connect to the database
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}*/

	if r.Method!="POST"{
		glog.Errorln("Method not implemented")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Method not implemented"))
		return
	}

	//Get the user data
	user:=new (models.User)
	err:=json.NewDecoder(r.Body).Decode(user)
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Data.Please contact admin"))
		return
	}

	fmt.Println(user)

	//Authenticate if the data is correct
	//ctx,cancel:=context.WithDeadline(context.Background(),time.Now().Add(10))
	//defer cancel()
	_,err=u.DB.Find(context.TODO(),user)
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Data.Please contact admin"))
		return
	}
	//just testing
	w.Header().Set("Content-type","plain/text")
	w.Write([]byte("Succesful login"))
} 








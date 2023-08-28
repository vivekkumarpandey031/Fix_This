package handlers

import (
	"context"
	"encoding/json"
	"golang-project/database"
	"golang-project/models"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserHandler struct {
	DB *database.User
}

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
		w.Write([]byte("Cannot stoe hashed password"))
	}
	user.Password=string(hashedPassword)
	//Inserting into the database
	result,err:=u.DB.Insert(context.TODO(),user)
	if err!=nil{
		glog.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result.(primitive.ObjectID).String()))
}

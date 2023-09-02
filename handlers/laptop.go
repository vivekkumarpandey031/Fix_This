package handlers

import (
	"context"
	"encoding/json"
	"golang-project/database"
	"net/http"
	"time"

	"github.com/golang/glog"
)

type Laptophandler struct {
	DB *database.Laptop
}

func (l *Laptophandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		glog.Errorln("Method not Implemented")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Method not implemented"))
	}

	session, err := store.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve our session values
	username := session.Values["username"]
	password := session.Values["password"]
	if username == nil || password == nil {
		glog.Errorln("Unauthorized access")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Try Login to get Access"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	documents, err := l.DB.GetAll(ctx)
	if err != nil {
		glog.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something went wrong.Please contact admin"))
		return
	}
	var docs string = ""
	for _, doc := range documents {
		data, err := json.Marshal(doc)
		if err != nil {
			glog.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong.Please contact admin"))
			return
		} else {
			docs += string(data) + ","
		}
	}
	w.Write([]byte(docs[:len(docs)-1]))

}

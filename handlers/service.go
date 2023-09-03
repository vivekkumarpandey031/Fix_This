package handlers

import (
	"context"
	"fmt"
	"golang-project/database"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

type ServiceHandler struct {
	DB *database.Service
}

func (s *ServiceHandler) Receipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		glog.Errorln("Method not implemented")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Method not implemented"))
	}
	id := mux.Vars(r)["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Query the userproblems database for type mobile and user id
	cost1, time1, err := s.DB.TimeAndCost(ctx, id, "mobile")
	if err != nil {
		glog.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return

	}

	//Querty the userproblems database for type laptop and user id
	cost2, time2, err := s.DB.TimeAndCost(ctx, id, "laptop")
	if err != nil {
		glog.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return

	}
	duration := time.Duration((time1 + time2) * 60000000000)

	w.Write([]byte(fmt.Sprintf("The cost is %v and expected delivery is at  %v", cost1+cost2, time.Now().Add(duration))))
}

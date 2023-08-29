package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

func WriteToConsole(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		glog.Infoln(fmt.Sprintf("%v called",r.Method))
		next.ServeHTTP(w,r)

	})
}

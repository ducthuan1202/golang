package middleware

import (
	"net/http"
	"packages/mylog"
)

func AddLogRaquest(next http.Handler) http.Handler {
	handle := func(w http.ResponseWriter, r *http.Request) {		
		mylog.Add(r)
		next.ServeHTTP(w, r)		
	}
	return http.HandlerFunc(handle)
}

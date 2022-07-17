package handler

import (
	"fmt"
	"net/http"
	"time"
)

// NewLoggerWrapper create new logging wrapper over http handler.
func NewLoggerWrapper(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		t := time.Now()

		if h != nil {
			h.ServeHTTP(rw, r)
		}

		fmt.Printf("%s - [%s] %s req: %s\n", t.Format("2 Jan 2006 15:04:05"), time.Since(t), r.Method, r.URL)
	})
}

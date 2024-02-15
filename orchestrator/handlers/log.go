package handlers

import (
	"orchestrator/logging"
	"net/http"
	"io"
	"fmt"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		logs.ReportAction(fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, string(body)))
		next.ServeHTTP(w, r)
	})
}


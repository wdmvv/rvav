package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"orchestrator/logging"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusInternalServerError)
			return
		}
		dst := &bytes.Buffer{}
		json.Compact(dst, body) // I should handle error properly here, but it will be handled later anyway (json decoders)
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		logging.ReportAction(fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, dst.String()))
		next.ServeHTTP(w, r)
	})
}

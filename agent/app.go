package main

// endpoint bindings
import (
	"agent/logging"
	"fmt"
	"net/http"
)

func StartServer(port int) {
	mux := http.NewServeMux()
	status := http.HandlerFunc(StatusHandler)
	eval := http.HandlerFunc(EvalHandler)
	workers := http.HandlerFunc(WorkHandler)

	mux.Handle("/status", loggingMiddleware(status))
	mux.Handle("/eval", loggingMiddleware(eval))
	mux.Handle("/workers", loggingMiddleware(workers))

	logging.ReportAction(fmt.Sprintf("started agent on %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		logging.ReportErr("error ocurred on agent", err)
	}
}

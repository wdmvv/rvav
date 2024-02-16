package app

//basically for binding endpoints not in main

import (
	"fmt"
	"net/http"
	"orchestrator/handlers"
	logs "orchestrator/logging"
)

var port int = 8000
var mux *http.ServeMux

func BindEndpoints() {
	mux = http.NewServeMux()

	chtime := http.HandlerFunc(handlers.ChtimeHandler)
	timeouts := http.HandlerFunc(handlers.TimeoutsHandler)
	status := http.HandlerFunc(handlers.StatusHandler)

	mux.Handle("/chtime", handlers.LoggingMiddleware(chtime))
	mux.Handle("/timeouts", handlers.LoggingMiddleware(timeouts))
	mux.Handle("/status", handlers.LoggingMiddleware(status))
	
}

func StartServer() {
	logs.ReportAction(fmt.Sprintf("started agent on %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		logs.ReportErr("error ocurred on agent", err)
	}
}

package app

//basically for binding endpoints not in main

import (
	"fmt"
	"net/http"
	"orchestrator/handlers"
	"orchestrator/logging"
)

var mux *http.ServeMux

func StartServer(port int) {
	handlers.StartJobs()
	mux = http.NewServeMux()

	chtime := http.HandlerFunc(handlers.ChtimeHandler)
	timeouts := http.HandlerFunc(handlers.TimeoutsHandler)
	status := http.HandlerFunc(handlers.StatusHandler)
	addexpr := http.HandlerFunc(handlers.AddExprHandler)
	jobs := http.HandlerFunc(handlers.JobsHandler)

	mux.Handle("/chtime", handlers.LoggingMiddleware(chtime))
	mux.Handle("/timeouts", handlers.LoggingMiddleware(timeouts))
	mux.Handle("/status", handlers.LoggingMiddleware(status))
	mux.Handle("/addexpr", handlers.LoggingMiddleware(handlers.JobsLogger(addexpr)))
	mux.Handle("/jobs", handlers.LoggingMiddleware(jobs))

	logging.ReportAction(fmt.Sprintf("started orchestrator on %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		logging.ReportErr("error ocurred on orchestrator", err)
	}
}

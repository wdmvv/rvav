package main

// endpoint bindings
import (
	"net/http"
	"agent/logging"
	"fmt"
)


func StartServer(port int){
	mux := http.NewServeMux()
	status := http.HandlerFunc(StatusHandler)
	eval := http.HandlerFunc(EvalHandler)
	workers := http.HandlerFunc(WorkHandler)

	mux.Handle("/status", loggingMiddleware(status))
	mux.Handle("/eval", loggingMiddleware(eval))
	mux.Handle("/workers", loggingMiddleware(workers))

	logs.ReportAction(fmt.Sprintf("started agent on %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		logs.ReportErr("error ocurred on agent", err)
	}
}
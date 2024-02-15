package main

// main load balancer
import (
	"agent/logging"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/semaphore"
)

var Workers int
var Limit semaphore.Weighted

func main() {
	logs.LoggerSetup()

	env := os.Getenv("MAX_WORKERS")
	workers, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		logs.ReportAction("did not find env MAX_WORKERS, setting default 10")
		workers = 10
	}
	Limit = *semaphore.NewWeighted(workers)

	mux := http.NewServeMux()
	status := http.HandlerFunc(statusHandler)
	eval := http.HandlerFunc(evalHandler)

	mux.Handle("/status", loggingMiddleware(status))
	mux.Handle("/eval", loggingMiddleware(eval))

	port := 8081
	logs.ReportAction(fmt.Sprintf("started agent on %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		logs.ReportErr("error ocurred on agent", err)
	}
}

package main

// main load balancer
import (
	"fmt"
	"golang.org/x/sync/semaphore"
	"net/http"
	"os"
	"strconv"
)

var Workers int
var Limit semaphore.Weighted

func main() {
	env := os.Getenv("MAX_WORKERS")
	workers, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		workers = 10
	}
	Limit = *semaphore.NewWeighted(workers)

	mux := http.NewServeMux()
	status := http.HandlerFunc(statusHandler)
	eval := http.HandlerFunc(evalHandler)
	mux.Handle("/status", status)
	mux.Handle("/eval", eval)
	fmt.Println("Started agent on 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

// mux := http.NewServeMux()
// 	helloHandler := http.HandlerFunc(HelloHandler)
// 	mux.Handle("/", Sanitize(SetDefaultName(helloHandler)))

// 	if err := http.ListenAndServe(":8080", mux); err != nil {
// 		panic(err)
// 	}

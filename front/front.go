package main

import (
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	expr := http.HandlerFunc(ExprHandler)
	time := http.HandlerFunc(TimeHandler)
	jobs := http.HandlerFunc(JobsHandler)
	m := http.HandlerFunc(MainHandler)
	
	mux.Handle("/expr", CORS(expr))
	mux.Handle("/time", CORS(time))
	mux.Handle("/jobs", CORS(jobs))
	mux.Handle("/", CORS(m))

	http.ListenAndServe(":7999", mux)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
        // w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// if r.Method == "OPTIONS" {
		// 	w.WriteHeader(http.StatusOK)
		// }
		next.ServeHTTP(w, r)
	})
}

func ExprHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/expr.html")
}

func TimeHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/time.html")
}

func JobsHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/jobs.html")
}

func MainHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/main.html")
}
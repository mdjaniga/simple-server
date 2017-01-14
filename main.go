package main

import "net/http"

var chttp = http.NewServeMux()

func serveIt(w http.ResponseWriter, r *http.Request) {
	chttp.ServeHTTP(w, r)
}

func main() {
	chttp.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", serveIt)
	http.ListenAndServe(":8000", nil)
}

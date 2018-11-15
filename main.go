package main

import (
	"fmt"
	"net/http"
)

func tony(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tony's new go web app")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tony's new go web app")
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
	}
	http.HandleFunc("/tony", tony)
	http.HandleFunc("/hello", hello)
	server.ListenAndServe()
}

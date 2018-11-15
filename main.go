package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

// anything that has ServeHTTP is an handler
func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tony's new go web app %v", r.URL.Path[1:])

}

func main() {
	handler := myHandler{}
	
	// Here we defind a default handler which will be the same
	// all the page you visit
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}

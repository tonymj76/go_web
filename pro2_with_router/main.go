package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello world %s\n", p.ByName("name"))
}

type user struct {
	name string
}

func (t *user) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "what is your name: %s\n", t.name)
}

func main() {
	userName := &user{
		name: "Jerry",
	}

	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	http.Handle("/", userName)

	server := http.Server{
		Addr:    "127.0.0.1:8000",
	}
	server.ListenAndServe()
}

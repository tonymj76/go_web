package main

import (
	"fmt"
	"net/http"
)

func form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "the body %s--\n url: %s\n-- form : %s", r.Body, r.URL, r.PostForm.Get("post"))
}

func main() {
	http.HandleFunc("/", form)
	mux := http.Server{
		Addr: "127.0.0.1:8080",
	}
	mux.ListenAndServe()
}

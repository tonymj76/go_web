package main

import (
	"html/template"
	"net/http"
)

func parseHTML(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templ.html")
	t.Execute(w, "hello world")
}

func main() {
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/helo", parseHTML)
	mux.ListenAndServe()
}

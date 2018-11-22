package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User   string
	Threat []string
}

func rWriter(w http.ResponseWriter, r *http.Request) {
	data := `<html>
		<head><title>Go web Programing</title></head>
		<body><div><p>this is the body of the webpage</p></div></body>
		<html>
		`
	w.Write([]byte(data))
}

func rWriterHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(504)
	fmt.Fprintf(w, "this isn't working fine")
}

func rHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.google.com")
	w.WriteHeader(302)
}

func jsonWriter(w http.ResponseWriter, r *http.Request) {
	p := &Post{
		User:   "Anthony",
		Threat: []string{"first", "last", "something else"},
	}
	data, _ := json.Marshal(p)
	w.Write(data)
}
func jsonWriter2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(200)
	p := &Post{
		User:   "Anthony",
		Threat: []string{"first", "last", "something else"},
	}
	data, _ := json.Marshal(p)
	w.Write(data)
}

func main() {
	serverMux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/writer", rWriter)
	http.HandleFunc("/writeHeader", rWriterHeader)
	http.HandleFunc("/redirect", rHeader)
	http.HandleFunc("/json", jsonWriter)
	http.HandleFunc("/json2", jsonWriter2)
	serverMux.ListenAndServe()
}

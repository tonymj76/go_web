package main

import (
	"net/http"
	"./register"
)

func main() {
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/signup", register.SignUPHander)
	mux.ListenAndServe()
}
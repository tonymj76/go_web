package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)
type giveName struct {
	name string
}

func (g *giveName) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "my name is %s", g.name)
}

func logGiveName(h http.Handler) http.Handler{
	return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request)  {
		fmt.Printf("this is the func name %T\n", h)
		h.ServeHTTP(w,r)
	})
}

func tony(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yes i thing am getting this 's new go web app")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tony's new go web app")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler version call " + name)
		h(w, r)
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	gName := &giveName{name:"Jerry"}
	http.Handle("/give/", logGiveName(gName))
	
	tj := http.HandlerFunc(tony)
	http.Handle("/", tj)
	http.HandleFunc("/hello", log(hello))
	//http.HandleFunc("/", log(tony))
	server.ListenAndServe()
}

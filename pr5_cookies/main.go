package main

import (
	"fmt"
	"net/http"
)

func setCookies(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: 		"first_cookie",
		Value: 		"Go Web Programming",
		HttpOnly: 	true,
	}
	c2 := http.Cookie{
		Name: 		"second_cookie",
		Value: 		"Build2day enterprice",
		HttpOnly:	true,
	}
	c3 := http.Cookie{
		Name: 		"third_cookie",
		Value: 		"Build2day enterprice",
		HttpOnly:	true,
	}
	c4 := http.Cookie{
		Name: 		"fouth_cookie",
		Value: 		"Build2day enterprice",
		HttpOnly:	true,
	}
	// Either we do this below 
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())

	// or simple this
	http.SetCookie(w, &c3)
	http.SetCookie(w, &c4)
	// it doesn't make too much of a difference, though you should take note tht you need to
	// pass in the cookies by reference instead.
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// r.Header.Get() will get the string of all the cookies using ["cookie"] will get a slice of cookies
	getC := r.Header.Get("Cookie")
	// or you can use r.Cookie() to get the value of the cookie that you need
	co1, err:= r.Cookie("first_cookie")
	if err != nil{
		fmt.Fprintln(w,"there is a problem")
	}
	col2:= r.Cookies()
	w.Write([]byte(getC))
	fmt.Fprintf(w, "%s\n", getC)
	fmt.Fprintf(w, "co1: %s\n", co1)
	fmt.Fprintf(w, "col2: %s\n", col2)
}

func main() {
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/set_cookie", setCookies)
	http.HandleFunc("/get-cookie", getCookie)
	mux.ListenAndServe()
}
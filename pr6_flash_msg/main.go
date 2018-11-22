// Flash message with Cookie

package main

import (
	"time"
	"encoding/base64"
	"net/http"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	flashMsg := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	// w.Header().Set("Set-Cookie", flashMsg.String())
	http.SetCookie(w, &flashMsg)
}

func showMessage(w http.ResponseWriter, r *http.Request)  {
	msg, err := r.Cookie("flash")
	if err != nil{
		if err == http.ErrNoCookie{
			w.Write([]byte("Message not found\n"))
		}
		// return kill the func so that rc won't run or just put rc in an else statement
		return
	}
	rc := http.Cookie{
				Name: "flash",
				MaxAge: -1,
				Expires: time.Unix(1, 0),
			}
	http.SetCookie(w, &rc)
	flash, _ := base64.URLEncoding.DecodeString(msg.Value)
	w.Write([]byte(flash))
}

func main() {
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/show-msg", showMessage)
	http.HandleFunc("/set-msg", setMessage)
	mux.ListenAndServe()
}
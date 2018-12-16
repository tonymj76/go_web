package main

import (
	"os"
	"strings"
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"testing"
)

var (
	mux	*http.ServeMux
	w		*httptest.ResponseRecorder
)
func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	w = httptest.NewRecorder()
}

func TestWebServer(t *testing.T) {
	r, _ := http.NewRequest("GET", "/post/6", nil)
	mux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Errorf("connection broken %v \n", w.Code)
	}

	var post Post
	err := json.Unmarshal(w.Body.Bytes(), &post)
	if err != nil {
		t.Error(err)
	}
	if post.Author != "Tony" {
		t.Error("Cannot retrieve Json post")
	}
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader( `{"content":"write put request", "author":"Tony"}`)
	r, err := http.NewRequest("PUT", "/post/6", json)
	if err != nil{
		t.Error(err)
	}
	mux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Error("not successful ", w.Code)
	}
}

func TestHandlePost(t *testing.T) {
	json := strings.NewReader( `{"content":"write post request", "author":"thomson"}`)
	r, _ := http.NewRequest("POST", "/post/", json)
	mux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Error("not successful")
	}
	if w.Header().Get("Content-Type") != "application/json"{
		t.Error(w.Header().Get("Content-Type"))
	}
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
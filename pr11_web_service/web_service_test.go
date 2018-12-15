package main

import (
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"testing"
)

func TestWebServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)	
	write := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/6", nil)
	mux.ServeHTTP(write, request)
	if write.Code != 200 {
		t.Errorf("connection broken %v \n", write.Code)
	}

	var post Post
	err := json.Unmarshal(write.Body.Bytes(), &post)
	if err != nil {
		t.Error(err)
	}
	if post.Author != "jerry" {
		t.Error("Cannot retrieve Json post")
	}
}
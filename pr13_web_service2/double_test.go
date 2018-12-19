package main

import (
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"testing"
)
// FakePost is emurating Post
type FakePost struct {
	ID				int
	Comment		string
	Author 		string 	
	Msg 			string
}

func (post FakePost)Fetch(id int, p interface{}) {
	post.ID = id
	jsonf := `{"id":1, "comment":"funny", "author":"stackoverflow", "text":"i can't"}`
	json.Unmarshal([]byte(jsonf), &p)
}

func (post FakePost) Createdb(p interface{}) {
	return
}

func (post FakePost)Delete(id int, p interface{}) {
	return
}
func (post FakePost)Update(id int, p ...interface{}) {
	return
}

func TestWebservice(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/1", RequestHandler(FakePost{}))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(w,r)
	if w.Code != 200 {
		t.Errorf("connection broken %v \n", w.Code)
	}

	var post Post
	if err := json.Unmarshal(w.Body.Bytes(), &post); err != nil{
		t.Error(err)
	}
	
	if post.ID != 1 {
		t.Error("Cannot retrieve Json post")
	}
}

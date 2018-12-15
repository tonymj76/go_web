package main

import (
	"time"
	"testing"
)
// go test -v -cover -short -parallel 3 -bench .

func TestDecoder(t *testing.T) {
	post, err:= decoder("jsonfile.json")
	if err != nil {
		t.Error(err)
	}
	if post.ID!= 1 {
		t.Error("Wrong id, was expecting 1 but got", post.ID)
	}
	
}

func TestEncode(t *testing.T) {
	if testing.Short(){

		t.Skip("skipping encoding for now")
	}
	time.Sleep(10*time.Second)
}


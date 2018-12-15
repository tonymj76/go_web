package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"encoding/json"
	"os"
	"time"
)


// Post to get json data
type Post struct {
	ID			int			`json:"ID"`
	Author	Author		`json:"author"`
	Text		string		`json:"msg"`
	Comment	[]Comment	`json:"comment"`
	Time		time.Time	`json:"-"`
}

// Comment is the post comments
type Comment struct {
	ID 		int		`json:"ID"`
	Author	Author	`json:"author"`
	Text		string	`json:"msg"`
	Time		time.Time `json:"-"`
}
// Author is one to one relationship
type Author struct {
	ID			int		`json:"ID"`
	Name		Name	`json:"name"`
}

// Name is one to one relationship
type Name struct {
	FName			string		`json:"first_name"`
	LName			string   	`json:"last_name"`
}

func decoder(filename string) (post Post, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil{
		log.Println(err)
		return
	}
	err = json.NewDecoder(file).Decode(&post)
	return
}

func unmarshaling(filename string) (post Post, err error) {
	/* file, err := os.Open(filename)
	if err != nil{
		return
	}
	r, err := ioutil.ReadAll(file)
	err = json.Unmarshal(r, &post) */
	// 
	var file []byte
	file, err = ioutil.ReadFile(filename)
	err = json.Unmarshal(file, &post)
	return
}
func main() {
	post, err:= decoder("jsonfile.json")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%v\n", post)

	p, _ := unmarshaling("jsonfile.json")
	fmt.Println(p.Comment[1])
}


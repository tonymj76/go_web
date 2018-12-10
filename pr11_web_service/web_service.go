package main

import (
	"encoding/json"
	"strconv"
	"path"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	
)


// Post of the blog
type Post struct {
	ID			int		`json:"_" gorm:"index"`
	Content	string	`json:"content" gorm:"not null"`
	Author	string	`json:"author" gorm:"not null;unque"`
}

// DB database
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "dbname=webservice_post user=tonymj password=t sslmode=disable")
	if err !=nil {
		panic(err)
	}
	DB.AutoMigrate(&Post{})
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w,r)
	}
	if err != nil{
		log.Fatalln(err)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	var (
		post Post
		id int
	)
	id, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	DB.Where("id = ?", id).First(&post)
	output, err := json.MarshalIndent(&post, " ", "\t")
	if err != nil{
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	var post Post
	err = json.NewDecoder(r.Body).Decode(&post)
	DB.Create(&post)
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	var id int
	id, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	var post Post
	var newPost Post
	DB.Where("id = ?", id).First(&post)
	err = json.NewDecoder(r.Body).Decode(&newPost)
	DB.Model(&post).Update(newPost)
	// try this
	//DB.Model(&post).Where("id= ?",id).Update(newPost)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	var id int
	var post Post
	id, err = strconv.Atoi(path.Base(r.URL.Path))
	DB.Where("id = ?", id).Delete(&post)
	return
}

func main() {
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/post/", handleRequest)
	mux.ListenAndServe()
}
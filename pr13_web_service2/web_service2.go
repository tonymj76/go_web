package main

import (
	"encoding/json"
	"path"
	"strconv"
	"net/http"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Post one to more to comment
type Post struct {
	ID				int		`json:"id" gorm:"index"`
	Comment		string 	`json:"comment" gorm:"not null"`
	//Comment		[]Comment `json:"comment" gorm:"not null;foreignkey:Comments"`
	Author 		string 	`json:"author" gorm:"not null"`
	Msg 			string 	`json:"text" gorm:"not null"`
	DB				*gorm.DB	`json:"-" gorm:"-"`
}

// // Comment for post
// type Comment struct {
// 	ID				int	 `json:"id" gorm:"index"`
// 	Author		string `json:"author" gorm:"not null"`
// 	Msg			string `json:"text" gorm:"not null"`
// 	Comments		uint	 `json:"_"`
// }



// CRUD functions

// Createdb is use to add post to db
func (post *Post) Createdb(p *Post) {
	post.DB.Create(&p)
}

// Fetch is use to query a db with id
func (post *Post) Fetch(id int, p *Post) {
	post.DB.Where("id = ?", id).First(&p)
}

// Update is use to update db
func (post *Post) Update(id int, p ...*Post) {
	post.DB.Model(&p[0]).Update(&p[1])
}

// Delete is use to delete
func (post *Post) Delete(id int, p *Post) {
	post.DB.Where("id = ?", id).Delete(&p)
}

// Test is interface
type test interface {
	Createdb(*Post)
	Fetch(int, *Post)
	Update(int, ...*Post)
	Delete(int, *Post)
}

// RequestHandler is used to update the mux
func RequestHandler(t test) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGet(w http.ResponseWriter,r *http.Request, p test) (err error) {
	var (
		id int
		jsonPost []byte
	)

	id, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	var post Post
	p.Fetch(id, &post)
	jsonPost, err = json.MarshalIndent(&post, " ", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPost)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, p test) (err error) {
	var post Post
	err = json.NewDecoder(r.Body).Decode(&post)
	p.Createdb(&post)
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, p test) (err error) {
	var (
		post Post
		newPost Post
	)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	p.Fetch(id, &post)
	err = json.NewDecoder(r.Body).Decode(&newPost)
	p.Update(id, &post, &newPost)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, p test) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	var post Post
	p.Delete(id, &post)
	return
}

func main() {
	DB, err := gorm.Open("postgres", "dbname=webservice2 user=tonymj password=t sslmode=disable")
	if err != nil {
		log.Fatalln("unable to open database", err)
	}
	DB.AutoMigrate(&Post{})
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/post/", RequestHandler(&Post{DB:DB}))
	mux.ListenAndServe()
}

// json: cannot unmarshal object into Go struct field Post.comment of type []main.Comment
// i think is due to Association and related
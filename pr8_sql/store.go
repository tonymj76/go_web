package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

var DB *sql.DB

// connecting with DB
func init() {
	var err error
	DB, err = sql.Open("postgres", "user=tonymj dbname=gwp password=t sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := DB.Query("SELECT id, content, author FROM posts LIMIT $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

// GetPost Gets a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("SELECT id, content, author FROM posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create creates a new post
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update updates a post
func (post *Post) Update() (err error) {
	_, err = DB.Exec("UPDATE posts SET content = $2, author = $3 WHERE id = $1", post.ID, post.Content, post.Author)
	return
}

//Delete delets a post
func (post *Post) Delete() (err error) {
	_, err = DB.Exec("DELETE FROM posts WHERE id = $1", post.ID)
	return
}

func main() {
	post := Post{Content: "Hello world", Author: "tony"}
	post2 := Post{Content: "Hello tom", Author: "tom"}
	post3 := Post{Content: "Hello adr", Author: "adr"}
	post4 := Post{Content: "Hello joy", Author: "joy"}
	post.Create()
	post2.Create()
	post3.Create()
	post4.Create()
	p, _ := Posts(3)
	for _, x := range p {
		fmt.Println(x)
	}
}

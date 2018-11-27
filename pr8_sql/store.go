package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Post struct {
	ID       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	CommentID int
	Content   string
	Author    string
	P         *Post
}

// DB is the database
var DB *sql.DB

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

// connecting with DB
func init() {
	file, errs := os.OpenFile("gwp.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.SetPrefix("TRACE")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if errs != nil {
		panic(errs)
	}
	defer file.Close()

	var err error
	DB, err = sql.Open("postgres", "user=tonymj dbname=gwp password=t sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// GetPosts return query given by limit and go tro each row and asign it to the struct
func GetPosts(limit int) (posts []Post, err error) {
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
	post.Comments = []Comment{}
	err = DB.QueryRow("SELECT id, content, author FROM posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)

	rows, err := DB.Query("SELECT commentsid, content, author FROM comments")
	defer rows.Close()
	check(err)
	for rows.Next() {
		comm := Comment{P: &post}
		err = rows.Scan(&comm.CommentID, &comm.Content, &comm.Author)
		check(err)
		post.Comments = append(post.Comments, comm)
	}
	return
}

// Create creates a new post
func (post *Post) Create() (err error) {
	// return id here returns the id in the sql and adds it to the post.id
	// this return the id of this particular insert which is den added to post.ID
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	// .Scan takes the returned id and assign it to addr of post.ID
	// where post.Content and Author are the value we are adding to the db
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

func (post Post) String() string {
	return fmt.Sprintf("%v: %s \t%s\n", post.ID, post.Author, post.Content)
}

// Create Comment
func (comment *Comment) Create() (err error) {
	if comment.P == nil {
		err = errors.New("Post not found")
		return
	}
	err = DB.QueryRow("INSERT INTO comments (content, author, post_id) VALUES ($1, $2, $3) returning commentsId", comment.Content, comment.Author, comment.P.ID).Scan(&comment.CommentID)
	check(err)
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
	c1 := Comment{Content: "Good Post", Author: "Joe", P: &post}
	c1.Create()
	c2 := Comment{Content: "Good Post", Author: "jerry", P: &post2}
	c2.Create()
	c3 := Comment{Content: "Good Post", Author: "bob", P: &post3}
	c3.Create()
	rd, _ := GetPost(4)
	fmt.Println(rd)
	fmt.Println(rd.Comments)
	fmt.Println()
	fmt.Println(rd.Comments[0].P)
	fmt.Println(rd.Comments[0].Author)
}

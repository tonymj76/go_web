package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	gorm.Model
	ID       int
	Content  string `gorm:"not null"`
	Author   string `gorm:"not null"`
	Comments []Comment
}

type Comment struct {
	CommentID int
	Content   string `gorm:"not null"`
	Author    string
	PostID    int `gorm:"index"`
	CreatedAt time.Time
}

// DB is the database
var DB *gorm.DB

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

// connecting with DB
func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=tonymj dbname=gwp password=t sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello world", Author: "tony"}

	comment := Comment{Content: "Good post", Author: "Sau"}
	DB.Model(&post).Association("Comments").Append(comment)
	DB.Create(&post)

	var readPost Post
	DB.Where("author = $1", "tony").First(&readPost)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	var comments []Comment
	DB.Model(&readPost).Related(&comments)
	fmt.Println(comments)

}

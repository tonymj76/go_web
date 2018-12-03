package register

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	// this is postgres driver init
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SignUP registeration form for the user
type SignUP struct {
	gorm.Model
	User      string `json:"User" gorm:"not null;unique"`
	FirstName string `json:"FirstName" gorm:"not null"`
	LastName  string `json:"LastName" gorm:"not null"`
	Email     string `json:"Email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "dbname=test_signup_user user=tonymj sslmode=disable password=t")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&SignUP{})
}

// SignUPHander gets the json from the request and store it in the struct
func SignUPHander(w http.ResponseWriter, r *http.Request) {
	signup := &SignUP{}
	err := json.NewDecoder(r.Body).Decode(signup)
	if err != nil {
		log.Fatalln(err)
	}
	hashPass, err :=
		Db.Create(signup)
}

// Login is a hander that authenticate the user
func Login(w http.ResponseWriter, r http.Request) {
	type login struct {
		Password string `json:"user"`
		Email    string `json:"email"`
	}

}

package data

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=telehealth-app sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db open")
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

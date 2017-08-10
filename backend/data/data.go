package data

import (
	"_githubcom/lib/pq"
	"data/base/sql"
	"log"
)

var Db *sql.DB

func main() {
	var err error
	Db, err = sql.open("postgres", "dbname=previ-app, sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

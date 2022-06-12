package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Tiger"
	dbname   = "postgres"
)

var db *sqlx.DB

func main() {
	http.HandleFunc("/getusers", getUsers)
	http.ListenAndServe(":8550", nil)

}
func OpenConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println(psqlInfo)
	fmt.Println("Successfully connected!")
	return db
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	OpenConnection()
}

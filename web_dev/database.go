package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	db_user  = "postgres"
	password = "Tiger"
	dbname   = "postgres"
)

type user struct {
	UserId   int    `json:"user_id" db:"user_id"`
	FullName string `json:"full_name" db:"full_name"`
	Age      int    `json:"age" db:"age"`
	Address  string `json:"address" db:"address"`
}

type person struct {
	PersonId int    `json:"person_id" db:"person_id"`
	Name     string `json:"name" db:"name"`
	Number   int    `json:"number" db:"number"`
	Address  string `json:"address" db:"address"`
}

//var db *sqlx.DB

func main() {
	http.HandleFunc("/getusers", getUsers)
	http.HandleFunc("/createuser", createUser)
	http.HandleFunc("/updateuser", updateuser)
	http.HandleFunc("/deleteuser", deleteuser)
	http.HandleFunc("/createUniqueUser", createUniqueUser)

	fmt.Println("connected")
	http.ListenAndServe(":8550", nil)

}
func OpenConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, db_user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//	fmt.Println(psqlInfo)
	fmt.Println("Successfully connected!")
	return db
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var User []user

	str := "SELECT * FROM users" // query string

	err := db.Select(&User, str) //  Database -> Server -> Client
	if err != nil {
		fmt.Print("err", err)
	}
	byteUser, _ := json.Marshal(&User)

	w.Write(byteUser)
}

func createUser(w http.ResponseWriter, r *http.Request) { // Client -> Server -> Database
	db := OpenConnection()

	var User user

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		panic(err)
	}

	str := "INSERT INTO users VALUES (:user_id,:full_name,:age,:address)"

	_, err = db.NamedExec(str, &User)
	if err != nil {
		fmt.Print("err", err)
	}

	fmt.Fprintln(w, "data has been persited sucessfully")

}

func updateuser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var User user
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		fmt.Println("err-", err)
	}
	str := `UPDATE users SET full_name=$1 WHERE user_id=$2`

	_, err = db.Exec(str, User.FullName, User.UserId)
	if err != nil {
		fmt.Println("err--", err)
	}
	fmt.Fprintln(w, "Data has been updated successfully")
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	Uid := r.URL.Query().Get("user_id")
	str := `DELETE FROM users WHERE user_id=$1`

	_, err := db.Exec(str, Uid) //here Uid is used to store the value of user_id
	if err != nil {
		fmt.Println("err--", err)
	} else {
		fmt.Fprintln(w, "Data has been deleted successfully")
	}
}

func createUniqueUser(w http.ResponseWriter, r *http.Request) { // AutoIncreement of ID
	db := OpenConnection()
	var Person person

	err := json.NewDecoder(r.Body).Decode(&Person)
	if err != nil {
		fmt.Println("err--", err)
	}

	// INSERT INTO person(name,) VALUES(:fields...)
	str := `INSERT INTO person(name,number,address) VALUES (:name,:number,:address)`

	_, err = db.NamedExec(str, &Person)
	if err != nil {
		fmt.Println("err--", err)
	} else {
		fmt.Fprintln(w, "Inserted into db")
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

type User struct {
	UserId   int    `json:"user_id" db:"user_id"`
	FullName string `json:"full_name" db:"full_name"`
	Age      int    `json:"age" db:"age"`
	Address  string `json:"address" db:"address"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getuser", GetUser).Methods("GET")
	r.HandleFunc("/single/{user_id}", SingleUser).Methods("GET") // fetching single user as per user_id requested by the Client
	// r.HandleFunc("/single/{user_id:[0-9]+}", SingleUser).Methods("GET")
	r.HandleFunc("/createuser", CreateUser).Methods("POST")
	r.HandleFunc("/updateuser/{full_name}/{user_id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/deleteuser/{user_id}", DeleteUser)

	fmt.Println("Connected...")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var user []User
	str := `SELECT * FROM users`

	err := db.Select(&user, str)
	if err != nil {
		fmt.Println(err.Error())
		panic("There's some issue with connection!!")
	}

	j, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "Db->Server->Client")
		w.Write(j)
	}
}

func SingleUser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	vars := mux.Vars(r)

	var user []User
	str := `SELECT * FROM users WHERE user_id=$1`

	err := db.Select(&user, str, vars["user_id"])
	if err != nil {
		fmt.Println(err.Error())
		panic("Please try again!!")

	}
	j, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "Db->Server->Client")
		w.Write(j)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	str := `INSERT INTO users VALUES(:user_id,:full_name,:age,:address)`
	_, err = db.NamedExec(str, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write([]byte("Data has been inserted"))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	vars := mux.Vars(r)

	str := `UPDATE users SET full_name=$1 WHERE user_id=$2`

	_, err := db.Exec(str, vars["full_name"], vars["user_id"])
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write([]byte("Data has been Updated"))
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	vars := mux.Vars(r)
	str := `DELETE FROM users WHERE user_id=$1`

	_, err := db.Exec(str, vars["user_id"])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "Data has been deleted")
	}

}

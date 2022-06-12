package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	http.HandleFunc("/name", handleIt)
	fmt.Println("Connected")
	http.ListenAndServe(":8086", nil)
}

func handleIt(w http.ResponseWriter, r *http.Request) {
	// Decoding from request URL
	Name := r.URL.Query().Get("name")
	//	Title := r.URL.Query().Get("title")

	fmt.Println("Printing Name :", Name)

	FullName := Name + " Kr Sharma"

	fmt.Fprint(w, "URL :"+FullName)

	var user User

	// Decoding from request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Body Parsing", user)

	u, _ := json.Marshal(user)

	w.Write(u)
}

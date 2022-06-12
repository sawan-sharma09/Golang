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
	http.HandleFunc("/", Decode)
	fmt.Println("Connected")
	http.ListenAndServe(":8090", nil)
}

func Decode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Decode using request URL")

	Name := r.URL.Query().Get("myname")
	fmt.Println("First name is:", Name)
	fullname := Name + " kr sharma"
	f, _ := json.Marshal(fullname)
	w.Write(f)

	///decode using request body

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Body Parsing", user)

	u, _ := json.Marshal(user)
	w.Write(u)

}

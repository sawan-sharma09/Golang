package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type myUser struct {
	Name string `json:"name"` // the struct variable should always start with capital letter here and the Json data should always be in small letters.
}

func main() {
	http.HandleFunc("/", htmlVsPlain)
	fmt.Println("Connected")
	http.ListenAndServe(":8000", nil)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	var User myUser
	fmt.Println("html vs plain")
	fmt.Fprint(w, "Welcome to Json")
	a := "SNEH"
	User.Name = "SAURAV"
	c, _ := json.Marshal(User.Name)
	b, _ := json.Marshal(a)
	w.Write(b)
	w.Write(c)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/mypath", getProfile)
	fmt.Println("Connected:")
	http.ListenAndServe(":8000", nil)
}
func getProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world, This is my first API, Hurray !!!")
}

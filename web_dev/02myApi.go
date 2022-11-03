package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getProfile) // we use http.handleFunc function to attach a handler to DefaultServeMux.
	fmt.Println("Connected:")
	http.ListenAndServe(":8080", nil)
}

func getProfile(w http.ResponseWriter, r *http.Request) { // This is handler function....In general, anything that has
	// a method called ServeHTTP with a method signature as above is called as handler

	fmt.Fprint(w, "Hello World\n")
	fmt.Fprint(w, "This is my 1st API")
}

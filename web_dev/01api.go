package main

import (
	//"http"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/profilename", getProfile)
	fmt.Println("CONNECTED")
	http.ListenAndServe(":8080", nil)
}
func getProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Print("CHECKING")
	fmt.Fprint(w, "SAWAN KR SHARMA")
}

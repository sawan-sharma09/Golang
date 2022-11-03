package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	http.HandleFunc("/", theHandler)
	fmt.Println("Connected")
	http.ListenAndServe(":8000", nil)
}

func theHandler(w http.ResponseWriter, r *http.Request) {
	courses := []course{
		{"Golang", 778, "dev.org", "abc123", []string{"docker", "Kubernetes"}},
		{"Java", 669, "Java_T_point", "xyz14", []string{"Automation", "Selenium"}},
		{"Python", 889, "python-gorilla", "hya123", nil},
	}
	finalJson, _ := json.MarshalIndent(courses, "", "   ")
	w.Write(finalJson)
}

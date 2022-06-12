package main

import ( //                                Normal Json without using a client-server
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // we have written "-" because it hides or doesn't print the password to protect privacy
	Tags     []string `json:"tags,omitempty"` // "omitempty", this is used to omit any nil or empty values. There should not be any space befote omitempty.
}

func main() {
	encodeJson()
}

func encodeJson() {
	theCourses := []course{
		{"Golang", 999, "GFG", "abc123", []string{"go_dev", "blockchain"}},
		{"Java", 299, "we3", "xyz123", []string{"java_dev", "automation"}},
		{"python", 456, "this", "agh123", nil},
	}

	// Package this data as JSON data

	finalJson, err := json.MarshalIndent(theCourses, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", finalJson)
}

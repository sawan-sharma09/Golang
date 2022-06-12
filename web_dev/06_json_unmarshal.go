package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	decodeJson()
}

func decodeJson() { //                copied this Json data from previous json_marshal program so that we can decode it
	jsonDataFromWeb := []byte(` 
	{
		"coursename": "Java",
		"Price": 669,
		"website": "Java_T_point",
		"tags": [
		   "Automation",
		   "Selenium"
		]
	 }
	`)

	var golangCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &golangCourse)
		fmt.Printf("%#v\n", golangCourse) // we will be printing a struct, that has been unmarshal using json
	} else {
		fmt.Println("Json was not valid")
	}

	//-------------------------------------------------------------------------------------------------------

	//if we want to iterate over the Unmarshal json data then we simply can't range over the struct
	// we need to store that json data over a map[string]interface{}

	var newPython map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &newPython)
	fmt.Printf("%#v", newPython)

	for key, val := range newPython {
		fmt.Printf("key is %v and val is %v and type is %T\n", key, val, val)
	}
}

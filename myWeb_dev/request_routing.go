package main

import ( //Reques routing and request method
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", letsHandle)
	fmt.Println("Connected:")
	http.ListenAndServe(":8000", nil)
}

func letsHandle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "This is the most basic URL Path")
	case "/name":
		fmt.Fprint(w, "Facebook was designed for communicating with friends")
	case "/ninja":
		fmt.Fprintln(w, "Lets have a fun and add a new line under it")
		fmt.Fprint(w, "this was an easy one, added the line successfully")
	default:
		fmt.Fprint(w, "Not a valid URL path, please exit the window")
	}

	fmt.Printf("Handling function with %s request\n", r.Method) // the output will be printed only when we open any
	//one of the above URLs in the browser. Eachtime we open the URL, the output will be printed in the below terminal.
}

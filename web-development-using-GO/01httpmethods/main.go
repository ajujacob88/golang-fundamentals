package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "This is a GET request")
		case "POST":
			fmt.Fprintf(w, "This is a POST request")
		case "PUT":
			fmt.Fprintf(w, "This is a PUT request")
		case "DELETE":
			fmt.Fprintf(w, "This is a DELETE request")
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":6080", nil)
}

/*
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe("", nil)

}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "this is get request")

	case "POST":
		fmt.Fprintf(w, "this is post request")
	default:
		http.Error(w, "invalid", http.StatusMethodNotAllowed)

	}
}
*/

//https://www.youtube.com/watch?v=BUSn3GgR6DU&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=5

package main

import (
	"net/http"
)

// this is not full
// this web requests thing already don in hitesh video
func handlefunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url": //when the path is url
		url(w, r)
		//http://localhost/url?name=Aju
	case "/body":
		body(w, r)
		//http://localhost/body
		//with body in json: {"name": "Aju"}
	default:
		w.Write([]byte("Hello World"))
		//http://localhost
	}

}

func main() {
	http.HandleFunc("/", handlefunc)
	http.ListenAndServe("", nil)
}

//https://www.youtube.com/watch?v=DiZuxLbpkGg&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=4
//last part of the above video

package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world mux")

	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello WORLD Session</h1>")

}

func main() {

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxAju http.ServeMux
	server.Handler = &muxAju
	muxAju.HandleFunc("/", helloWorldPage)
	server.ListenAndServe()
}

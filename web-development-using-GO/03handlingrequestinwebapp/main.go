//https://www.youtube.com/watch?v=DiZuxLbpkGg&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=4

package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World Aju")
	case "/aju":
		fmt.Fprint(w, "My name is Aju Jacob")
	default:
		fmt.Fprint(w, "Invalid URL entered")
	}

	fmt.Printf("Handling function with %s request\n", r.Method)

}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("html vs Plain handler")

	//fmt.Fprint(w, "<h1>Hello WORLD type 1 html</h1>")

	// w.Header().Set("Content-Type", "text/plain")
	// fmt.Fprint(w, "<h1>Hello WORLD type 2 plain</h1>")

	w.Header().Set("Content-Type", "text/html") //without this line also, the browser will treat the below line tags as html since it contains html tags
	fmt.Fprint(w, "<h1>Hello WORLD type 3 html</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("timeout attemp")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "no timeout,") //this ine will not run because sleep is 2s and timeout set was 1 sec, before this line time out will happens
}

/*
func main() {
	//http.HandleFunc("/", helloWorldPage) //starting our server,,,"/" means passing the root passing the root path for the page

	http.HandleFunc("/", htmlVsPlain)
	http.ListenAndServe("", nil) //passing empty port no meas passing default port no which is 80,,starting the server by calling listen and serve and passing the port no, not passing port no means default port no 80 will be pased, http.ListenAndServe(addr:"", handler: nil)
}
*/

func main() {
	//http.HandleFunc("/", helloWorldPage) //starting our server,,,"/" means passing the root passing the root path for the page

	// http.HandleFunc("/", htmlVsPlain)
	//http.ListenAndServe("", nil) //passing empty port no meas passing default port no which is 80,,starting the server by calling listen and serve and passing the port no, not passing port no means default port no 80 will be pased, http.ListenAndServe(addr:"", handler: nil)

	http.HandleFunc("/timeout", timeout)
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()
}

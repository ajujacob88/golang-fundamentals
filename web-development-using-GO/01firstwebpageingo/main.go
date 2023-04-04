//https://www.youtube.com/watch?v=BOxBcuvASag&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=2

//https://www.youtube.com/watch?v=1-Rx9sWQYL0&list=PLDtWoQ-cxqiz05TeCdeqwACNzk_rYreDD&index=2

package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to first web page in GOlang by url %s", r.URL.Path)
	fmt.Fprint(w, "Hello World Aju") //writing to http response writer using this
}

func main() {
	http.HandleFunc("/", helloWorldPage) //starting our server,,,"/" means passing the root passing the root path for the page
	//http.ListenAndServe(":8000", nil)
	http.ListenAndServe("", nil) //passing empty port no meas passing default port no which is 80,,starting the server by calling listen and serve and passing the port no, not passing port no means default port no 80 will be pased, http.ListenAndServe(addr:"", handler: nil)
}

/*

package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to first web page in GOlang by url %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handlerFunc) //"/" means passing the root
	http.ListenAndServe(":8000", nil)
}
*/

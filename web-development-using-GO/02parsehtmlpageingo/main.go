//https://www.youtube.com/watch?v=JbFh2iLEdj0&list=PLDtWoQ-cxqiz05TeCdeqwACNzk_rYreDD&index=3

package main

import (
	"net/http"
	"text/template"
)

var templ *template.Template

func init() { //The init() function is reserved and is used for specific reasons. This function is defined to take no arguments and return nothing. This function is meant to run before any other piece of code – even before the main() function, which supposedly is the first function from where execution of our program begins. The init() function serves some specific purposes such as to initialize a specific state of the application.
	templ = template.Must(template.ParseGlob("template/*.html"))

}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "welcome to <h1> Go Web Page </h1> - Aju JAcob")
	templ.ExecuteTemplate(w, "welcomepage.html", nil)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "welcome to <h1> Go Web Page </h1> - Aju JAcob")
	templ.ExecuteTemplate(w, "about.html", nil)
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/about", handleAbout)
	http.ListenAndServe(":9999", nil)

}

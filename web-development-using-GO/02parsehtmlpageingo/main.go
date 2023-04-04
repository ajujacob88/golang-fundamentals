//https://www.youtube.com/watch?v=JbFh2iLEdj0&list=PLDtWoQ-cxqiz05TeCdeqwACNzk_rYreDD&index=3

package main

import (
	"net/http"
	"text/template"
)

var templ *template.Template

func init() {
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

// https://www.youtube.com/watch?v=YBBX1k9lDKw&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=6
// i skipped the htttps and substitution part(important one),embedded functions in the video, which is also useful lateron..
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"             //this is the file we are going to display on the browser when a request is received from browser
	t, err := template.ParseFiles(fileName) //also manually change the import to "html/template" instead of "text/template"
	if err != nil {

		fmt.Println("Error when parsing file - aju")
		return
		//panic(err)
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {

		fmt.Println("Error when executing template")
		return
		//panic(err)
	}

}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	//fmt.Println(username, password)

	if userDataBase[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You are now logged in. Welcome to Home ")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Didn't find matching credentials ")
	}
}

var userDataBase = map[string]string{
	"aju": "mypassword",
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": //todo: input boxes and button for credentials
		login(w, r)
	case "/login-submit": //todo: handle login credentials
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "not pres")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}

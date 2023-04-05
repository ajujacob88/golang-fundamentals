//https://www.youtube.com/watch?v=uey1OHKKDXk&list=PLDtWoQ-cxqiz05TeCdeqwACNzk_rYreDD&index=21&t=1272s

// https://www.youtube.com/watch?v=YBBX1k9lDKw&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=6
// i skipped the htttps and substitution part(important one),embedded functions in the video, which is also useful lateron..

//for sessions part here some code writtten in html and here, to understand that watch the abouve dojo link and watch the substitution part in the video

package main

import (
	"html/template"
	"net/http"

	"github.com/icza/session"
)

var temp *template.Template

func init() { //The init() function is reserved and is used for specific reasons. This function is defined to take no arguments and return nothing. This function is meant to run before any other piece of code – even before the main() function, which supposedly is the first function from where execution of our program begins. The init() function serves some specific purposes such as to initialize a specific state of the application.

	temp = template.Must(template.ParseGlob("template/*.html"))

}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

var userDataBase = map[string]string{
	"aju":       "mypassword",
	"brototype": "kochi",
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	//need to create a session first, session is created only when the user is valid
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDataBase[username] == password && password != "" { //or we can simply write if username == "Aju" && password =="passw",, password != "" is written because, if required is written in html(line 33 and line 36 in html) then no need of this, but if required is not written, in html then if a username is only entered, then also it will enter into the login page..
		//if matching, create a session using icza package, also import the package for doing this
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{
				"username": username,
			},
		})
		session.Add(sess, w)

		//if login then redirect to welcome directory
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		//if username and password is not matching
		data := map[string]interface{}{
			"error": "Invalid Username and Password Entered",
		}
		temp.ExecuteTemplate(w, "index.html", data) //here instead of nil data is given and done some coding in htl file also, this is sustitution part in golang dojo video, watch that also..
	}

}

func welcomeHandle(w http.ResponseWriter, r *http.Request) {

	//To prevent users from accessing the previous page after logging out, you can use HTTP headers to set cache-control and no-store directives. This will ensure that the browser does not cache the pages that require authentication and prevents users from accessing them after logging out.

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") //This tells the client that the response should not be cached by any intermediary (such as a proxy server) or by the client itself. It also tells the client that it must revalidate the response with the server before using a cached copy, even if it has a fresh copy in its cache.
	w.Header().Set("Pragma", "no-cache")                                   //This is an older header that was used to achieve similar behavior to "Cache-Control: no-cache" in older versions of HTTP. It is included here for compatibility with older clients that may not recognize the "Cache-Control" header.
	w.Header().Set("Expires", "0")                                         //line sets the "Expires" header to "0", which means that the response has already expired and should not be used. This is included for compatibility with older clients that do not support the newer caching headers.

	sess := session.Get(r)
	username := sess.CAttr("username")
	data := map[string]interface{}{
		"username": username,
	}
	temp.ExecuteTemplate(w, "welcome.html", data)

}

func logoutHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}

	//if logout then redirect to root directory
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/welcome", welcomeHandle)
	http.HandleFunc("/logout", logoutHandle)
	//now start the server
	http.ListenAndServe("", nil)

}

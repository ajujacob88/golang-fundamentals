package main

import (
	"html/template"
	"net/http"

	"github.com/icza/session"
)

/*
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
/*
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb" //this is the basics of url format, 3000 or something likethat is the port no when hosted in cloud, & is used instead of comma in url

func main() {
	fmt.Println("Handling url")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
}
*/

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/welcome", welcomeHandle)
	http.HandleFunc("/logout", logoutHandle)

	http.ListenAndServe("", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

var userDB = map[string]string{
	"aju": "mypassword",
	"anu": "pass",
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{
				"username": username,
			},
		})
		session.Add(sess, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"error": "Invalid Username and Password Entered",
		}
		temp.ExecuteTemplate(w, "indexhtml", data)
	}
}

func welcomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/*
var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/login", loginHandle)
	// http.HandleFunc("/welcome", welcomeHandle)
	// http.HandleFunc("/logout", logoutHandle)

	http.ListenAndServe("", nil)
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

var userDB = map[string]string{
	"aju": "mypassword",
	"anu": "pass",
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		setSession(username, w)
		http.Redirect(w, r, "/welcome", http.StatusFound)
	} else {
		data := map[string]interface{}{
			"error": "Invalid Username and Password Entered",
		}
		temp.ExecuteTemplate(w, "indexhtml", data)
	}
}

func setSession(username string, w http.ResponseWriter) {
	value := map[string]string{
		"name": username,
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, &cookie)
	}
}

func welcomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")


	temp.ExecuteTemplate(w, "welcome.html", data)

}
*/

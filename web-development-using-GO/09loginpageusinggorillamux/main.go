//by stebin using mux

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("static/*.html"))
}

var verify = map[string]string{
	"stebinsabu369@gmail.com": "Stebin@333",
	"hihello333@gmail.com":    "Hihelo@333",
}

type user struct {
	Name  string
	Place string
}

func indexpage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	err := t.ExecuteTemplate(w, "login.html", "Please login")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
}
func loginhandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("email")
	password := r.FormValue("password")
	if verify[username] == password {
		setSession(username, w)
		http.Redirect(w, r, "/homepage", http.StatusFound)
	} else {
		err := t.ExecuteTemplate(w, "login.html", "Invalid Credintials")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
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
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	var s user
	s.Name = getUsername(r)
	if s.Name == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	s.Place = "India"
	err := t.ExecuteTemplate(w, "home.html", s)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
func getUsername(r *http.Request) (userName string) {
	cookie, err := r.Cookie("session")
	if err == nil {
		Value := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &Value)
		if err == nil {
			userName = Value["name"]
		}
	}
	return
}
func logouthandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", indexpage)
	router.HandleFunc("/login-submit", loginhandler).Methods("POST")
	router.HandleFunc("/homepage", homeHandler)
	router.HandleFunc("/logout", logouthandler).Methods("POST")
	// url, err := router.Host(localhost).Schemes(http).Methods("PATCH")
	// http.Handle("/", router)
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:2670",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe() //http://localhost:2670/
}

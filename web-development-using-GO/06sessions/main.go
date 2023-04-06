package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("very-secret")))

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	// if session.Values["authentication"] != nil {
	if !session.Values["authentication"].(bool) {
		http.Error(w, "Unauthorized accesss", http.StatusForbidden)
		return
	}
	fmt.Fprint(w, "Authorized user home page")
	// }
}
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authentication"] = true
	session.Save(r, w)
	fmt.Fprint(w, "Successfully logged in")
}
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authentication"] = false
	session.Save(r, w)
	fmt.Fprint(w, "Successfully logged out")
}
func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":1570", nil)
}

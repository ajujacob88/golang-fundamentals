package main

import (
	"fmt"
	"os"
	"strings"
)

//standard convention for go

const (
	usage = `
Usage: [username] [password]

 # username: Aju
 # password: man
eg: go run main.go Aju man`

	errUser  = "Access denied for %q\n"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q\n"

	user, user2 = "Aju", "Nikhila"
	pass, pass2 = "man", "2004"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}

}

//for single user
/*
const (
	usage = `
Usage: [username] [password]

 # username: Aju
 # password: man
eg: go run main.go Aju man`

	errUser  = "Access denied for %q\n"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q\n"

	user = "Aju"
	pass = "man"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u != user  {
		fmt.Printf(errUser, user)

	} else if p != pass  {
		fmt.Printf(errPwd, user)
	} else {
		fmt.Printf(accessOK, user)
	}

}

*/
//no need
/*
const usage = `
Usage: [username] [password]

 # username: Aju
 # password: man
eg: go run main.go Aju man`

func main() {

	if len(os.Args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	user, pw := os.Args[1], os.Args[2]

	if user == "Aju" && pw == "man" {
		fmt.Printf("Access granted to %q\n", user)

	} else if user == "Aju" && pw != "man" {
		fmt.Printf("Invalid password for %q\n", user)
	} else {
		fmt.Printf("Access denied for %q\n", user)
	}

}
*/

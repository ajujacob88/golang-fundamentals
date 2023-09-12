package main

import (
	"fmt"
	"strconv"
)

func main() {
	//short if statement

	if n, err := strconv.Atoi("25a"); err != nil { //but while using this short if, n and err variable can only be used within this if statement and its branches,, this variable wont be available outside of the if statement
		fmt.Println("there is an error:", err)
	} else {
		fmt.Println(n)
	}
	//fmt.Println(n)

}

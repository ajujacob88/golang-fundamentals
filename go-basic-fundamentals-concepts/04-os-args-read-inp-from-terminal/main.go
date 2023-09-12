package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("please enter a value in the console eg: go run main.go 100 200 300")
		return
	}
	fmt.Println(os.Args)
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("path:", os.Args[0])
	fmt.Println("1st argumnent:", os.Args[1])
	fmt.Println("2nd argument", os.Args[2])
	fmt.Println("no of args inside os.Args:", len(os.Args))

	//var name string
	name := os.Args[1]
	fmt.Println("my name is ", name)
}

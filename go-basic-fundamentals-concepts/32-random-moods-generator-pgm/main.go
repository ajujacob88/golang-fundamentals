package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter the name too, eg: go run main.go aju")
		return
	}
	name := args[0]

	mood := [...]string{"feeling good", "feeling sad", "feeling angry", "feeling happy"}
	n := rand.Intn(len(mood))

	fmt.Printf("%s feels %s\n", name, mood[n])
}

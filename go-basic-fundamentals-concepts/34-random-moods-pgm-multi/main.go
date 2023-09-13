package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Enter the name and mood [positive / negative], eg: go run main.go aju positive")
		return
	}
	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy", "good", "awesome"},
		{"sad", "bad", "terrible"},
	}

	n := rand.Intn(len(moods))

	var p int
	if mood == "negative" {
		p = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[p][n])
}

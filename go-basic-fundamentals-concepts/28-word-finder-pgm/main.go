package main

import (
	"fmt"
	"os"
	"strings"
)

const str = "Hi, good morning Today is Monday and good morning"

func main() {
	words := strings.Fields(str)

	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d:  %q\n", i+1, q)
				break
			}

		}
	}

}

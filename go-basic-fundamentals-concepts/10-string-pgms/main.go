//get input form the terminal and return like the following example
// input - go run main.go aju
//output - AJU!!!

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inp := os.Args[1]

	count := len(inp)
	rep := strings.Repeat("!", count)

	fmt.Println(strings.ToUpper(inp) + rep)

}

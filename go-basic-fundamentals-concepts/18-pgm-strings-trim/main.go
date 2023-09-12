//pgm feet to meter

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet To Metres
--------------
This program converts feet into metres.

Usage:
feet [feets to convert]
eg:
go run main.go 100`

func main() {

	if len(os.Args) < 2 {
		//fmt.Println(usage)
		fmt.Println(strings.TrimSpace(usage)) // trimspace is to remove unnecessary spaces
		return
	}
	args := os.Args[1]
	feet, _ := strconv.ParseFloat(args, 64)
	metre := feet * 0.3048
	fmt.Printf("%g feet is %g metres\n", feet, metre)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("please enter a value in the console eg: go run main.go 100")
		return
	}
	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	metre := feet * 0.3048

	fmt.Printf("%f feet is %f metres\n", feet, metre)
	fmt.Printf("%g feet is %g metres\n", feet, metre) //%g will remove unnecessary decimal parts

	string := "520"
	convertedInt, _ := strconv.Atoi(string)
	fmt.Println(convertedInt)

	integer := 25
	convertedstr := strconv.Itoa(integer)
	fmt.Println(convertedstr)

	eq := strconv.FormatBool(true)
	fmt.Println(eq)

}

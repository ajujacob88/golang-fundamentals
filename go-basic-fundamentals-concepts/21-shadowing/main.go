package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	if n, err := strconv.Atoi("25"); err != nil {
		fmt.Println("there is an error:", err)
	} else {
		fmt.Println(n)
	}
	fmt.Println("The value of n is:", n, "  OOPS! You've been shadowed ") //yes here the n is declared first, but that n is not used in the if statement, there another n is used, so this is called shadowing,,so here there are 2 n variables,, this may cause errors for programmers, so be careful

	var (
		num int
		err error
	)
	if num, err = strconv.Atoi("125"); err != nil {
		fmt.Println("there is an error:", err)
	} else {
		fmt.Println(num)
	}

	fmt.Println("Here, The value of n is:", num, "NOT SHADOWED") //now here is only 1 num variable

	//by doing some settings in vs code, a warning can be generated in vs code if shadowing occurs

}

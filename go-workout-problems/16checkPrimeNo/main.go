package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	flag := 0
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag++
			break
		}
	}
	if flag == 1 || num == 1 {
		fmt.Println("The entered number ", num, " is Not a Prime number")
	} else {
		fmt.Println("The entered number ", num, " is a Prime number")
	}
}

package main

import "fmt"

func main() {
	var limit int
	fmt.Println("Enter the limit")
	fmt.Scan(&limit)
	var sum int = 0
	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("The sum of odd numbers upto ", limit, " is ", sum)
}

package main

import "fmt"

func main() {
	//ARITHMETIC OPERATIONS

	ratio := 1.0 / 10
	fmt.Printf("%.60f\n", ratio)

	ratio = 3 / 2
	fmt.Println(ratio)

	ratio = 3.0 / 2
	fmt.Println(ratio)

	fmt.Println(2 + 2*5/2)

	fmt.Println(1 + 5 - 3*10/2)

	num := 10
	num += 2
	fmt.Println(num)

	num *= 5
	fmt.Println(num)

	num /= 2
	fmt.Println(num)

	num %= 4
	fmt.Println(num)
}

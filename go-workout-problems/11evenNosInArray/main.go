package main

import "fmt"

func main() {
	var size int

	fmt.Println("Enter the size of the Array: ")
	fmt.Scan(&size)
	fmt.Println("Enter the values of the array:")
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	flag := 0
	for i := 0; i < size; i++ {
		if array[i]%2 == 0 {
			flag++
		}
	}

	/* Another method
	for _, val := range array {
		if val%2 == 0 {
			flag++
		}
	}

	*/

	fmt.Println("The no of even numbers in the array is: ", flag)
}

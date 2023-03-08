package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)
	fmt.Println("Enter the values of the array")
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("The entered array is: ", array)

	//sort.Ints(array)  //sorting in ascending order
	// fmt.Println(array)

	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println("The sorted array in descending Order is: ", array)

}

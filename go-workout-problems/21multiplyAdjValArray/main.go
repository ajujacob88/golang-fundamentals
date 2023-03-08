package main

import "fmt"

func main() {
	var size, value int
	var sliceArray []int
	var resultSlice []int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)
	fmt.Println("Enter the values of the array")
	for i := 0; i < size; i++ {
		fmt.Scan(&value)
		sliceArray = append(sliceArray, value)
	}

	for i := 0; i < size-1; i++ {
		resultSlice = append(resultSlice, sliceArray[i]*sliceArray[i+1])
	}

	fmt.Println("The entered array is: ", sliceArray, "\nThe output array after multiply adjacent values is: ", resultSlice)

}

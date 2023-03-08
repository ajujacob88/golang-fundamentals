package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)
	fmt.Println("Enter the values of array 1")
	slice1 := getArray(size)
	fmt.Println("Enter the values of array 1")
	slice2 := getArray(size)
	resultSlice := sumArray(slice1, slice2, size)
	fmt.Println("The input Array 1 enterd is: ")
	displayArray(slice1)
	fmt.Println("The input Array 2 entered is: ")
	displayArray(slice2)
	fmt.Println("The sum of Array1 and 2 is: ")
	displayArray(resultSlice)

}

func getArray(length int) [][]int {
	var value int
	var tempSlice []int
	var slice [][]int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		slice = append(slice, tempSlice)
		tempSlice = nil
	}
	return slice
}

func sumArray(inpSlice1 [][]int, inpSlice2 [][]int, length int) [][]int {
	var tempSlice []int
	var addedSlice [][]int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			tempSlice = append(tempSlice, inpSlice1[i][j]+inpSlice2[i][j])
		}
		addedSlice = append(addedSlice, tempSlice)
		tempSlice = nil
	}
	return addedSlice
}

func displayArray(input [][]int) {
	fmt.Println(input)
}

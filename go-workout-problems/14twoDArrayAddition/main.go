package main

import "fmt"

func main() {
	var size, value int
	fmt.Println("Enter the size of Arrays")
	fmt.Scan(&size)
	var slice1 [][]int
	var slice2 [][]int
	var sumSlice [][]int
	var tempSlice []int

	fmt.Println("Enter the values of first array")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		slice1 = append(slice1, tempSlice)
		tempSlice = nil
	}

	fmt.Println("Enter the values of second array")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		slice2 = append(slice2, tempSlice)
		tempSlice = nil
	}

	fmt.Println("Array 1 is:\n", slice1)
	fmt.Println("Array 2 is:\n", slice2)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			value = slice1[i][j] + slice2[i][j]
			tempSlice = append(tempSlice, value)
		}
		sumSlice = append(sumSlice, tempSlice)
		tempSlice = nil
	}
	fmt.Println("The sum of arrays are:\n", sumSlice)
}

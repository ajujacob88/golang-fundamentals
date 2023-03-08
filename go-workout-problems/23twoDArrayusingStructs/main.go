package main

import (
	"fmt"
)

func main() {
	mySlice := myStructArray{}
	// mySlice2 := myStructArray{}
	// matrixvalues := mySlice2.getArray()
	// mySlice2.displayArray(matrixvalues)
	mySlice.displayArray(mySlice.getArray())
}

type myStructArray struct {
	userSlice [][]int
	size      int
}

func (m *myStructArray) getArray() [][]int {
	var value int
	var tempSlice []int

	fmt.Println("Enter the size of the array")
	fmt.Scan(&m.size)
	fmt.Println("Enter the values of array")
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		m.userSlice = append(m.userSlice, tempSlice)
		tempSlice = nil
	}
	return m.userSlice
}

func (m myStructArray) displayArray(matrix [][]int) {
	fmt.Println("The array is: ")
	fmt.Println(matrix)
}

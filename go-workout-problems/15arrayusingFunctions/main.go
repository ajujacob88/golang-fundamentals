package main

import "fmt"

func main() {

	slice1 := getArray()
	slice2 := getArray()
	displayArray(slice1)
	displayArray(slice2)
}

func getArray() []int {
	var size, value int
	var sliceInp []int
	fmt.Println("Enter the size of the Array:")
	fmt.Scan(&size)
	fmt.Println("Enter the values of the array")
	for i := 0; i < size; i++ {
		fmt.Scan(&value)
		sliceInp = append(sliceInp, value)
	}
	return sliceInp
}

/* another method
func getArray() []int {
	var size int
	fmt.Println("Enter the size of the Array:")
	fmt.Scan(&size)
	sliceInp := make([]int, size)
	fmt.Println("Enter the values of the array")
	for i := 0; i < size; i++ {

		fmt.Scan(&sliceInp[i])
	}
	return sliceInp
}
*/
func displayArray(sliceOut []int) {
	fmt.Println("The entered array is: \n", sliceOut)
}

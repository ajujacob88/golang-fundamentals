package main

import "fmt"

func main() {
	fmt.Println("Enter the size of the arrays")
	var size int
	fmt.Scan(&size)
	fmt.Println("Enter the values of array 1")
	var array1 = make([]int, size)
	arrayIn(array1, size)
	fmt.Println("Enter the values of array 2")
	var array2 = make([]int, size)
	arrayIn(array2, size)
	fmt.Println("Array 1 entered is: ", array1)
	fmt.Println("Array 2 entered is: ", array2)

	temp := array1
	array1 = array2
	array2 = temp

	fmt.Println("The Arrays after swapping:")
	fmt.Println("Array 1: ", array1)
	fmt.Println("Array 2: ", array2)

}

func arrayIn(array []int, len int) {
	for i := 0; i < len; i++ {
		fmt.Scan(&array[i])
	}
}

//ANOTHER METHOD BY GANESH
/*
package main

import (
    "fmt"
)

func main() {
    var slice1, slice2 []int
    var limit, value, temp int
    fmt.Println("Enter a limit ")
    fmt.Scan(&limit)
    fmt.Println("Enter the first array elements")
    for i := 0; i < limit; i++ {
        fmt.Scan(&value)
        slice1 = append(slice1, value)
    }
    fmt.Println("Enter the second array elements")
    for i := 0; i < limit; i++ {
        fmt.Scan(&value)
        slice2 = append(slice2, value)
    }
    fmt.Println("This is the swapped arrays")
    for i := 0; i < limit; i++ {
        temp = slice1[i]
        slice1[i] = slice2[i]
        slice2[i] = temp
    }
    fmt.Printf("Array 1: %v", slice1)
    fmt.Println("")
    fmt.Printf("Array 2: %v", slice2)

}
*/

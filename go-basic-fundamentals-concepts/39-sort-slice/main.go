package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{5, 4, 6, 2, 8, 1, 2, 4, 0, 54, 21, 4, 6}

	fmt.Println(slice)
	sort.Ints(slice)

	fmt.Println("sorted ascendig is:", slice)

	slice2 := []string{"aju", "ra", "c", "t", "g", "s"}
	fmt.Println(slice2)
	sort.Strings(slice2)

	fmt.Println("sorted ascendig is:", slice2)

	//but array can be sorted directly, because its having fixed sizes
	//we can use an array like a slice by putting [:] can then sort

	array := [6]int{5, 1, 9, 4, 1, 0}
	sort.Ints(array[:]) //[:] means array can be also used like a slice

	fmt.Println(array)
	fmt.Printf("%#v\n", array)

}

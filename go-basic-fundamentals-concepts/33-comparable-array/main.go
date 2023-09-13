package main

import "fmt"

func main() {
	//An array can be comparable, if both the type and size of the 2 arrays are same,,, also, maps and slices are not comparable

	red := [3]int{4, 50, 62}
	blue := [3]int{4, 50, 62}

	fmt.Println(red == blue)

	green := [3]int{41, 50, 62}
	fmt.Println(red == green)

	//black := [4]int{41, 50, 62}
	//fmt.Println(red == black)

	brown := red
	fmt.Println(brown == red)

	//multi dimensional array

	student := [2][3]int{
		[3]int{5, 6, 7},
		[3]int{8, 9, 4},
	}

	fmt.Println(student)

	//we can write simply like this
	student1 := [2][3]int{
		{5, 6, 7},
		{8, 9, 4},
	}

	fmt.Println(student1)

	student2 := [...][3]int{
		{5, 6, 7},
		{8, 9, 4},
	}

	fmt.Println(student2)

	// ACCESSING MULTI DI ARRAY USINF FO LOOP
	var sum int

	for _, v := range student1 {
		for _, v2 := range v {
			sum += v2
		}
	}

	fmt.Println(sum)

}

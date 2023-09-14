package main

import "fmt"

func main() {
	var arr [20]int
	var slic []int

	fmt.Printf("%#v\n", arr)
	fmt.Printf("%#v\n", slic)

	//slic[1] = 25
	//fmt.Printf("%#v\n", slic)

	slic2 := []int{2, 5, 6}

	fmt.Printf("%#v\n", slic2)

	slic3 := slic2
	fmt.Printf("%#v\n", slic3)

	for _, v := range slic3 {
		fmt.Println(v)
	}

	//slic[1] = 25
	//fmt.Printf("%#v\n", slic)

	slic = append(slic, 25, 52, 355)
	fmt.Printf("%#v\n", slic)
}

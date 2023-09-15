package main

import "fmt"

func main() {
	nums := []int{4, 5, 6}

	nums = append(nums, 5, 4, 7) //remeber, append() returns a new slice value.,save it over your original slice or another slice, otherwise it will be lost
	fmt.Println("nums:", nums)

	nums2 := append(nums, 11, 12, 13)
	fmt.Println("nums:", nums)
	fmt.Println("nums:", nums2)

	temp := []int{100, 101, 102}
	nums2 = append(nums2, temp...) //by using ellipsis ...

	fmt.Println("nums2-", nums2)

	//an append trick to append value to the middle of a slice
	val := []int{1, 2, 3, 7, 8, 9}
	fmt.Println(val)
	val = append(val, val[3:]...)
	fmt.Println(val)
	val = append(val[:3], 4, 5, 6)
	fmt.Println(val)
	fmt.Println(val[:])
	fmt.Println(val[:9])

	//anothe one
	digi := []int{11, 12, 13, 17, 18, 19, 20, 21, 22}
	fmt.Println(digi)
	digi = append(digi, digi[3:]...)
	fmt.Println(digi)
	digi = append(digi[:3], 14, 15, 16)
	fmt.Println(digi)
	fmt.Println(digi[:15])
	digi = append(digi[:6], digi[9:15]...)
	fmt.Println(digi)

}

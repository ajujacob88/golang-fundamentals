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
}

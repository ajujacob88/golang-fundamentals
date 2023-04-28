//https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

/*
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

*/

package main

import "fmt"

func main() {

	var slice1 = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(slice1))
}

func removeDuplicates(nums []int) int {
	prev := nums[0]
	x := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[x] = nums[i]
			x++
		}
		prev = nums[i]
	}
	fmt.Println(nums)
	return x
}

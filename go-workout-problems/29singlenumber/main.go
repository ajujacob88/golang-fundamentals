// https://leetcode.com/problems/single-number/description/
/*
136. Single Number
Easy
13.6K
522
Companies
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1
*/

package main

import "fmt"

func main() {

	var slice1 = []int{25, 32, 53, 25, 53}
	fmt.Println(singleNumber(slice1))
}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}

	return result
}

/*
485. Max Consecutive Ones
Easy
4.3K
425
Companies
Given a binary array nums, return the maximum number of consecutive 1's in the array.



Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2

//https://leetcode.com/problems/max-consecutive-ones/description/

*/

package main

import "fmt"

func main() {

	var nums = []int{1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	flag := 0
	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			flag++
			if flag > count {
				count = flag
			}
		} else {
			flag = 0
		}

	}
	return count
}

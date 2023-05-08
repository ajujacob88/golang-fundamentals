/*
448. Find All Numbers Disappeared in an Array
Easy
8.4K
431
Companies
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.



Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]

https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

*/

package main

import "fmt"

func main() {

	var nums = []int{1, 1, 0, 1, 1}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	flag := 0
	var number []int
	var j int
	for i := 1; i <= len(nums); i++ {
		flag = 0
		for j = 0; j < len(nums); j++ {
			if i == nums[j] {

				flag = 1
				//fmt.Println("j is", j)
			}

		}
		if flag == 0 {
			number = append(number, i)
		}
	}
	return number
}

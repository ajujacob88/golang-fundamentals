/*
1394. Find Lucky Integer in an Array
Easy
964
26
Companies
Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

Return the largest lucky integer in the array. If there is no lucky integer return -1.



Example 1:

Input: arr = [2,2,3,4]
Output: 2
Explanation: The only lucky number in the array is 2 because frequency[2] == 2.
Example 2:

Input: arr = [1,2,2,3,3,3]
Output: 3
Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.
Example 3:

Input: arr = [2,2,2,3,3]
Output: -1
Explanation: There are no lucky numbers in the array.

https://leetcode.com/problems/find-lucky-integer-in-an-array/description/

*/

package main

import "fmt"

func main() {

	var nums1 = []int{1, 2, 2, 3, 3, 3}

	fmt.Println(findLucky(nums1))
}

func findLucky(arr []int) int {
	flag := 0
	temp := 0
	lucky := -1
	for i := 0; i < len(arr); i++ {
		flag = 0
		temp = 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				temp = arr[i]
				flag++
			}
		}
		if temp == flag && flag > lucky {
			lucky = flag
		}
	}
	return lucky
}

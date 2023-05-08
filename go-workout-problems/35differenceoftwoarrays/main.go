/*
https://leetcode.com/problems/find-the-difference-of-two-arrays/description/

2215. Find the Difference of Two Arrays
Easy
1.6K
59
Companies
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.



Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
Example 2:

Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

*/

package main

import "fmt"

func main() {

	var nums1 = []int{1, 2, 3}
	var nums2 = []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {

	distinctnums1 := DistinctSlice(nums1)
	distinctnums2 := DistinctSlice(nums2)
	var output [][]int
	var temp []int
	tempnum := 0
	flag := 0

	for i := 0; i < len(distinctnums1); i++ {
		flag = 0
		for j := 0; j < len(distinctnums2); j++ {
			if distinctnums1[i] == distinctnums2[j] {
				flag = 1
				break
			}

			tempnum = distinctnums1[i]

		}
		if flag == 0 {
			temp = append(temp, tempnum)
		}

	}
	output = append(output, temp)
	var temp2 []int
	for i := 0; i < len(distinctnums2); i++ {
		flag = 0
		for j := 0; j < len(distinctnums1); j++ {
			if distinctnums2[i] == distinctnums1[j] {
				flag = 1
				break
			}

			tempnum = distinctnums2[i]

		}
		if flag == 0 {
			temp2 = append(temp2, tempnum)
		}

	}
	output = append(output, temp2)
	return output
}

func DistinctSlice(input []int) []int {
	distinct := make(map[int]bool)
	result := make([]int, 0)

	for _, item := range input {
		if _, found := distinct[item]; !found {
			distinct[item] = true
			result = append(result, item)
		}
	}

	return result
}

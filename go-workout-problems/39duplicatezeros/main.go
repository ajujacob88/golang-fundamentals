/*
1089. Duplicate Zeros
Easy
2.3K
662
Companies
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.



Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
Example 2:

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]

https://leetcode.com/problems/duplicate-zeros/description/

*/

package main

import "fmt"

func main() {

	var nums1 = []int{1, 0, 2, 3, 0, 4, 5, 0}

	duplicateZeros(nums1)
}

func duplicateZeros(arr []int) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
	fmt.Println(arr)
}

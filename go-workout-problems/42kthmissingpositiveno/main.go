/*
1539. Kth Missing Positive Number
Easy
5.3K
342
Companies
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Return the kth positive integer that is missing from this array.

Example 1:

Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
Example 2:

Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.

https://leetcode.com/problems/kth-missing-positive-number/description/
*/
package main

import "fmt"

func main() {
	//num1:= []int{1,2,3,4}
	var num = []int{2, 3, 4, 7, 11}
	fmt.Println(findKthPositive(num, 5))
}

func findKthPositive(arr []int, k int) int {
	var missingarr []int
	j := 1
	for i := 0; i < len(arr); i++ {
		for arr[i] != j {
			missingarr = append(missingarr, j)
			j++
		}
		j++
	}

	for len(missingarr) < k {

		missingarr = append(missingarr, j)
		j++

	}

	return missingarr[k-1]

}

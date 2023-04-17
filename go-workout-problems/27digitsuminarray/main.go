// leetcode - problem 2535. Difference Between Element Sum and Digit Sum of an Array

/*
You are given a positive integer array nums.

The element sum is the sum of all the elements in nums.
The digit sum is the sum of all the digits (not necessarily distinct) that appear in nums.
Return the absolute difference between the element sum and digit sum of nums.

Note that the absolute difference between two integers x and y is defined as |x - y|.

*/

// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/description/

package main

import (
	"fmt"
	"math"
)

func main() {

	var slice1 = []int{2, 15, 42, 3}
	fmt.Println("absoulte diff is :", differenceOfSum(slice1))

}

func differenceOfSum(nums []int) int {
	var elementsum int
	for i := 0; i < len(nums); i++ {
		elementsum = elementsum + nums[i]
	}
	// return elementsum
	var digitsum int
	var temp int
	for i := 0; i < len(nums); i++ {
		temp = nums[i]
		for temp > 0 {
			digit := temp % 10
			digitsum = digitsum + digit
			temp = temp / 10
		}
	}
	//return digitsum
	absolute := math.Abs(float64(elementsum) - float64(digitsum))
	return int(absolute)
}

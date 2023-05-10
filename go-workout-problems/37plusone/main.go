/*
66. Plus One
Easy
7.1K
4.9K
Companies
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].





https://leetcode.com/problems/plus-one/description/



*/

package main

import "fmt"

func main() {

	var nums1 = []int{1, 2, 2, 3, 3, 3}

	fmt.Println(plusOne(nums1))
}

/* This wont work for large number because int doesnt hold that much value

func plusOne(digits []int) []int {
    var num int
    power := 1
    //length := len(digits)
    for i:=len(digits)-1;i>=0;i--{
        num = num + digits[i]*power
        power = power * 10
        fmt.Println(num)
    }
    num = num + 1
    fmt.Println(num)
    var resultarr []int
    var resultarr1 []int

    for num > 0 {
			digit := num % 10
			resultarr = append(resultarr,digit)
			num = num / 10
           // resultarr = append(resultarr,num)
		}

    for i:=len(resultarr)-1;i>=0;i--{
        resultarr1 = append(resultarr1,resultarr[i])
    }
    //resultarr = append(resultarr,num)
     return resultarr1
}

*/

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		}

		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
	}
	return digits
}

// leetcode - problem 10 - Given an integer x, return true if x is a palindrome, and false otherwise.solve it without converting the integer to a string
//https://leetcode.com/problems/palindrome-number/
//https://www.code-recipe.com/post/palindrome-number

package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter no")
	fmt.Scanln(&num)
	fmt.Println(isPalindrome(num))
}

func isPalindrome(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reversedNum
}

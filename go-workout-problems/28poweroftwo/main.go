//231. Power of Two
/*

https://leetcode.com/problems/power-of-two/description/

Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.



Example 1:

Input: n = 1
Output: true
Explanation: 20 = 1

Could you solve it without loops/recursion?

*/

package main

import "fmt"

func main() {

	fmt.Println(isPowerOfTwo(32))
}

/*

func isPowerOfTwo(n int) bool {
	flag := 0
	for n > 1 {
		if n%2 == 0 {
			n = n % 2
		} else {
			flag = 1
			break
		}

	}
	if flag == 0 {
		return true
	} else {
		return false
	}

}

*/

// Could you solve it without loops/recursion?

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	if (n & (n - 1)) == 0 { // Here it is bitwise operation is done because single & isused (if n=8, then 1000 & 0111 = 0000)
		return true
	}

	return false

}

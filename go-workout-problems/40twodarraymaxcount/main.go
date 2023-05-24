/*

1351. Count Negative Numbers in a Sorted Matrix
Easy
3.3K
93
Companies
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.



Example 1:

Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.
Example 2:

Input: grid = [[3,2],[1,0]]
Output: 0

https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/description/
*/

package main

import "fmt"

func main() {

	//var grid = [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}

	var grid2 = [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}}
	fmt.Println(countNegatives(grid2))
}

func countNegatives(grid [][]int) int {
	count := 0
	rows := len(grid)    // Number of rows
	cols := len(grid[0]) // Number of columns (assuming all rows have the same length)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}

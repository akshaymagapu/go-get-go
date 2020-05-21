/**
Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.

Example 1:
Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
*/
package exercise

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	}
	return z
}

func countSquares(matrix [][]int) int {
	count := 0
	rowL := len(matrix)
	colL := len(matrix[0])
	for i := 0; i < rowL; i++ {
		for j := 0; j < colL; j++ {
			if i > 0 && j > 0 && matrix[i][j] != 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1])
			}
			count += matrix[i][j]
		}
	}
	return count
}

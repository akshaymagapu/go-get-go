/**
We write the integers of A and B (in the order they are given) on two separate horizontal lines.

Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:

A[i] == B[j];
The line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.

Return the maximum number of connecting lines we can draw in this way.
*/

package exercise

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxUncrossedLines(A []int, B []int) int {
	a_length := len(A)
	b_length := len(B)

	crossedMat := make([][]int, a_length+1)

	for i := 0; i <= a_length; i++ {
		crossedMat[i] = make([]int, b_length+1)
	}
	for i := 1; i <= a_length; i++ {
		for j := 1; j <= b_length; j++ {
			if A[i-1] == B[j-1] {
				crossedMat[i][j] = crossedMat[i-1][j-1] + 1
			} else {
				crossedMat[i][j] = max(crossedMat[i-1][j], crossedMat[i][j-1])
			}
		}
	}
	return crossedMat[a_length][b_length]
}

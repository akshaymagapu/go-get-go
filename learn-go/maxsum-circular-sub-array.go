/**
Given a circular array C of integers represented by A, find the maximum possible sum of a non-empty subarray of C.

Here, a circular array means the end of the array connects to the beginning of the array.
(Formally, C[i] = A[i] when 0 <= i < A.length, and C[i+A.length] = C[i] when i >= 0.)

Also, a subarray may only include each element of the fixed buffer A at most once.
(Formally, for a subarray C[i], C[i+1], ..., C[j], there does not exist i <= k1, k2 <= j with k1 % A.length = k2 % A.length.)

Example 1:

Input: [1,-2,3,-2]
Output: 3
Explanation: Subarray [3] has maximum sum 3
*/

package exercise

import "sort"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func kadane(data []int, negate bool) int {

	maxSum := 0
	temp := 0

	for _, val := range data {
		cur_val := val
		if negate {
			cur_val = -val
		}

		temp = max(0, temp+cur_val)
		maxSum = max(maxSum, temp)
	}

	return maxSum
}

func maxSubarraySumCircular(A []int) int {

	sortedArr := make([]int, len(A))
	copy(sortedArr, A)
	sort.Ints(sortedArr)
	if sortedArr[len(sortedArr)-1] < 0 {
		return sortedArr[len(sortedArr)-1]
	}

	totalSum := 0

	for _, val := range A {
		totalSum = totalSum + val
	}

	return max(kadane(A, false), totalSum+kadane(A, true))

}

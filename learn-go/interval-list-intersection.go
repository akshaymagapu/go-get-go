/**
Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

(Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.
The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.
For example, the intersection of [1, 3] and [2, 4] is [2, 3].)

Example 1:

Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.
*/

package exercise

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	a_length := len(A)
	b_length := len(B)

	intersection := make([][]int, 0)

	i := 0
	j := 0

	for i < a_length && j < b_length {

		lowNum := max(A[i][0], B[j][0])
		amax := A[i][1]
		bmax := B[j][1]
		highNum := min(amax, bmax)

		if amax < bmax {
			i++
		} else {
			j++
		}

		if lowNum <= highNum {
			intersection = append(intersection, []int{lowNum, highNum})
		}
	}

	return intersection
}

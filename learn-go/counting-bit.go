/**
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's
in their binary representation and return them as an array.

Example 1:
Input: 2
Output: [0,1,1]
*/
package exercise

func countBits(num int) []int {
	result := make([]int, num+1)

	for i := 1; i <= num; i++ {
		result[i] = result[(i-1)&i] + 1
	}

	return result
}

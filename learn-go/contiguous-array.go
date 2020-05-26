/**
Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

Example 1:
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
*/

package exercise

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	countsMap := map[int]int{
		0: -1,
	}
	lenth := len(nums)
	count := 0
	maxLen := 0

	for i := 0; i < lenth; i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		value, ok := countsMap[count]
		if !ok {
			countsMap[count] = i
		} else {
			maxLen = max(maxLen, i-value)
		}
	}
	return maxLen
}

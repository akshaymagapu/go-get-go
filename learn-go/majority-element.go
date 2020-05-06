/**
Given an array of size n, find the majority element.
The majority element is the element that appears more than ⌊ n/2 ⌋ times.
You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:
Input: [3,2,3]
Output: 3
*/

package exercise

func majorityElement(nums []int) int {

	result := map[string]int{
		"count": 0,
		"ans":   0,
	}
	count := map[int]int{}

	for _, val := range nums {
		count[val]++
		if count[val] > result["count"] {
			result["count"] = count[val]
			result["ans"] = val
		}
	}

	return result["ans"]
}

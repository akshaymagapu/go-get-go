// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Input: [2,2,1]
// Output: 1

package exercise

func singleNumber(nums []int) int {

	res := 0

	for _, val := range nums {
		res ^= val
	}

	return res
}

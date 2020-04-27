// Given an array nums of n integers where n > 1,
// return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
// Input:  [1,2,3,4]
// Output: [24,12,8,6]

package exercise

func productExceptSelf(nums []int) []int {

	product := 1
	size := len(nums)
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = product
		product *= nums[i]
	}

	product = 1
	for i := size - 1; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}
	return result
}

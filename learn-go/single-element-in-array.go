/**
You are given a sorted array consisting of only integers where every element appears exactly twice,
except for one element which appears exactly once. Find this single element that appears only once.

Example 1:
Input: [1,1,2,3,3,4,4,8,8]
Output: 2
*/

package exercise

func singleNonDuplicate(nums []int) int {

	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2

		if mid % 2 != 0 {
			mid--
		}

		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}

	}

	return nums[low]
}

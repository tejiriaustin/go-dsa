package two_pointer_techniques

// Problem:
//Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.
//Do not allocate extra space for another array; you must do this by modifying the input array in-place with O(1) extra memory.

func RemoveDuplicates(nums []int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[left] == nums[right] {
			right++
			continue
		}
		left++
		nums[left] = nums[right]
		right++
	}
	return left + 1
}

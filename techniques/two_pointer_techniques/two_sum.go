package two_pointer_techniques

// Problem:
// Given an array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number. Return the indices of the two numbers (1-indexed) as an integer array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.
// You may assume that each input would have exactly one solution and if you may not use the same element twice.

// Example:
//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore, index 1 and 2 are returned.

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < len(nums) {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target && sum > 0 {
			left++
		} else if sum > target && sum > 0 {
			right--
		} else {
			left++
		}
	}

	return []int{0, 0}
}

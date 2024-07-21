package two_pointer_techniques

func twoSumUnordered(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < len(nums) {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum != target && right != left+1 {
			right--
		} else {
			left++
			right = len(nums) - 1
		}
	}

	return []int{0, 0}
}

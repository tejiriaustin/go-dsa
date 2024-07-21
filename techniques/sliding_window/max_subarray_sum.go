package sliding_window

// Problem: Maximum Sum Subarray of Size k
// Given an array of integers and an integer k, find the maximum sum of any subarray of size k.

func maxSubArraySum(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	for i := k; i < len(nums); i++ {
		windowSum := maxSum + nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

package sliding_window

// Problem: Longest Substring Without Repeating Characters
// Given a string, find the length of the longest substring without repeating characters.

func longestSubstring(s string) int {
	ru := []rune(s)

	left, right := 0, 0
	var answer int

	for left < len(ru)-1 {
		for i := left; i <= right; i++ {
			if ru[left] == ru[right] {
				left++
				i = left
				continue
			}
		}
		answer++
		right++
	}
	return answer
}

package two_pointer_techniques

func isAnagram(s string) bool {
	ru := []rune(s)
	left, right := 0, len(ru)-1

	for left < right {
		if ru[left] != ru[right] {
			return false
		}
		left++
		right--
	}
	return true
}

package two_pointer_techniques

func moveZero(list []int) []int {
	left, right := 0, len(list)-1

	for left < right {
		if list[left] == 0 {
			list[left], list[right] = list[right], list[left]
		}
		left++
	}
	return list
}

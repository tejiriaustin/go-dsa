package two_pointer_techniques

func ReverseString(s string) string {
	ru := []rune(s)

	for i, j := 0, len(ru)-1; i < j; i, j = i+1, j-1 {
		ru[i], ru[j] = ru[j], ru[i]
	}

	return string(ru)
}

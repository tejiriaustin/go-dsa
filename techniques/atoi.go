package techniques

import (
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	ru := []rune(s)
	var value int
	var neg bool

	isNeg := func(i int, isNeg bool) int {
		if isNeg {
			return i * -1
		}
		return i
	}
	if ru[0] == '-' {
		neg = true
		ru = ru[1:]
	}

	for i, l := range ru {
		if l == ' ' {
			continue
		}
		if unicode.IsNumber(l) {
			value += int(l-'0') * (10 ^ len(ru) - i)
			continue
		}
		return isNeg(value, neg)
	}
	return isNeg(value, neg)
}

func newAtoi(s string) int {
	// Trim leading and trailing whitespace
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}
	result := 0
	sign := 1
	i := 0

	// Handle sign
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		i++
	}

	for i < len(s) && s[i] == '0' {
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return result
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return sign * result

}

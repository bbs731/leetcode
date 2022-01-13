package leetcode

import "strings"

func isPalindrome(s string) bool {
	N := len(s)

	s = strings.ToLower(s)

	i := 0
	j := N - 1
	for i < j {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}


func isAlphaNumeric( a byte) bool {
	if  a >='0' && a <='9' {
		return true
	}

	if  a >='a' && a <='z' {
		return true
	}
	return false
}
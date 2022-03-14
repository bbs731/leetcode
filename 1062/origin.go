package leetcode

import (
	"math"
)

// use Rabin-Karp algorithms to check whether text has L's length repeated substring
func search(m int, text []int, q int) bool {
	n := len(text)
	base := 26
	h := 1
	p := 0
	seen := make(map[int]bool)

	if n < m {
		return false
	}
	// h = pow(base, m-1)%q
	for i := 0; i < m-1; i++ {
		h = (h * base) % q
	}

	// calculate the hash value of pattern start at text' 0 index
	for i := 0; i < m; i++ {
		p = (base*p + text[i]) % q
	}
	for i := 0; i <= n-m; i++ {
		if _, ok := seen[p]; ok {
			return true
		}
		seen[p] = true

		if i < n-m {
			p = (base*(p-text[i]*h) + text[i+m]) % q
			if p < 0 {
				p = p + q
			}
		}
	}
	return false
}

func longestRepeatingSubstring(s string) int {
	n := len(s)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - 'a')
	}

	left, right := 1, n
	modules := int(math.Pow(2, 24))

	for left <= right {
		L := left + (right-left)>>1
		if search(L, nums, modules) {
			left = L + 1
		} else {
			right = L - 1
		}
	}
	return left - 1
}

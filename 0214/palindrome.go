package leetcode

import (
	"math"
	"strings"
)

func shortestPalindrome(s string) string {
	n := len(s)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - '0')
	}

	if n == 0 {
		return ""
	}
	h := 0
	r := 0
	module := int(math.Pow(2, 24))
	//module := 1000000007
	base := 131
	bases := make([]int, n)
	bases[0] = 1

	for i := 1; i < n; i++ {
		bases[i] = bases[i-1] * base % module
	}

	// 我们来计算 s[i] 是不是 palindrome
	best := 0
	for i := 0; i < n; i++ {
		h = (h*base + nums[i]) % module
		r = (nums[i]*bases[i] + r) % module
		if h == r {
			// s[i] and reverse s[i] hash is equal. so s[:i+1] is palindrome
			best = i
		}
	}

	if best == n-1 {
		return s
	}

	return reverse(s[best+1:]) + s
}

func reverse(s string) string {
	n := len(s)
	b := strings.Builder{}
	for i := n - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}

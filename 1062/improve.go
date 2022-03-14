package leetcode

import (
	"fmt"
	"math"
)

// use Rabin-Karp algorithms to check whether text has L's length repeated substring
func search(L int, nums []int, q int) bool {
	n := len(nums)
	base := 26
	h := 0
	aL := 1
	seen := make(map[int]bool)

	for i := 0; i < L; i++ {
		h = (h*base + nums[i]) % q
	}
	seen[h] = true

	// calculate the hash value of pattern start at text' 0 index
	for i := 1; i <= L; i++ {
		aL = (aL * base) % q
	}


	for start := 1; start < n-L+1; start++ {
		h = (h*base - nums[start-1]*aL%q + q) % q
		h = (h + nums[start+L-1]) % q
		if _, ok := seen[h]; ok {
			return true
		}
		seen[h] = true
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
	return left-1
}

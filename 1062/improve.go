package leetcode

import (
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
	// 如何选取 modules 的值？  需要根据，生日悖论：  https://baike.baidu.com/item/%E7%94%9F%E6%97%A5%E6%82%96%E8%AE%BA
	// n ~ O(N^1/2) 此时 n = 1500 最大可能出现的子串长度，  N = n^2 = 1500 * 1500 = 2250000 = 2.4* 10 ^6   所以，这里可以用 2^24 来近似
	modules := int(math.Pow(2, 24))

	for left < right {
		L := left + (right-left)>>1
		if search(L, nums, modules) {
			left = L + 1
		} else {
			right = L
		}
	}
	return left - 1
}

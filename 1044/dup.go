package leetcode

import "math"

// 这道题，给我们的经验，就是，
// base 要取质数， 然后， mod 需要足够大， 这道题，需要至少 2 ^ 36
// 或者，按照，官方leetcode题解的意思，使用双哈希， 或者 k 哈希，类似bloom-filter 去重一样。因为这道题， hash conflict 很严重啊！
func longestDupSubstring(s string) string {
	n := len(s)
	left, right := 1, n-1

	if n <= 1 {
		return ""
	}

	mod := int(math.Pow(2, 36)) // 这道题，告诉我们， n = 3*10^4  所以至少  mod 需要取 9*10^9 左右 mod 需要足够大！
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - 'a')
	}
	var ans string

	search := func(L int, nums []int, mod int) (bool, int) {
		seen := make(map[int]bool)
		base := 31 // 要取一个质数
		bases := make([]int, L)
		h := 0

		for i := 0; i < L; i++ {
			h = (h*base + nums[i]) % mod
		}
		bases[0] = 1
		for i := 1; i < L; i++ {
			bases[i] = bases[i-1] * base
		}
		seen[h] = true

		for i := 1; i < n; i++ {
			j := i + L - 1
			if j < n {
				h = ((h - nums[i-1]*bases[L-1]%mod) + mod) % mod
				h = (h*base + nums[j]) % mod

				if _, ok := seen[h]; ok {
					return true, i
				} else {
					seen[h] = true
				}
			}
		}

		return false, -1
	}

	for left < right {
		L := left + (right-left)>>1

		if ok, pos := search(L, nums, mod); ok {
			left = L + 1
			ans = s[pos : pos+L]
		} else {
			right = L
		}
	}
	return ans

}

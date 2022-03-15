package leetcode

import "math"

func findRepeatedDnaSequences(s string) []string {
	digits := map[byte]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	n := len(s)
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = digits[s[i]]
	}
	ansmap := make(map[string]struct{})
	ans := make([]string, 0)
	if n <= 10 {
		return ans
	}
	search := func(L int, nums []int, module int) {
		base := 4
		al := 1
		h := 0
		seen := make(map[int]bool)
		for i := 0; i < L; i++ {
			al = al * base % module
			h = (h*base + nums[i]) % module
		}
		seen[h] = true
		for start := 1; start+L-1 < n; start++ {
			h = (h*base - (nums[start-1]*al)%module + module) % module // 这里容易出错， h 需要乘以 base
			h = (h + nums[start+L-1]) % module
			if _, ok := seen[h]; ok {
				//ans = append(ans, s[start:start+L])
				ansmap[s[start:start+L]] = struct{}{}
			} else {
				seen[h] = true
			}
		}
	}
	search(10, nums, int(math.Pow(2, 30)))
	for k := range ansmap {
		ans = append(ans, k)
	}
	return ans
}

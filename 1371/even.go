package leetcode

/*
typical prefixsum problem
 */
func findTheLongestSubstring(s string) int {
	n := len(s)
	prefixsum := make([]int, n+1)
	saved := make(map[int]int)

	ans := 0

	dict := map[byte]int{
		'a': 1,
		'e': 2,
		'i': 4,
		'o': 8,
		'u': 16,
	}
	saved[0] = 0
	for i := 1; i <= n; i++ {
		if v, ok := dict[s[i-1]]; ok { // one of a e i o u
			if prefixsum[i-1]&v > 0 {
				prefixsum[i] = prefixsum[i-1] & (^v)
			} else {
				prefixsum[i] = prefixsum[i-1] | v
			}

		} else {
			prefixsum[i] = prefixsum[i-1]
		}

		if index, ok := saved[prefixsum[i]]; !ok {
			saved[prefixsum[i]] = i
		} else {
			ans = max(ans, i-index)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

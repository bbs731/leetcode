package leetcode

// implement Rabin-Karp algorithms  又名： 滚动哈希
// return the starting index that pattern occurs in text
func search(text, pattern string, q int) []int {
	pos := make([]int, 0)
	m := len(pattern)
	n := len(text)
	base := 256
	h := 1
	p, t := 0, 0

	if n < m {
		return pos
	}
	// h = pow(base, m-1)%q
	for i := 0; i < m-1; i++ {
		h = (h * base) % q
	}

	// calculate the hash value of pattern and first window of text
	for i := 0; i < m; i++ {
		p = (base*p + int(pattern[i])) % q
		t = (base*t + int(text[i])) % q
	}

	for i := 0; i <= n-m; i++ {
		if p == t { // hash value equal then check the string
			if text[i:i+m] == pattern {
				pos = append(pos, i)
			}
		}

		if i < n-m {
			t = (base*(t-int(text[i])*h) + int(text[i+m])) % q
			if t < 0 {
				t = t + q
			}
		}
	}
	return pos
}

func longestRepeatingSubstring(s string) int {
	n := len(s)
	ans := 0

	// 这是朴素的算法， 时间复杂度是 O(n^2) 会超时.
	for l := 1; l <= n-1; l++ {
		for i := 0; i < n; i++ {
			j := l - 1 + i
			if j < n {
				pattern := s[i : j+1]
				indexes := search(s[i+1:], pattern, 101)
				if len(indexes) > 0 {
					ans = l
				}

			}
		}
	}
	return ans
}

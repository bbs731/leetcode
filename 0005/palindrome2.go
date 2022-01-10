package leetcode

func longestPalindrome(s string) string {
	N := len(s)
	ans := s[:1]
	r := make([][]bool, N)
	for i := 0; i < N; i++ {
		r[i] = make([]bool, N)
		r[i][i] = true
	}

	for l := 1; l <= N; l++ {
		for i := 0; i < N-l; i++ {
			j := i + l
			if i+1 <= j-1 {
				if r[i+1][j-1] {
					r[i][j] = s[i] == s[j]
					if r[i][j] && l > len(ans) {
						ans = s[i : j+1]
					}
				} else {
					r[i][j] = false
				}
			} else {
				// l == 1
				if s[i] == s[j] {
					r[i][j] = true
					if l > len(ans) {
						ans = s[i : j+1]
					}
				} else {
					r[i][j] = false
				}
			}
		}
	}
	return ans
}

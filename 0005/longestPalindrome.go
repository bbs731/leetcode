package leetcode

type subString struct {
	start int
	end   int
}

func longestPalindrome2(s string) string {
	res := s[:1]
	candidates := make([]subString, 0, 2*len(s))
	holders := make([]subString, 0, 2*len(s))
	// initialized palindrome of len 1 and 2
	for i := 0; i < len(s); i++ {
		candidates = append(candidates, subString{i, i})
		j := i + 1
		if j < len(s) && s[i] == s[j] {
			candidates = append(candidates, subString{i, j})
			if j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}

	for len(candidates) > 0 {
		for _, p := range candidates {
			i := p.start - 1
			j := p.end + 1

			if i >= 0 && j < len(s) && s[i] == s[j] {
				holders = append(holders, subString{i, j})
				if j-i+1 > len(res) {
					res = s[i : j+1]
				}
			}
		}
		candidates = holders
		holders = holders[:0]
	}
	return res
}

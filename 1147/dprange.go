package leetcode

func longestDecomposition(text string) int {
	n := len(text)
	if n == 0 {
		return 0
	}

	ans := 0

	candidates := make([]int, 0)
	for j := n - 1; j >= 0; j-- {
		if text[j] == text[0] {
			candidates = append(candidates, j)
		}
	}

	if candidates[0] == 0 {
		return 1 // text as a whole
	}

	for i := 1; i < len(candidates); i++ {
		pos := candidates[i]
		l := n - pos
		if text[:l] == text[pos:] {
			ans = max(ans, longestDecomposition(text[l:pos])+2)
			//match greedy
			break
		} else {
			ans = max(ans, 1)
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

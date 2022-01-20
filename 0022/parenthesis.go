package leetcode


/*
这道 backstracking, 你写了将近40分钟， 多练习吧！
 */
var ans []string

func generate(candidates []byte, pos int, n int) {
	// invariants left >= right
	left, right := 0, 0
	for i := 0; i < pos; i++ {
		if candidates[i] == '(' {
			left++
		} else {
			right++
		}
		if right > left || left > n {
			// invalid
			// break invariants
			return
		}
	}

	// valid ans save it
	if pos == 2*n {
		ans = append(ans, string(candidates[:2*n]))
		return
	}

	candidates[pos] = '('
	pos++
	generate(candidates, pos, n)
	pos--

	candidates[pos] = ')'
	pos++
	generate(candidates, pos, n)
	pos--

	return
}

func generateParenthesis(n int) []string {
	candidates := make([]byte, 2*n)
	candidates[0] = '('

	ans = make([]string, 0)
	generate(candidates, 1, n)
	return ans
}

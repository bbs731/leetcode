package leetcode

/*
stack will overflow , with no cache
 */
func wordsTyping(sentence []string, rows int, cols int) int {

	l := len(sentence)

	ans := 0
	var dfs func(int, int, int)
	dfs = func(r int, c int, pos int) {

		if pos == l { // mod to zero
			pos = 0
		}

		if r == rows {
			return
		}

		if c >= cols {
			dfs(r+1, 0, pos)
			return
		}

		if c+len(sentence[pos])-1 < cols {
			if pos == l-1 {
				ans++
			}
			dfs(r, c+len(sentence[pos])+1, pos+1)
			return
		} else {
			dfs(r+1, 0, pos)
			return
		}
	}

	dfs(0, 0, 0)
	return ans

}

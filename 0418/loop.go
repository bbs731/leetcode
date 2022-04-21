package leetcode

func wordsTyping(sentence []string, rows int, cols int) int {

	l := len(sentence)
	ans := 0

	r := 0
	pos := 0
	c := 0
	for r < rows {
		if c >= cols {
			r++
			c = 0
			continue
		}

		if c == 0 && pos == 0 && r != 0 {
			// here is a cycle
			ans = rows / r * ans
			rows = rows % r
			r = 0
			continue
		}

		if c+len(sentence[pos])-1 < cols { // have enough space
			c = c + len(sentence[pos]) + 1
			pos++
			if pos == l {
				pos = 0
				ans++

			}
		} else {
			r++
			c = 0
		}
	}
	return ans
}

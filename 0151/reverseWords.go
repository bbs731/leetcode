package leetcode

func reverseWords(s string) string {
	stack := make([]string, 0)
	start, end := 0, 0
	ans := ""

	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			// possible word end pos
			end = i + 1
		} else {
			if end > start {
				stack = append(stack, s[start:end])
			}
			// possible wordstart from i + 1
			start = i + 1
		}
	}

	if end > start {
		stack = append(stack, s[start:end])
	}

	for j := len(stack) - 1; j > 0; j-- {
		ans += stack[j]
		ans += " "
	}

	ans += stack[0]

	return ans
}

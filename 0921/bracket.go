package leetcode

func minAddToMakeValid(s string) int {
	stack := make([]byte, 0) // save all the '('
	ans := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				ans++
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	ans += len(stack)
	return ans
}

package leetcode

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	ans := 0
	carry := 1
	for i := n - 1; i >= 0; i-- {
		ans += (int(columnTitle[i]-'A') + 1) * carry
		carry = carry * 26
	}
	return ans
}

package leetcode

func longestPalindrome(s string) int {

	count := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	odd := false
	ans := 0
	for _, v := range count {
		if v%2 == 0 {
			ans += v
		} else {
			ans += v - 1
			odd = true
		}
	}

	if odd {
		ans++
	}
	return ans
}

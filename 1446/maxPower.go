package leetcode

func maxPower(s string) int {
	max := 1
	cum := 1

	if len(s) == 0 {
		return 0
	}

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cum++
		} else {
			cum = 1
		}

		if cum > max {
			max = cum
		}
	}

	return max
}

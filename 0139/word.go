package leetcode

var dict map[string]bool

/*
经典题目啊， 为啥能一次过啊， 走狗屎运了吗？
 */
func wordBreak(s string, wordDict []string) bool {
	dict = make(map[string]bool)

	for _, s := range wordDict {
		dict[s] = true
	}

	N := len(s)
	dp := make([]bool, N)
	record := make([]int, 0, N)

	for i := 0; i < N; i++ {
		candidate := s[:i+1]
		if _, ok := dict[candidate]; ok {
			dp[i] = true
			record = append(record, i)
			continue
		}

		// otherwise
		for j := 0; j < len(record); j++ {
			pos := record[j]
			candidate = s[pos+1 : i+1]
			if _, ok := dict[candidate]; ok {
				dp[i] = true
				record = append(record, i)
				break
			}
		}
	}

	return dp[N-1]
}

package leetcode

// 这道题， 可以用DP 来解答！

// 哎看看，这一坨屎一样的答案！
func canCross(stones []int) bool {
	// stones must start with 0 and 1
	// we use a start and end index to represent a queue

	n := len(stones)
	dp := make([]map[int]struct{}, n)
	start, end := 1, 1 // points to stones index at 1
	//dp[1] = 1          // dp[i] stands for stones at index 1 with last jump max step's number
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]struct{})
	}
	dp[1][1] = struct{}{}
	if stones[1] != 1 {
		return false
	}

	for start <= end {
		// top of the queue
		top := start
		for step := range dp[top] {
			steps := step + 1 // far we can reach
			for i := top + 1; i < n; i++ {
				if stones[top]+steps < stones[i] {
					break
				}
				// this far we can reach
				if stones[i]-stones[top] >= steps-2 {
					//dp[i] = max(dp[i], min(steps, stones[i]-stones[top]))
					dp[i][stones[i]-stones[top]] = struct{}{}
					if i > end {
						// update end if necessary
						end = i
					}
				}
			}
		}
		start++
	}

	return len(dp[n-1]) != 0

}


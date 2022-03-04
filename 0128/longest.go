package leetcode

func longestConsecutive(nums []int) int {

	graph := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		graph[nums[i]] = false
	}

	var dfs func(int) int // start pos and return dfs travel length
	dfs = func(start int) int {
		if visited, ok := graph[start]; !ok || visited {
			return 0
		}
		graph[start] = true
		return dfs(start-1) + dfs(start+1) + 1
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dfs(nums[i]))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

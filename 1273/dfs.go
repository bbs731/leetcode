package leetcode

/*
两遍 dfs,  第一遍统计 子树的 sum, 第二遍，遍历统计 nodes 的个数
 */
func deleteTreeNodes(nodes int, parent []int, value []int) int {

	// use parent to calculate children
	children := make([][]int, nodes)
	for i := 1; i < nodes; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	sums := make([]int, nodes)
	visited := make([]bool, nodes)

	var dfs func(int) int
	dfs = func(root int) int {
		visited[root] = true
		sum := value[root]
		for _, c := range children[root] {
			if !visited[c] {
				sum += dfs(c)
			}
		}
		sums[root] = sum
		return sum
	}

	dfs(0)

	ans := 0
	visited = make([]bool, nodes)
	if sums[0] == 0 {
		return ans
	}
	var cdfs func(int)
	cdfs = func(root int) {
		visited[root] = true
		for _, c := range children[root] {
			if !visited[c] && sums[c] != 0 {
				cdfs(c)
			}
		}
		ans++
	}
	cdfs(0)
	return ans
}

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(root *TreeNode) int {

	take := make(map[*TreeNode]int)
	notake := make(map[*TreeNode]int)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		take[root] = root.Val + notake[root.Left] + notake[root.Right]
		notake[root] = max(take[root.Left], notake[root.Left]) + max(take[root.Right], notake[root.Right])
	}

	dfs(root)
	return max(take[root], notake[root])
}

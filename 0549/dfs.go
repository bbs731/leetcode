package leetcode

func longestConsecutive(root *TreeNode) int {

	increase := make(map[*TreeNode]int)
	decrease := make(map[*TreeNode]int)

	if root == nil {
		return 0
	}

	ans := 1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)

		if root.Left != nil && root.Left.Val+1 == root.Val {
			decrease[root] = max(decrease[root], decrease[root.Left]+1)
		}
		if root.Right != nil && root.Right.Val+1 == root.Val {
			decrease[root] = max(decrease[root], decrease[root.Right]+1)
		}

		if root.Left != nil && root.Left.Val == root.Val+1 {
			increase[root] = max(increase[root], increase[root.Left]+1)
		}
		if root.Right != nil && root.Right.Val == root.Val+1 {
			increase[root] = max(increase[root], increase[root.Right]+1)
		}

		total := 0

		if increase[root] != 0 {
			if decrease[root] != 0 {
				total = increase[root] + decrease[root] + 1
			} else {
				total = increase[root] + 1
			}
		} else {
			if decrease[root] != 0 {
				total = decrease[root] + 1
			}
		}

		ans = max(ans, total)
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

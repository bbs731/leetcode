package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


/*
写的像 shit 一样。 这样答题，是过不去的。
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(root *TreeNode) int {

	parentsNotaken := make(map[*TreeNode]int)
	parentsTaken := make(map[*TreeNode]int)

	var dfs func(*TreeNode, bool) int
	dfs = func(root *TreeNode, parentTaken bool) int {
		if root == nil {
			return 0
		}
		// if cached return the value directly
		if parentTaken {
			if _, ok := parentsTaken[root]; ok {
				return parentsTaken[root]
			}
		} else {
			if _, ok := parentsNotaken[root]; ok {
				return parentsNotaken[root]
			}
		}

		if _, ok := parentsNotaken[root.Left]; !ok {
			parentsNotaken[root.Left] = dfs(root.Left, false)
		}
		if _, ok := parentsNotaken[root.Right]; !ok {
			parentsNotaken[root.Right] = dfs(root.Right, false)
		}

		parentsTaken[root] = parentsNotaken[root.Left] + parentsNotaken[root.Right]

		if _, ok := parentsTaken[root.Left]; !ok {
			parentsTaken[root.Left] = dfs(root.Left, true)
		}
		if _, ok := parentsTaken[root.Right]; !ok {
			parentsTaken[root.Right] = dfs(root.Right, true)
		}

		// store the cache
		parentsNotaken[root] = max(parentsNotaken[root.Left]+parentsNotaken[root.Right], root.Val+parentsTaken[root.Left]+parentsTaken[root.Right])
		return parentsNotaken[root]
	}

	return dfs(root, false)
}

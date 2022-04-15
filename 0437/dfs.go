package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {

	counts := 0
	var dfs func(*TreeNode, []int)
	dfs = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		// check answer
		sum := 0
		for i := len(path) - 1; i >= 0; i-- {
			sum += path[i]
			if sum == targetSum {
				counts++
			}
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
		//restore path
		path = path[:len(path)-1]
	}

	dfs(root, []int(nil))
	return counts
}

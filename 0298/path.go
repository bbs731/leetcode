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

func longestConsecutive(root *TreeNode) int {

	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		ln := dfs(n.Left)
		rn := dfs(n.Right)
		if n.Left != nil && n.Val+1 == n.Left.Val {
			ln++
		} else {
			ln = 1
		}
		if n.Right != nil && n.Val+1 == n.Right.Val {
			rn++
		} else {
			rn = 1
		}
		ans = max(ans, ln, rn)
		return max(ln, rn)
	}

	dfs(root)
	return ans
}

func max(nums ...int) int {
	m := 0
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

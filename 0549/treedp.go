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
	var dfs func(*TreeNode) (int, int)

	dfs = func(n *TreeNode) (int, int) {

		if n == nil {
			return 0, 0
		}

		var ca, cd int
		la, ld := dfs(n.Left)
		ra, rd := dfs(n.Right)

		if n.Left != nil && n.Left.Val+1 == n.Val {
			ld++
		} else {
			ld = 1 // here is tricky
		}
		if n.Left != nil && n.Left.Val == n.Val+1 {
			la++
		} else {
			la = 1
		}
		if n.Right != nil && n.Right.Val+1 == n.Val {
			rd++
		} else {
			rd = 1
		}
		if n.Right != nil && n.Right.Val == n.Val+1 {
			ra++
		} else {
			ra = 1
		}

		ca = max(la, ra)
		cd = max(ld, rd)

		ans = max(ans, ca)
		ans = max(ans, cd)

		// consider special case
		if n.Left != nil && n.Right != nil {
			if n.Left.Val+1 == n.Val && n.Val+1 == n.Right.Val {
				ans = max(ans, ld+ra-1)
			}
			if n.Left.Val-1 == n.Val && n.Val == n.Right.Val+1 {
				ans = max(ans, rd+la-1)
			}
		}
		return ca, cd
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

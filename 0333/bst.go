package leetcode

/*
错误的答案， 看 bst1.go
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestBSTSubtree(root *TreeNode) int {
	ans := 0

	if root != nil {
		ans = 1
	}

	var dfs func(*TreeNode) (int, int, int)
	dfs = func(r *TreeNode) (int, int, int) {
		var lmin, lmax, lsize, rmin, rmax, rsize, cmin, cmax, csize int
		if r.Left != nil && r.Left.Val < r.Val {
			lmin, lmax, lsize = dfs(r.Left)
			if lmax < r.Val {
				csize = lsize + 1
				ans = max(ans, csize)
				cmin = lmin
			} else {
				csize = 2
				cmin = r.Left.Val
			}
		} else {
			csize = 1
			cmin = r.Val
		}

		if r.Right != nil && r.Right.Val > r.Val {
			rmin, rmax, rsize = dfs(r.Right)
			if rmin > r.Val {
				csize = csize + rsize
				cmax = rmax
			} else {
				csize = csize + 1
				ans = max(ans, csize)
				cmax = r.Right.Val
			}
		} else {
			cmax = r.Val
		}

		return cmin, cmax, csize
	}

	if root != nil {
		dfs(root)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

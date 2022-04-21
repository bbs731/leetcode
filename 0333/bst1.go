package leetcode

/*
一个 binary search 一个 dfs 好像是你的田地 ！
 */
func largestBSTSubtree(root *TreeNode) int {
	ans := 0

	if root != nil {
		ans = 1
	}

	var dfs func(*TreeNode) (int, int, int, bool)
	dfs = func(r *TreeNode) (int, int, int, bool) {
		var lmin, lmax, lsize, rmin, rmax, rsize, csize int
		var lb, rb bool

		/*
		只要， r.Left, 或者 r.Right 不是 Nil 就需要 traverse, 因为，  BST 可能只是下面的子树
		 */
		if r.Left != nil {
			lmin, lmax, lsize, lb = dfs(r.Left)
			if lmax >= r.Val {
				lb = false
			}

		} else {
			// r.Left is nil
			lmin = r.Val
			lsize = 0
			lb = true
		}

		if r.Right != nil {
			rmin, rmax, rsize, rb = dfs(r.Right)
			if rmin <= r.Val {
				rb = false
			}
		} else {
			//r. Right is nil
			rmax = r.Val
			rsize = 0
			rb = true
		}

		if lb && rb {
			csize = 1 + lsize + rsize
			ans = max(ans, csize)
			return lmin, rmax, csize, true
		}
		return 0, 0, 0, false

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

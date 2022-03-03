package leetcode

import "math"

/**
amazing 为啥能一次救过呢！
思路： 对于任何一条路径，都可以找到一个 node, 然后这条路径的一半在这个node 的左树上，一半路径在这个node 的右树上。当然路径的长度可以是0
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var table map[*TreeNode][]int // 0 for left and 1 for right
// return the max sum from the tree, max sum =  max(left sum, right sum) + root.val
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		table[root] = []int{0, 0}
		return root.Val
	}

	ls := dfs(root.Left)
	rs := dfs(root.Right)
	table[root] = []int{ls, rs}
	if max(ls, rs) <= 0 {
		return root.Val
	}
	return max(ls, rs) + root.Val
}

func maxPathSum(root *TreeNode) int {
	table = make(map[*TreeNode][]int)
	ans := math.MinInt64
	dfs(root)
	for k, v := range table {
		sum := 0
		if k != nil {
			if v[0] > 0 {
				sum += v[0]
			}
			if v[1] > 0 {
				sum += v[1]
			}
			sum += k.Val
			ans = max(ans, sum)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

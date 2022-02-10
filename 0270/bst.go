package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
没有充分利用性质， 有更优秀的解法！ 
 */
func closestValue(root *TreeNode, target float64) int {

	var dfs func(*TreeNode) []float64
	dfs = func(root *TreeNode) []float64 {
		if root == nil {
			return []float64{}
		}

		l := dfs(root.Left)
		l = append(l, float64(root.Val))
		l = append(l, dfs(root.Right)...)
		return l
	}

	nums := dfs(root)

	pos := sort.SearchFloat64s(nums, target)
	if pos == len(nums) {
		return int(nums[pos-1])
	}
	if pos == 0 {
		return int(nums[0])
	}
	if nums[pos]-target > target-nums[pos-1] {
		return int(nums[pos-1])
	} else {
		return int(nums[pos])
	}

}

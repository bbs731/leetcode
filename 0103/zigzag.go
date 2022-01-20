package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	reverse := false
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	ans := make([][]int, 0)

	for len(stack) != 0 {
		row := []int{}
		if !reverse {
			for i := 0; i < len(stack); i++ {
				row = append(row, stack[i].Val)
			}
			reverse = true
		} else {
			for i := len(stack) - 1; i >= 0; i-- {
				row = append(row, stack[i].Val)
			}
			reverse = false
		}
		ans = append(ans, row)
		newstack := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			if stack[i].Left != nil {
				newstack = append(newstack, stack[i].Left)
			}
			if stack[i].Right != nil {
				newstack = append(newstack, stack[i].Right)
			}
		}
		stack = newstack
	}
	return ans
}

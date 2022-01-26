package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {

	ans := make([]int, 0)

	var traverse func(root *TreeNode, num int)
	traverse = func(root *TreeNode, num int) {
		if root.Left == nil && root.Right == nil {
			num = num*10 + root.Val
			ans = append(ans, num)
			return
		}

		if root.Left != nil {
			traverse(root.Left, num*10+root.Val)
		}
		if root.Right != nil {
			traverse(root.Right, num*10+root.Val)
		}
	}

	traverse(root, 0)

	sum := 0
	for i := 0; i < len(ans); i++ {
		sum += ans[i]
	}
	return sum
}

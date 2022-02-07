package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	queue := make([]*TreeNode, 0)
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue = append(queue, root)

	for len(queue) > 0 {

		tmp := make([]*TreeNode, 0)
		level := make([]int, 0)
		for i := 0; i < len(queue); i++ {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		queue = tmp
		ans = append(ans, level)
	}
	return ans
}

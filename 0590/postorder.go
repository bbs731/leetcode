package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	for _, child := range root.Children {
		ans = append(ans, postorder(child)...)
	}

	ans = append(ans, root.Val)
	return ans
}

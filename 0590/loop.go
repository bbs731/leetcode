package leetcode

func postorder(root *Node) []int {
	stack := make([]*Node, 0)
	ans := make([]int, 0)
	var prev *Node

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			if len(root.Children) > 0 {
				root = root.Children[0]
			} else {
				root = nil
			}
		}

		root = stack[len(stack)-1]
		// pop the root out of the stack
		stack = stack[:len(stack)-1]

		if len(root.Children) == 0 || root.Children[len(root.Children)-1] == prev {
			// add into ans
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			// root 重新入栈
			stack = append(stack, root)
			i := 0
			for ; i < len(root.Children); i++ {
				if prev == root.Children[i] {
					break
				}
			}
			if i == len(root.Children)-1 {
				root = nil
			} else {
				root = root.Children[i+1]
			}
		}
	}
	return ans
}

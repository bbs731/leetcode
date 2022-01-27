package leetcode

/*
这个版本，就没有必要了。 看 loop.go,  loop.go 和 binary search tree 的 post-order loop 版本， 框架是一样的
 */
func postorder(root *Node) []int {
	stack := make([]*Node, 0)
	ans := make([]int, 0)
	var prev *Node

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			if len(root.Children) != 0 {
				root = root.Children[0]
			} else {
				root = nil
			}
		}

		root = stack[len(stack)-1]
		// pop the root out of the stack
		stack = stack[:len(stack)-1]

		i := 0
		found := false
		for ; i < len(root.Children); i++ {
			if prev == root.Children[i] {
				found = true
				break
			}
		}

		if !found { // e.g in case that root does not have any children
			ans = append(ans, root.Val)
			prev = root
			root = nil
			continue
		}
		// found
		if i == len(root.Children)-1 {
			// last child
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			// 重新入队
			stack = append(stack, root)
			root = root.Children[i+1]
		}
	}
	return ans
}

package leetcode

/*
	感觉， 这个 recursive 有点难啊！， 一次写不对啊！
 */
func connect_recursive(root *Node, nextLevel *Node) *Node {
	if root.Left == nil {
		return root
	}

	root.Left.Next = root.Right
	root.Right.Next = nextLevel

	root.Left = connect_recursive(root.Left, root.Right.Left)
	if nextLevel == nil {
		root.Right = connect_recursive(root.Right, nil)
	} else {
		root.Right = connect_recursive(root.Right, nextLevel.Left)
	}
	return root

}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	return connect_recursive(root, nil)
}

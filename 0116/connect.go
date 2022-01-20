package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		save := make([]*Node, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				save = append(save, queue[i].Left)
			}
			if queue[i].Right != nil {
				save = append(save, queue[i].Right)
			}
			if i != len(queue)-1 {
				queue[i].Next = queue[i+1]
			}
		}
		queue[len(queue)-1].Next = nil
		queue = save
	}
	return root
}

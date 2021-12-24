package leetcode

var m map[*Node]*Node

func docopy(head *Node) *Node {
	if head == nil {
		return nil
	}
	if n, ok := m[head]; ok {
		return n
	}

	c := new(Node)
	m[head] = c
	c.Val = head.Val
	c.Random = docopy(head.Random)
	c.Next = docopy(head.Next)

	return c
}

func copyRandomList2(head *Node) *Node {
	m = make(map[*Node]*Node)

	return docopy(head)
}

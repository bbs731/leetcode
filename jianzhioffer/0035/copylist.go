package offer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
一遍，写对，还是不容易啊， 多练吧！
 */
var table map[*Node]*Node

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	if table == nil {
		table = make(map[*Node]*Node)
	}

	if _, ok := table[head]; !ok {
		copyhead := &Node{
			Val: head.Val,
		}
		table[head] = copyhead
		copyhead.Next = copyRandomList(head.Next)
		copyhead.Random = copyRandomList(head.Random)
	}

	return table[head]
}

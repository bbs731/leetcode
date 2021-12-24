package leetcode

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

不是最优解。
看官方题解：
https://leetcode-cn.com/problems/copy-list-with-random-pointer/solution/fu-zhi-dai-sui-ji-zhi-zhen-de-lian-biao-rblsf/

进阶： 如果链表，是带环的， 怎么办？
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var dup *Node
	dup = new(Node)
	dup.Val = head.Val
	dup.Next = nil

	if head.Random == nil {
		dup.Random = nil
	}
	if head.Random == head {
		dup.Random = dup
	}

	// recursion help us deep copy the remaining list
	dup.Next = copyRandomList(head.Next)

	h := head
	d := dup
	// we need to traverse the list and find which random pointer point to head, and we do the same for dup
	// at same time if head.Random point to some node, do the same for dup
	for h.Next != nil {
		h = h.Next
		d = d.Next

		if h.Random == head {
			d.Random = dup
		}

		if h == head.Random {
			dup.Random = d
		}
	}

	return dup
}

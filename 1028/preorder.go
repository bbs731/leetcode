package leetcode

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	if traversal == "" {
		return nil
	}
	pos := strings.IndexByte(traversal, '-')
	if pos == -1 {
		val, _ := strconv.Atoi(traversal)
		return &TreeNode{val, nil, nil}
	}

	//treat the traversal[pos:] to remove one '-' for each node's value
	// e.g. -2--3---4-5--6--7 change to 2-3--45-6-7
	sb := &strings.Builder{}
	right := -1
	cnt := 0
	for i := pos + 1; i < len(traversal); i++ {
		if traversal[i] != '-' {
			if cnt > 0 {
				// treat previous dashes
				for j := 1; j <= cnt-1; j++ {
					// 为了少写一个 dash
					sb.WriteByte('-')
				}
				if cnt == 1 {
					// we have a right child , record the position
					right = sb.Len()
				}
				// reset counter
				cnt = 0
			}
			sb.WriteByte(traversal[i])
		} else {
			cnt++
		}
	}

	val, _ := strconv.Atoi(traversal[:pos])
	root := &TreeNode{val, nil, nil}
	if right == -1 {
		// we only have left child
		root.Left = recoverFromPreorder(sb.String())
	} else {
		s := sb.String()
		root.Left = recoverFromPreorder(s[:right])
		root.Right = recoverFromPreorder(s[right:])
	}

	return root
}

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}
/*
不敢相信，如此的简单啊，
再一次证明了， 你写的像个 shit
 */
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		sb.WriteByte(',')
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")

	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}

		num, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		node := &TreeNode{num, build(), build()}
		return node
	}
	return build()
}

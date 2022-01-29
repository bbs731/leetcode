package leetcode

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}

}

func (this *Codec) serialize(root *Node) string {
	sb := &strings.Builder{}

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			sb.WriteString("null")
			return
		}
		// otherwise, we use Val * 10 + numofChildren as the serialized value
		// noc 代表 number of children
		// 如果 num of Children > 10  就有问题了。 那么我们要不然就用字符串拼接如  val-noc
		// 或者让 这个乘以的数，足够大 譬如按照提要求 total nodes < 10^4
		sb.WriteString(strconv.Itoa(node.Val*10000 + len(node.Children)))
		for i := 0; i < len(node.Children); i++ {
			sb.WriteByte(',')
			dfs(node.Children[i])
		}
	}
	dfs(root)
	return sb.String()
}

func (this *Codec) deserialize(data string) *Node {
	sp := strings.Split(data, ",")
	var build func() *Node
	build = func() *Node {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		num, _ := strconv.Atoi(sp[0])
		val := num / 10000
		n := num % 10000
		sp = sp[1:]
		children := make([]*Node, n)
		for i := 0; i < n; i++ {
			children[i] = build()
		}
		return &Node{val, children}
	}
	return build()
}

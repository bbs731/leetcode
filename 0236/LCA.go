package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
这道题， 递归就能解！ 
 */
func find(root, n *TreeNode, path []*TreeNode) (bool, []*TreeNode) {
	if root == nil {
		return false, path
	}
	path = append(path, root)
	if root == n {
		return true, path
	}
	if root.Left != nil {
		found, p := find(root.Left, n, path)
		if found {
			return true, p
		}
	}
	// must be found in right path
	return find(root.Right, n, path)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	_, p1 := find(root, p, []*TreeNode{})
	_, p2 := find(root, q, []*TreeNode{})

	ans := root
	for i := 0; i < len(p1) && i < len(p2); i++ {
		if p1[i] == p2[i] {
			ans = p1[i]
		} else {
			break
		}
	}
	return ans
}

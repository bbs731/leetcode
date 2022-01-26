package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
感觉，你退步了啊！ 是题做麻木了吗？
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}
	root := postorder[len(postorder)-1]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root {
			break
		}
	}

	return &TreeNode{
		root,
		buildTree(inorder[:i], postorder[:i]),
		buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}

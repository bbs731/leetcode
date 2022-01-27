package leetcode

/* 进阶: 递归算法很简单，你可以通过迭代算法完成吗？


这个实现有明显的错误， 因为是树， 所以无环， 那就不需要使用 visited hash table
*/

func isVisited(n *TreeNode, visited map[*TreeNode]struct{}) bool {
	if n == nil {
		return true
	}
	_, ok := visited[n]
	return ok
}

func inorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ans []int
	var stack []*TreeNode
	stack = append(stack, root)
	// 记录 TreeNode 只入队stack一次
	visited := make(map[*TreeNode]struct{})
	visited[root] = struct{}{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]

		if isVisited(top.Left, visited) {
			ans = append(ans, top.Val)
			//pop the top element
			stack = stack[:len(stack)-1 ]

			if top.Right != nil {
				stack = append(stack, top.Right)
			}
		} else {
			stack = append(stack, top.Left)
			visited[top.Left] = struct{}{}
		}
	}
	return ans
}

package leetcode

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0

	dfs(root, 0, root.Val, &max)

	return max
}

func dfs(node *TreeNode, counter int, target int, max *int) {
	if node == nil {
		return
	}

	if node.Val == target {
		counter++
	} else {
		counter = 1
	}

	*max = getMax(counter, *max)

	dfs(node.Left, counter, node.Val+1, max)
	dfs(node.Right, counter, node.Val+1, max)
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

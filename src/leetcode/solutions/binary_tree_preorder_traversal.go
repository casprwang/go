package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	if(root == nil) {
		return []
	}
}

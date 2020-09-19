package data_structure

import "fmt"

type TreeNode struct {
	Data int
	Left *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 根节点
	fmt.Print(tree.Data, " ")
	// 左子树
	PreOrder(tree.Left)
	// 右子树
	PreOrder(tree.Right)
}

// 中序遍历
func InOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 左子树
	InOrder(tree.Left)
	// 根节点
	fmt.Print(tree.Data, " ")
	// 右子树
	InOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 左子树
	PostOrder(tree.Left)
	// 右子树
	PostOrder(tree.Right)
	// 根节点
	fmt.Print(tree.Data, " ")
}
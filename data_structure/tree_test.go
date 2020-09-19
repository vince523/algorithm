package data_structure

import (
	"fmt"
	"testing"
)

func TestTreeOrder(t *testing.T) {
	tree := &TreeNode{
		Data:  1,
	}

	tree.Left = &TreeNode{
		Data: 2,
	}
	tree.Right = &TreeNode{
		Data: 3,
	}
	tree.Left.Left = &TreeNode{
		Data: 4,
	}

	tree.Left.Right = &TreeNode{
		Data: 5,
	}

	tree.Right.Left = &TreeNode{
		Data: 6,
	}

	fmt.Println("前序排序：")
	PreOrder(tree)
	fmt.Println("\n中序排序：")
	InOrder(tree)
	fmt.Println("\n后序排序：")
	PostOrder(tree)
}

package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	list := make([]int, 0)
	PreOrder(root, &list)
	assert.Equal(t, 3, len(list))
	assert.Equal(t, []int{1, 2, 3}, list)
}

func TestPreOrder2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	list := make([]int, 0)
	PreOrder2(root, &list)
	assert.Equal(t, 7, len(list))
	assert.Equal(t, []int{1, 2, 4, 5, 3, 6, 7}, list)
}

func TestInOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	list := make([]int, 0)
	InOrder(root, &list)
	assert.Equal(t, 3, len(list))
	assert.Equal(t, []int{2, 1, 3}, list)
}

func TestBackOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	list := make([]int, 0)
	BackOrder(root, &list)
	assert.Equal(t, 3, len(list))
	assert.Equal(t, []int{2, 3, 1}, list)
}

func TestLayerOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	list := make([]int, 0)
	LayerOrder(root, &list)
	assert.Equal(t, 5, len(list))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list)
}

func TestInOrder2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	list := make([]int, 0)
	InOrder2(root, &list)
	assert.Equal(t, 7, len(list))
	assert.Equal(t, []int{4, 2, 5, 1, 6, 3, 7}, list)
}

func TestBackOrder2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	list := make([]int, 0)
	BackOrder2(root, &list)
	assert.Equal(t, 7, len(list))
	assert.Equal(t, []int{4, 5, 2, 6,  7, 3, 1}, list)
}

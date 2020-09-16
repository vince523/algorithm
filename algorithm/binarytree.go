package algorithm

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 前序遍历: 先遍历根节点, 再遍历左子树, 再遍历右子树
func PreOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	PreOrder(root.Left, list)
	PreOrder(root.Right, list)
}

//PreOrder2 非递归(深度优先算法) 前序遍历
func PreOrder2(root *TreeNode, list *[]int) {
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		if root != nil {
			*list = append(*list, root.Val)
			stack = append(stack, root)
			root = root.Left
		}else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = node.Right
		}
	}
}

// InOrder 中序遍历: 先遍历左子树, 再遍历根节点, 再遍历左子树
func InOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, list)
	*list = append(*list, root.Val)
	InOrder(root.Right, list)
}

// InOrder2 中序遍历非递归
func InOrder2(root *TreeNode, list *[]int) {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		}else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			*list = append(*list, node.Val)
			root = node.Right
		}
	}
}

// BackOrder 后序遍历: 先遍历左子树, 再遍右子树, 再遍历根
func BackOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	BackOrder(root.Left, list)
	BackOrder(root.Right, list)
	*list = append(*list, root.Val)
}

// BackOrder2 后序遍历(非递归)
func BackOrder2(root *TreeNode, list *[]int) {
	stack := make([]*TreeNode, 0)
	var visit *TreeNode
	for root != nil || len(stack) !=0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) != 0 {
			root = stack[len(stack) -1]
			stack = stack[:len(stack)-1]

			if root.Right == nil || visit == root.Right {
				*list = append(*list, root.Val)
				visit = root
				root = nil
			} else {
				stack = append(stack, root)
				root = root.Right
			}
		}
	}
}


// LayerOrder 层序遍历, 一层一层遍历子节点
func LayerOrder(root *TreeNode, list *[]int) {
	queue := make([]*TreeNode, 0)
	if root == nil {
		return
	}
	queue = append(queue, root)

	var node *TreeNode
	for len(queue) != 0 {
		node = queue[0]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		*list = append(*list, node.Val)
		queue = queue[1:]
	}
}
package data_structure

import "sync"

// 链表实现的栈
type Stack struct {
	head *LinkNode // 链表起点
	size int
	mu 	 sync.Mutex
}

// 入栈
func (s *Stack) Push(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 栈顶为空，增加节点
	if s.head == nil {
		s.head = new(LinkNode)
		s.head.Val = value
	} else {
		// 把节点加到链表头部
		preNode := s.head
		newNode := new(LinkNode)
		newNode.Val = value
		newNode.Next = preNode
		s.head = newNode
	}
	s.size += 1
}

// 出栈
func (s *Stack) Pop() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 判空
	if s.IsEmpty() {
		panic("stack empty")
	}

	// 弹出链表头节点
	headNode := s.head
	result := headNode.Val

	// 新的头节点
	s.head = headNode.Next
	s.size -= 1
	return result
}

// isEmpty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() int {
	if s.size == 0 {
		panic("empty")
	}
	return s.head.Val
}
/*
单链表, 线程不安全
*/
package algorithm

import (
	"errors"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	Val  interface{}
	Next *Node
}

// Push 添加一个节点到链表尾部
func (l *LinkedList) Push(value interface{}) {
	node := &Node{
		Val:  value,
		Next: nil,
	}

	// 空的单链表
	if l.length == 0 {
		l.head = node
		l.tail = node
		// 不为空的单链表
	} else {
		n := l.tail
		n.Next = node
		l.tail = node
	}

	l.length++
}

// Pop 删除一个节点到链表尾部, 并返回数据
func (l *LinkedList) Pop() (value interface{}, err error) {
	if l.length <= 0 {
		return nil, errors.New("empty linked list")
	}

	if l.length == 1 {
		l.length--
		value = l.head.Val
		l.head, l.tail = nil, nil
		return value, nil
	}

	node := l.head
	for {
		// 倒数第二个
		if node.Next.Next == nil {
			break
		}
		node = node.Next
	}
	value = l.tail.Val
	node.Next = nil
	l.tail = node
	l.length--
	return value, nil
}

func (l *LinkedList) Peek() (value interface{}, err error) {
	if l.length == 0 {
		return nil, errors.New("empty linked list")
	}

	return l.head.Val, nil
}

// Insert 插入第几个节点之后
func (l *LinkedList) Insert(n int, value interface{}) *LinkedList {
	node := &Node{
		Val:  value,
		Next: nil,
	}
	switch {
	case n == 0:
		node.Next = l.head
		l.head = node
		l.length++
	case n > l.length:
		l.tail.Next = node
		l.tail = node
		l.length++
	default:
		a := l.head
		for i := 0; i < n; i++ {
			a = a.Next
		}
		on := a.Next
		a.Next = node
		node.Next = on
	}

	return l
}

func (l *LinkedList) Len() int {
	return l.length
}

// Reverse 反转链表, 使用一个中间变量记录即可
func (l *LinkedList) Reverse() {
	// 无需反转
	if l.length <= 1 {
		return
	}

	node := l.head
	var pre *Node
	for node != nil {
		node.Next, pre, node = pre, node, node.Next
	}
	l.head, l.tail = l.tail, l.head

}

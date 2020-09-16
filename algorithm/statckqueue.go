package algorithm

import (
	"errors"
)

type stack struct {
	items []int
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *stack) Pop() int {
	if len(s.items) == 0 {
		panic(errors.New("empty stack"))
	}
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

type CQueue struct {
	stack1 stack
	stack2 stack
}

func ConstructorCQueue() CQueue {
	return CQueue{
		stack1: stack{},
		stack2: stack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.IsEmpty() {
		// 将stack1的所有值放到stack2中
		for !this.stack1.IsEmpty() {
			this.stack2.Push(this.stack1.Pop())
		}
	}

	// 假如stack2还是为空, 说明当前为空栈
	if this.stack2.IsEmpty() {
		return -1
	}

	return this.stack2.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

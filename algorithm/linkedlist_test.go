package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	var err error
	l := &LinkedList{}
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
	assert.Equal(t, 0, l.Len())

	l.Push(1)
	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Len())
	v, err := l.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 1, v.(int))

	l.Push(2)
	assert.Equal(t, 2, l.Len())
	assert.NotEqual(t, l.tail.Val, l.head.Val)
	assert.Equal(t, 1, l.head.Val.(int))
	assert.Equal(t, 2, l.tail.Val.(int))

	l.Push(3)
	l.Push(4)
	l.Push(5)
	assert.Equal(t, 5, l.Len())
	v, err = l.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 5, v.(int))

	l.Insert(0, 0)
	assert.Equal(t, 0, l.head.Val.(int))

	l.Insert(1, 8)
	assert.Equal(t, 8, l.head.Next.Next.Val.(int))
}

func TestLinkedList_Reverse(t *testing.T) {
	l := &LinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	l.Reverse()
	val, _ := l.Pop()
	assert.Equal(t, 1, val)
	val, _ = l.Pop()
	assert.Equal(t, 2, val)
	val, _ = l.Pop()
	assert.Equal(t, 3, val)
}

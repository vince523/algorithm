package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	mStack := &MinStack{
		list: []int{},
		min:  []int{},
	}
	mStack.Push(-2)
	mStack.Push(0)
	mStack.Push(-3)

	assert.Equal(t, -3, mStack.Min())
	mStack.Pop()
	assert.Equal(t, 0, mStack.Top())
	assert.Equal(t, -2, mStack.Min())
}

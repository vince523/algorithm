package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = &Stack{}

func init() {
	s.Push("1")
	s.Push(1)
	s.Push(2)
}

func TestStack(t *testing.T) {
	assert.Equal(t, 3, s.Len())
	value, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 2, value)
}

func TestStack_Empty(t *testing.T) {
	s = &Stack{}
	value, err := s.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestStack_Peak(t *testing.T) {
	s = &Stack{}
	s.Push(1)
	assert.Equal(t, 1, s.Peak().(int))
}

func TestStack_OverTurn(t *testing.T) {
	s = &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	ns := s.OverTurn()
	var val interface{}
	val, _ = ns.Pop()
	assert.Equal(t, 1, val.(int))
	val, _ = ns.Pop()
	assert.Equal(t, 2, val.(int))
	val, _ = ns.Pop()
	assert.Equal(t, 3, val.(int))
}



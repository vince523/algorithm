package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	queue = &Queue{}
)

func init() {
	queue.Push("1")
	queue.Push(1)
	queue.Push(2)
}

func TestQueue(t *testing.T) {
	assert.Equal(t, queue.Len(), 3)
	value, err := queue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, value.(string), "1")
}

func TestQueue_Empth(t *testing.T) {
	queue = &Queue{}
	item, err := queue.Pop()
	assert.Nil(t, item)
	assert.NotNil(t, err)
}

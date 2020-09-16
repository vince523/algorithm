package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	var ok bool
	circularQueue := Constructor(3)
	ok = circularQueue.EnQueue(1)
	assert.True(t, ok)
	assert.Equal(t, 1, circularQueue.Front())
	assert.Equal(t, 1, circularQueue.Rear())
}

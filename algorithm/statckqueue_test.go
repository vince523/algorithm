package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQueue(t *testing.T) {
	stackqueue := ConstructorCQueue()
	stackqueue.AppendTail(1)
	stackqueue.AppendTail(2)
	stackqueue.AppendTail(3)
	assert.Equal(t, 1, stackqueue.DeleteHead())
	assert.Equal(t, 2, stackqueue.DeleteHead())
	assert.Equal(t, 3, stackqueue.DeleteHead())
	assert.Equal(t, -1, stackqueue.DeleteHead())
}

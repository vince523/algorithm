package data_structure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkNode(t *testing.T) {
	link := LinkNode{
		Val:  1,
		Next: nil,
	}

	link.Append(2)
	assert.Equal(t, 2, link.Get(1))
	link.Append(3)
	assert.Equal(t, 3, link.Get(2))
	link.Insert(1, 4)
	assert.Equal(t, 4, link.Get(1))
	fmt.Println(link.Get(0))
	fmt.Println(link.Get(1))
	fmt.Println(link.Get(2))
	fmt.Println(link.Get(3))
	link.Delete(2)
	assert.Equal(t, 3, link.Get(2))
}

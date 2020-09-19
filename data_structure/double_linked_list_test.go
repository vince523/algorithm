package data_structure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLink_Append(t *testing.T) {
	dl := DoubleLink{
		size: 0,
		head: nil,
		tail: nil,
	}

	dl.Append(1)
	dl.Append(2)
	dl.Append(3)

	head := dl.head
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	assert.Equal(t, 3, dl.size)
}

func TestDoubleLink_Insert(t *testing.T) {
	dl := DoubleLink{
		size: 0,
		head: nil,
		tail: nil,
	}
	dl.Append(1)
	dl.Append(2)
	dl.Append(3)

	dl.Insert(1, 11)
	head := dl.head
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	assert.Equal(t, 4, dl.size)

}

func TestDoubleLink_Remove(t *testing.T) {
	dl := DoubleLink{
		size: 0,
		head: nil,
		tail: nil,
	}
	dl.Append(1)
	dl.Append(2)
	dl.Append(3)

	dl.Insert(1, 11)
	dl.Append(5)
	head := dl.head
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println()
	dl.Remove(5)
	dl.Remove(1)
	dl.Remove(1)
	s := dl.head
	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}
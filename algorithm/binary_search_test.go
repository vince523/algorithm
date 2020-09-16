package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	value := BinarySearch(nums, 3)
	assert.Equal(t, 1, value)
}


func TestLinerSearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	value := BinarySearch(nums, 3)
	assert.Equal(t, 1, value)
}
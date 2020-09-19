package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{3, 4, 1, 6, 8, 2, 9, 7, 5}
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	after := QuickSort(arr)
	assert.Equal(t, res, after)
}

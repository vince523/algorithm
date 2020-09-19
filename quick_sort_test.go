package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 4, 1, 6, 8, 2, 9, 7, 5}
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	after := QuickSort(arr)
	assert.Equal(t, res, after)
}

func TestQuickSort2(t *testing.T) {
	arr := []int{3, 4, 1, 6, 8, 2, 9, 7, 5}
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	after := QuickSort2(arr)
	fmt.Println(arr)
	fmt.Println(after)
	assert.Equal(t, res, after)
}

func TestQuickSort3(t *testing.T) {
	arr := []int{3, 4, 1, 6, 8, 2, 9, 7, 5}
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	after := QuickSort3(arr)
	fmt.Println(arr)
	fmt.Println(after)
	assert.Equal(t, res, after)
}
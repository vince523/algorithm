package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 4, 1, 6, 8, 2, 9, 7, 5}
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 := append([]int{}, arr...)
	arr2 := append([]int{}, arr...)
	arr3 := append([]int{}, arr...)
	fmt.Println("origin arr", arr)
	assert.Equal(t, res, BubbleSort1(arr1))
	assert.Equal(t, res, BubbleSort2(arr2))
	assert.Equal(t, res, BubbleSort3(arr3))
}

package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{3, 4, 1, 6, 8, 2, 9, 7, 5}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := InsertSort(arr)
	assert.Equal(t, want, actual)
}

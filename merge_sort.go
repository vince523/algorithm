package algorithm

// 归并排序是一种分治策略的算法
// 归并排序有两种递归做法：自顶向下，自底向上

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return  arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
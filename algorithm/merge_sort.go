package algorithm

// MergeSort 归并排序
func MergeSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	// 不断拆分为left, right 切片
	middle := length / 2
	left := list[:middle]
	right := list[middle:]
	return mergeSort(MergeSort(left), MergeSort(right))
}

func mergeSort(left []int, right []int) []int {
	// 合并left 和 right 切片
	var result []int
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// 有一些大小为奇数的需要加起来
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

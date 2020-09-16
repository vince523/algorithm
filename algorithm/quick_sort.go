package algorithm

// QuickSort 快速排序
func QuickSort(list []int) []int {
	quickSort(list)
	return list
}

func quickSort(list []int) {
	if len(list) <= 1 {
		return
	}

	pivot := list[len(list)-1]
	left, right := 0, len(list)-2

	for left < right {
		// 大于基准
		switch {
		case list[left] < pivot:
			left++
		case list[right] > pivot:
			right--
		default:
			list[left], list[right] = list[right], list[left]
			left++
		}
	}
	// 交换基准
	if list[left] > list[len(list)-1] {
		list[len(list)-1], list[left] = list[left], list[len(list)-1]
	}

	// 递归分组
	quickSort(list[:left])
	quickSort(list[left+1:])
}

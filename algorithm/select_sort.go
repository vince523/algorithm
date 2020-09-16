package algorithm

// SelectSort 选择排序
func SelectSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	// 一定需要遍历n-1次
	for i := 0; i < length-1; i++ {
		// 最小值得索引
		minIndex := i
		// 同样是遍历n-1次
		for j := i + 1; j < length; j++ {
			// 交换最小值的索引
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}

		// 交换最小值
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}

	return list
}

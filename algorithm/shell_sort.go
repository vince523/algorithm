package algorithm

// ShellSort 希尔排序
func ShellSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	// 增量排序
	for gap := length / 2; gap > 0; gap /= 2 {
		// 对每个分组进行插入排序
		for i := gap; i < length; i++ {
			tmp := list[i]
			j := i - gap
			for j >= 0 && list[j] > tmp {
				list[j+gap] = list[j]
				j -= gap
			}
			list[j+gap] = tmp
		}
	}
	return list
}

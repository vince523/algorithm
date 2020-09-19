package algorithm

func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 注意这里是len(arr)-1
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIdx := i
		for j := i+1; j < len(arr); j++ {
			// 找到比原来选择的min还要小
			if arr[j] < min {
				min = arr[j]
				minIdx = j
			}
		}
		if i != minIdx {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}

	return arr
}

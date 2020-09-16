package algorithm


func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 选择一个基准
	pivot := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	// 注意这里从 1 开始
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}

	low = QuickSort(low)
	high = QuickSort(high)
	return append(append(low, pivot), high...)
}

// 非递归
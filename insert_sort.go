package algorithm

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 从左边第二位置开始
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		// 将值比tmp大的右移
		for ; j >= 0 && arr[j] >= tmp; j-- {
			arr[j+1] = arr[j]
		}
		// 由于上面j--
		arr[j+1] = tmp
	}

	return arr
}
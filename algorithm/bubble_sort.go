package algorithm

// BubbleSort 冒泡排序
func BubbleSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	// 重点不在这里的i, 这里的i只是为了构建次数
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if list[j] > list[j+1] {
				// 两两交换, 最后一个为最大值
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	return list
}

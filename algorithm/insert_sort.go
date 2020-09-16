package algorithm

// InsertSort 插入排序
func InsertSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	// 从索引为1的开始
	for i := 1; i < length; i++ {
		// 定义索引和临时变量
		index := i
		tmp := list[index]

		// 将值比临时变量大的左移
		for index >= 0 && tmp < list[index-1] {
			list[index] = list[index-1]
			index--
		}

		// 赋值(因为是平移不是交换)
		list[index] = tmp
	}

	return list
}

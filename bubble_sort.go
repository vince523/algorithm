package algorithm

// 基础版本 O(n^2)
func BubbleSort1(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func BubbleSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		// i 左边已排序
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func BubbleSort3(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		swag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swag = true
			}
		}

		if !swag {
			break
		}
	}
	return arr
}


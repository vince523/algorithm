package algorithm

func ShellSort(arr []int) []int {
	length := len(arr)

	// gap
	gap := 1
	for gap < length / 3 {
		gap = gap * 3 - 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			tmp := arr[i]
			j := i - gap
			for ; j >= 0 && arr[j] > tmp; j = j-gap {
				arr[j+gap] = arr[j]
			}

			arr[j+gap] = tmp
		}
		gap = gap / 3
	}
	return arr
}
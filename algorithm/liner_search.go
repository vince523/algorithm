package algorithm

func LinerSearch(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
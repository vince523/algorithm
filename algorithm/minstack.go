package algorithm

type MinStack struct {
	list []int
	min  []int
}

/** initialize your data structure here. */

func (s *MinStack) Push(x int) {
	s.list = append(s.list, x)

	// 插入min栈
	if len(s.min) == 0 || x <= s.Min() {
		s.min = append(s.min, x)
	}
}

func (s *MinStack) Pop() int {
	value := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	if value == s.Min() {
		s.min = s.min[:len(s.min)-1]
	}
	return value
}

func (s *MinStack) Top() int {
	return s.list[len(s.list)-1]
}

func (s *MinStack) Min() int {
	return s.min[len(s.min)-1]
}

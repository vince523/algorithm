package algorithm

import "sync"

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

// 三向切分
func QuickSort2(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	quickSort2(res, 0, len(arr)-1)
	return res
}

func quickSort2(arr []int, begin, end int) {
	if begin < end {
		lt, gt := partition(arr, begin, end)
		// 左边三向快排
		quickSort2(arr, begin, lt-1)
		// 右边三向快排
		quickSort2(arr, gt+1, end)
	}
}

// 切分函数
func partition(arr []int, begin, end int) (int, int) {
	left := begin
	right := end
	i := begin + 1
	pivot := arr[begin]

	for i <= right {
		// 大于基准
		if arr[i] > pivot {
			arr[i], arr[right] = arr[right], arr[i]
			// 右指针左移
			right--
		} else if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			// 左指针右移
			left++
			// i 指针右移
			i++
		} else {
			// 不需交换，i 指针右移
			i++
		}
	}
	return left, right
}

// 非递归
type LinkNode struct {
	Val int
	Next *LinkNode
}

type Stack struct {
	head *LinkNode // 链表起点
	size int
	mu 	 sync.Mutex
}

// 入栈
func (s *Stack) Push(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 栈顶为空，增加节点
	if s.head == nil {
		s.head = new(LinkNode)
		s.head.Val = value
	} else {
		// 把节点加到链表头部
		preNode := s.head
		newNode := new(LinkNode)
		newNode.Val = value
		newNode.Next = preNode
		s.head = newNode
	}
	s.size += 1
}

// 出栈
func (s *Stack) Pop() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 判空
	if s.IsEmpty() {
		panic("stack empty")
	}

	// 弹出链表头节点
	headNode := s.head
	result := headNode.Val

	// 新的头节点
	s.head = headNode.Next
	s.size -= 1
	return result
}

// isEmpty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func QuickSort3(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	quickSort3(res)
	return res
}

func quickSort3(arr []int) {
	// 手动维护栈
	stack := new(Stack)

	stack.Push(len(arr)-1)
	stack.Push(0)

	// 栈非空， 证明还有未排序
	for !stack.IsEmpty() {
		begin := stack.Pop()
		end := stack.Pop()

		// 切分
		loc := partition3(arr, begin, end)

		/*
		 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[loc-1] < array[begin]
		 *                      -->  array[loc+1] ~ array[end] > array[begin]
		*/
		// 右边入栈
		if loc + 1 < end {
			stack.Push(end)
			stack.Push(loc+1)
		}

		// 左边入栈
		if begin < loc {
			stack.Push(loc-1)
			stack.Push(begin)
		}
	}

}

// 非递归的切分函数, 返回切分元素的下标
func partition3(arr []int, begin, end int) int {
	i := begin+1
	j := end
	pivot := arr[begin]

	// 没重合之前
	for i < j {
		if arr[i] > pivot {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}

	/* i == j 重合了
	 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
	*                       -->  array[i+1] ~ array[end] > array[begin]
	需要判断arr[i] 和 arr[begin] 决定 arr[i] 的位置
	 */
	//  这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
	if arr[i] >= arr[begin] {
		i--
	}
	arr[begin], arr[i] = arr[i], arr[begin]
	return i
}
package algorithm

func HeapSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	heapSort(res)
	return res
}

func heapSort(arr []int) []int {
	h := NewHeap(arr)
	for _, v := range arr {
		h.Push(v)
	}
	for range arr {
		h.Pop()
	}
	return arr
}

type Heap struct {
	Size int
	// 使用数组模拟树
	// 一个节点下标为i, 那么父亲节点下标为 (i-1)/2
	// 一个节点下标为i, 那么左儿子的下标为 2i+1, 右儿子下标为 2i+2k
	Array []int
}

func NewHeap(arr []int) *Heap {
	h := new(Heap)
	h.Array = arr
	return h
}

func (h *Heap) Push(value int) {
	// 堆没有元素，那把元素设置为顶点元素
	if h.Size == 0 {
		h.Array[0] = value
		h.Size++
		return
	}

	// i 是要插入的下标
	i := h.Size

	// 下标存在
	for i > 0 {
		// 父节点下标
		parent := (i-1) / 2

		// 如果插入的值小于等于父节点，退出
		if value <= h.Array[parent] {
			break
		}

		// 否则，交换
		h.Array[i] = h.Array[parent]
		i = parent
	}

	// 把该元素放到最终 i 的位置
	h.Array[i] = value
	h.Size++
}

// 最大堆删除节点元素
func (h *Heap) Pop() int {
	if h.Size == 0 {
		return -1
	}

	// 根节点 最大的元素
	ret := h.Array[0]
	h.Size--
	// 根节点删除后，将最后一个节点放到根节点上，进行比较，维持堆的特征
	last := h.Array[h.Size]
	h.Array[h.Size] = ret // 将要移除的元素放到最后

	i := 0
	for {
		left := 2 * i + 1
		right := 2 * i + 2

		// 左子树下标超出
		if left >= h.Size {
			break
		}

		// 有右子树, 取二者中较大的下标
		if right < h.Size && h.Array[right] > h.Array[left] {
			left = right
		}

		// 待定的父节点的值大于等于左右节点的较大者时，就可以退出循环了
		if last >= h.Array[left] {
			break
		}

		// 将较大的儿子与父亲交换
		h.Array[i] = h.Array[left]

		i = left
	}

	// 将最后一个元素放在不会再翻转的位置
	h.Array[i] = last
	return ret

}
package data_structure

import "sync"

// 数组队列
type ArrayQueue struct {
	arr []int
	size int
	lock sync.Mutex
}

// 入队
func (q *ArrayQueue) Enqueue(val int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.arr = append(q.arr, val)
	q.size++
}

// 出队
func (q *ArrayQueue) Dequeue() int{
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		panic("empty")
	}

	target := q.arr[0]
	tmp := make([]int, q.size-1, q.size-1)
	for i := 1; i < q.size; i++ {
		tmp[i-1] = q.arr[i]
	}

	q.arr = tmp
	q.size--
	return target
}

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// 入队
func (q *LinkQueue) Add(val int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.root == nil {
		q.root = new(LinkNode)
		q.root.Val = val
	} else {
		newNode := new(LinkNode)
		newNode.Val = val

		head := q.root
		for head.Next != nil {
			head = head.Next
		}

		head.Next = newNode
	}
	q.size++
}

// 出队
func (q *LinkQueue) Remove() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		panic("empty")
	}

	head := q.root
	res := head.Val

	q.root  = head.Next
	q.size--
	return res
}
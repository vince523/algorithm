package algorithm

// golang 循环队列实现
type MyCircularQueue struct {
	data []int
	front, rear int  // front标识出队元素的下标，rear标识入队的下标
	maxSize int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func CircleQueueConstructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data: make([]int, k + 1, k + 1),
		front: 0,
		rear: 0,
		maxSize: k + 1, // 多加一个用于给rear指针站位
	}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (c *MyCircularQueue) EnQueue(value int) bool {
	if c.IsFull() {
		return false
	}
	c.data[c.rear] = value
	c.rear++
	if c.rear == c.maxSize {
		c.rear = 0
	}

	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (c *MyCircularQueue) DeQueue() bool {
	if c.IsEmpty() {
		return false
	}
	c.front++
	if c.front == c.maxSize {
		c.front = 0
	}

	return true
}


/** Get the front item from the queue. */
func (c *MyCircularQueue) Front() int {
	if c.IsEmpty() {
		return -1
	}
	return c.data[c.front]
}


/** Get the last item from the queue. */
func (c *MyCircularQueue) Rear() int {
	if c.IsEmpty() {
		return -1
	}
	lastIndex := c.rear - 1
	// 注意边界
	if lastIndex < 0 {
		lastIndex = c.maxSize - 1
	}
	return c.data[lastIndex]
}


/** Checks whether the circular queue is empty or not. */
func (c *MyCircularQueue) IsEmpty() bool {
	return c.front == c.rear
}


/** Checks whether the circular queue is full or not. */
func (c *MyCircularQueue) IsFull() bool {
	next := c.rear + 1
	if next == c.maxSize {
		next = 0
	}
	return next == c.front
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

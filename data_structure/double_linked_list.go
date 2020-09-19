package data_structure

type DoubleLinkNode struct {
	Val int
	Prev *DoubleLinkNode
	Next *DoubleLinkNode
}
type DoubleLink struct {
	size int
	head *DoubleLinkNode
	tail *DoubleLinkNode
}

func (d *DoubleLink) Append(val int) {
	node := &DoubleLinkNode{
		Val:  val,
		Prev: nil,
		Next: nil,
	}

	if d.size == 0 {
		d.head = node
		d.tail = node
	} else {
		node.Prev = d.tail
		node.Next = nil
		d.tail.Next = node
		d.tail = node
	}

	d.size++
}

func (d *DoubleLink) Insert(i int, val int) {
	head := d.head
	j := 1
	for head != nil && j < i {
		head = head.Next
		j++
	}

	if head == d.tail {
		d.Append(val)
		return
	}

	if head == nil || j > i {
		return
	}

	node := &DoubleLinkNode{
		Val:  val,
		Prev: head,
		Next: head.Next,
	}

	head.Next = node
	node.Next.Prev = node
	d.size++
}

func (d *DoubleLink) Remove(i int) {
	head := d.head
	j := 1
	for head != nil && j < i {
		head = head.Next
		j++
	}
	if head == nil || j > i {
		return
	}
	prev := head.Prev
	pNext := head.Next

	if head == d.head {
		d.head = head.Next
	} else {
		prev.Next = pNext
	}

	if head == d.tail {
		d.tail = prev
	} else {
		pNext.Prev = prev
	}

	d.size--
}
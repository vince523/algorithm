package data_structure

type LinkNode struct {
	Val int
	Next *LinkNode
}

func (l *LinkNode) Insert( i int, val int) bool{
	p := l
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}

	if p == nil || j > i{
		return false
	}
	newNode := &LinkNode{
		Val:  val,
		Next: nil,
	}
	newNode.Next = p.Next
	p.Next = newNode
	return true
}


func (l *LinkNode) Append(val int) bool {

	head := l
	for head != nil {
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	node := &LinkNode{
		Val:  val,
		Next: nil,
	}
	if head == nil {
		return false
	}
	head.Next = node
	return true
}

func (l *LinkNode)Delete( i int) bool {
	head := l
	j := 1
	for head != nil && j < i{
		head = head.Next
		j++
	}
	if head == nil || j > i{
		return false
	}

	head.Next = head.Next.Next
	return true
}

func (l *LinkNode) Get(i int) int {
	head := l
	for j :=0; j < i; j++ {
		if head == nil {
			panic("err")
		}
		head = head.Next
	}

	return head.Val
}
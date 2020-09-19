package data_structure

type Ring struct {
	next, prev *Ring
	Value interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}

	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// 获取第n个节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}

	switch {
	// 往左边遍历
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
		// 往右边遍历
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 添加节点, 返回之前节点的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点后的n个节点
func (r *Ring) UnLink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n+1))
}

func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
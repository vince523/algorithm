package algorithm

/*
lru 算法实现
核心思想：
    1. 新数据插入到链表头部；
    2. 每当缓存命中（即缓存数据被访问），则将数据移到链表头部；
    3. 当链表满的时候，将链表尾部的数据丢弃。
 */

type LRUCache struct {
	cap int
	m map[int]*ListNode
	head, tail *ListNode
}

// 双向链表
type ListNode struct {
	key, val int
	prev, next *ListNode
}


func Constructor(capacity int) LRUCache {
	head := &ListNode{
		key:  0,
		val:  0,
		prev: nil,
		next: nil,
	}
	tail := &ListNode{
		key:  0,
		val:  0,
		prev: nil,
		next: nil,
	}
	// 构成双向链表
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:   capacity,
		m: make(map[int]*ListNode),
		head:  head,
		tail:  tail,
	}
}


func (l *LRUCache) Get(key int) int {
	cache := l.m
	// try get
	if node, ok := cache[key]; ok {
		// 移到头部
		l.MoveToHead(node)
		return node.val
	}
	return -1
}


func (l *LRUCache) Put(key int, value int)  {
	tail := l.tail
	cache := l.m
	if node, ok := cache[key]; ok {
		node.val = value
		l.MoveToHead(node)
		return
	}

	// 满载
	if len(cache) == l.cap {
		delete(cache, tail.prev.key)
		l.RemoveNode(tail.prev)
	}

	newNode := &ListNode{
		key:  key,
		val:  value,
		prev: nil,
		next: nil,
	}
	l.AddNode(newNode)
	l.m[key] = newNode
}

// 把节点加到头部
func (l *LRUCache) AddNode(node *ListNode) {
	head := l.head
	node.next = head.next
	node.next.prev = node
	node.prev = head
	head.next = node
}

// 删除当前节点
func (l *LRUCache) RemoveNode(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) MoveToHead(node *ListNode) {
	l.RemoveNode(node)
	l.AddNode(node)
}


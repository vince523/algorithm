/*
LRU 算法
*/
package algorithm

type LRUCache struct {
	cache map[int]*DLNode // 存储
	cap int // capacity
	head, tail *DLNode
}

type DLNode struct {
	key,value int
	prev, next *DLNode
}

func NewLRUCache(cap int) *LRUCache {
	head := &DLNode{-1, -1, nil, nil}
	tail := &DLNode{-1, -1, nil, nil}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		cache: make(map[int]*DLNode),
		cap: cap,
		head: head,
		tail: tail,
	}
}

func (lru *LRUCache) addNode(node *DLNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next = node
	node.next.prev = node
}

func (lru *LRUCache) removeNode(node *DLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *DLNode) {
	lru.removeNode(node)
	lru.addNode(node)
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key];ok {
		node.value = value
		lru.moveToHead(node)
	} else {
		n := &DLNode{key, value, nil, nil}
		if len(lru.cache) >= lru.cap {
			delete(lru.cache, lru.tail.prev.key)
			lru.removeNode(lru.tail.prev)
		}
		lru.cache[key] = n
		lru.addNode(n)
	}
}

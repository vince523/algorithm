package algorithm

import (
	"errors"
	"sync"
)

type Queue struct {
	items []interface{}
	sync.RWMutex
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Len() int {
	q.RLock()
	defer q.RUnlock()
	return len(q.items)
}

func (q *Queue) Pop() (item interface{}, err error) {
	q.RLock()
	defer q.RUnlock()
	if len(q.items) == 0 {
		return nil, errors.New("pop from empty")
	}

	item = q.items[0]
	q.items = q.items[1:]
	return item, nil
}

package lru

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data int
	Ptr  *list.Element
}

type LRUCache struct {
	Capacity int
	Queue    *list.List
	Items    map[int]*Node
}

func New(capacity int) *LRUCache {
	cache := LRUCache{
		Queue:    list.New(),
		Capacity: capacity,
		Items:    make(map[int]*Node),
	}

	return &cache
}

func (l *LRUCache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.Ptr)
		return item.Data
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {

	if item, ok := l.Items[key]; ok {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.Ptr)
	} else {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}

		ptr := l.Queue.PushFront(key)
		item := &Node{Data: value, Ptr: ptr}
		l.Items[key] = item
	}

}

func RunLRU() {

	fmt.Println("Test case 1")
	obj := New(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

	// Test case 2
	// 	["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]
	fmt.Println("Test case 2")
	obj = New(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))

}

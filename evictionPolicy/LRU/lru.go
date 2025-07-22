package LRU

import "errors"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type Cache struct {
	store    map[int]*Node
	capacity int
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *Cache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &Cache{
		store:    map[int]*Node{},
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (lru *Cache) Get(key int) (int, error) {
	if node, found := lru.store[key]; found {
		lru.deleteNode(node)
		lru.moveToFront(node)
		return node.value, nil
	} else {
		return -1, errors.New("key not found")
	}
}

func (lru *Cache) Set(key, value int) error {
	if lru.capacity < 1 {
		return errors.New("invalid capacity")
	}
	if node, found := lru.store[key]; found {
		lru.deleteNode(node)
	} else {
		if len(lru.store) >= lru.capacity {
			delete(lru.store, lru.tail.prev.key)
			lru.deleteNode(lru.tail.prev)
		}
	}
	tmp := &Node{key: key, value: value}
	lru.moveToFront(tmp)
	lru.store[key] = tmp
	return nil
}

func (lru *Cache) moveToFront(node *Node) {
	oldFirst := lru.head.next
	lru.head.next = node
	node.next = oldFirst
	oldFirst.prev = node
	node.prev = lru.head
}

func (lru *Cache) deleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

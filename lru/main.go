package main

type LRUCache struct {
	capacity int
	cache    map[int]*node
	list     list
}

type node struct {
	prev *node
	next *node
	key  int
	val  int
}

type list struct {
	head *node
	tail *node
}

func (this *LRUCache) insert(n *node) {
	if this.list.tail == nil {
		this.list.tail = n
	}

	if this.list.head == nil {
		this.list.head = n
		n.next = this.list.tail

		return
	}

	this.list.head.prev, n.next, this.list.head = n, this.list.head, n
}

func (this *LRUCache) popTail() {
	delete(this.cache, this.list.tail.key)

	if this.list.tail.prev == nil {
		this.list.tail = nil
		return
	}

	this.list.tail.prev.next = nil
	this.list.tail = this.list.tail.prev
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, list: list{nil, nil}, cache: map[int]*node{}}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	if this.list.head != nil && node.key == this.list.head.key {
		return node.val
	}

	if this.list.tail != nil && node.key == this.list.tail.key {
		this.popTail()
		this.insert(node)
		this.cache[node.key] = node

		return node.val
	}

	if node.prev != nil && node.next != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
		this.insert(node)
	}

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	// we let Get do the heavy lifting of moving an existing key to the front
	if this.Get(key) != -1 {
		this.cache[key].val = value
		return
	}

	n := &node{key: key, val: value}

	this.cache[key] = n
	this.insert(n)

	if len(this.cache) > this.capacity {
		this.popTail()
	}

	return
}

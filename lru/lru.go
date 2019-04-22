package lru_cache

type LRUCache struct {
	capacity int
	cache    map[int]*node
	list     list
}

type list struct {
	head *node
	tail *node
}

type node struct {
	prev *node
	next *node
	key  int
	val  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*node{}, list{nil, nil}}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}

	// if the node is head of the list, return immediately
	if node.key == l.list.head.key {
		return node.val
	}

	if node.next != nil && node.prev != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	} else if node.prev != nil {
		node.prev.next = node.next
	}

	node.next = l.list.head
	l.list.head.prev = node
	l.list.head = node

	if l.list.tail.key == node.key {
		l.list.tail = node.prev
	}

	return node.val
}

// why they don't like the single type for two params idk
func (l *LRUCache) Put(key int, value int) {
	// we let the Get function do the work on moving an already existing key on insert
	if l.Get(key) != -1 {
		l.cache[key].val = value
		return
	}

	node := &node{prev: nil, next: l.list.head, key: key, val: value}

	if l.list.tail == nil {
		l.list.tail = node
	}

	if l.list.head != nil {
		l.list.head.prev = node
	}

	l.list.head = node
	l.cache[key] = node

	// pop tail off if needed
	if len(l.cache) > l.capacity {
		delete(l.cache, l.list.tail.key)

		if l.list.tail.prev != nil {
			l.list.tail.prev.next = nil
		}

		l.list.tail = l.list.tail.prev
	}

	return
}

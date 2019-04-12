package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]*node
	list     list
}

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type list struct {
	head *node
	tail *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*node{}, list{}}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// if node is head of list, do nothing
	if this.list.head != nil && this.list.head.key == node.key {
		return node.val
	}

	// if node is tail
	if this.list.tail != nil && this.list.tail.key == node.key {
		// null prev's next
		if node.prev != nil {
			node.prev.next = nil
			this.list.tail = node.prev
		}

		node.prev = nil

		// set head
		node.next = this.list.head
		this.list.head = node

		return node.val
	}

	// at this point, the node is somewhere in the middle, join the sides
	if node.prev != nil && node.prev.next != nil {
		node.prev.next = node.next
	}

	if node.next != nil && node.next.prev != nil {
		node.next.prev = node.prev
	}
	node.prev = nil

	// set head
	node.next = this.list.head
	this.list.head = node

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	// we'll let the get function do heavy lifting with shifting things around on attempting
	// to set a key that already exists
	if this.Get(key) > 0 {
		this.cache[key].val = value

		return
	}

	// create node, add to cache, and set head (if needed also set as tail)
	node := &node{key, value, nil, this.list.head}
	this.cache[key] = node

	this.list.head = node

	if this.list.tail == nil {
		this.list.tail = node
	}

	// if cache length is over maximum, we need to remove the final node
	if len(this.cache) <= this.capacity {
		return
	}

	delete(this.cache, this.list.tail.key)

	if this.list.tail.prev != nil {
		this.list.tail.prev.next = nil
	}

	this.list.tail = this.list.tail.prev

	return
}

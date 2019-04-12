package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	list     list
}

type list struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
	key  int
	val  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*Node{}, list{}}
}

func (this *LRUCache) Put(key int, value int) {
	// check to see if key exists, set new value and return if it does
	node, ok := this.cache[key]
	if ok {
		node.val = value

		// top of list? do nothing
		if this.list.head.key == node.key {
			return
		}

		// if at tail, shift
		if this.list.tail != nil && this.list.tail.key == key {
			// set prev as tail
			this.list.tail = node.prev
			this.list.tail.next = nil

			// set new head as node, change old to point to this node
			node.next = this.list.head
			this.list.head.prev = node
			this.list.head = node
			node.prev = nil

			return
		}

		return
	}

	// create a new node and insert at top of list
	node = &Node{next: this.list.head, prev: nil, key: key, val: value}

	if this.list.head != nil {
		this.list.head.prev = node
		node.next = this.list.head
	}

	if this.list.tail == nil {
		this.list.tail = node
	}

	this.list.head = node
	this.cache[key] = node

	// pop off if needed
	if len(this.cache) > this.capacity {
		if this.list.tail != nil {
			delete(this.cache, this.list.tail.key)
		}
		this.list.tail = this.list.tail.prev
		this.list.tail.next = nil

	}

	return
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// top of list? do nothing
	if this.list.head.key == node.key {
		return node.val
	}

	// if at tail, shift
	if this.list.tail != nil && this.list.tail.key == key {
		// set prev as tail
		this.list.tail = node.prev
		this.list.tail.next = nil

		// set new head as node, change old to point to this node
		node.next = this.list.head
		this.list.head.prev = node
		this.list.head = node
		node.prev = nil

		return node.val
	}

	// if we're here, we know we're in the middle, join the sides
	node.prev.next = node.next
	node.next.prev = node.prev

	// change the head's prev to non-nil then change head of list to current node
	this.list.head.prev = node
	node.prev = nil
	this.list.head = node

	return node.val
}

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
	// check to see if key exists
	node, ok := this.cache[key]
	if ok {
		node.val = value

		/*
			// if already at head, do nothing
			if this.list.head != nil && this.list.head.key == key {
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

			// if we're here, we know we're in the middle, join the sides
			node.prev.next = node.next
			node.next.prev = node.prev

			// change the head's prev to non-nil then change head of list to current node
			this.list.head.prev = node
			node.prev = nil
			this.list.head = node

		*/
		return
	}

	// create and insert new node at top of list and into cache map
	node = &Node{next: this.list.head, prev: nil, key: key, val: value}
	if this.list.head != nil {
		this.list.head.prev = node
		node.next = this.list.head
	}

	this.list.head = node
	this.cache[key] = node

	if this.list.tail == nil {
		this.list.tail = node
	}

	// pop off if needed
	fmt.Println("\n BEFORE POP")
	fmt.Printf("%+v \n", this)
	fmt.Println(len(this.cache))
	if len(this.cache) > this.capacity {
		if this.list.tail != nil {
			delete(this.cache, this.list.tail.key)
		}
		this.list.tail = this.list.tail.prev
		this.list.tail.next = nil
		fmt.Println("AFTER POP \n")
		fmt.Printf("%+v \n", this)
		fmt.Println(len(this.cache))
	}

	return
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("\n GET")
	fmt.Printf("%+v \n", this)
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// shift list first then we can figure out if we need to pop

	// join the sides
	if node.key == this.list.head.key {
		return node.val
	}

	node.prev.next = node.next

	if node.next != nil {
		node.next.prev = node.prev
	}

	// change the head's prev to non-nil then change head of list to current node
	this.list.head.prev = node
	node.prev = nil
	this.list.head = node

	return node.val
}


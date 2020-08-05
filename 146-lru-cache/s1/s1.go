package main

import "fmt"

type LRUCache struct {
	items    map[int]*Node
	head     *Node
	tail     *Node
	capacity int
	size     int
}

type Node struct {
	pre   *Node
	next  *Node
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		items:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.items[key]; ok {
		value := node.value
		this.moveToTheTail(node)
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.items[key]; ok {
		node.value = value
		this.moveToTheTail(node)
		return
	}
	node := &Node{key: key, value: value}
	if this.size >= this.capacity {
		this.deleteFirst()
	}
	this.addToTheTail(node)
}

func (this *LRUCache) moveToTheTail(node *Node) {
	// 1. cut node
	node.pre.next = node.next
	node.next.pre = node.pre

	// 2. add to tail
	node.pre = this.tail.pre
	node.next = this.tail
	this.tail.pre.next = node
	this.tail.pre = node
}
func (this *LRUCache) addToTheTail(node *Node) {
	this.size++
	node.pre = this.tail.pre
	node.next = this.tail
	this.tail.pre.next = node
	this.tail.pre = node
	this.items[node.key] = node
}
func (this *LRUCache) deleteFirst() {
	this.size--
	delete(this.items, this.head.next.key)
	this.head.next.next.pre = this.head
	this.head.next = this.head.next.next
}

func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	_ = lru.Get(4)
	_ = lru.Get(3)
	//for p := lru.head.next; p != lru.tail; p = p.next {
	//	fmt.Printf("%+v\n", p)
	//}
	//lru.Get(2)
	//for p := lru.head.next; p != lru.tail; p = p.next {
	//	fmt.Printf("%+v\n", p)
	//}
	r := lru.Get(1)
	lru.Put(5, 5)
	r = lru.Get(1)
	r = lru.Get(2)
	r = lru.Get(3)
	r = lru.Get(4)
	r = lru.Get(5)
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", lru)
}

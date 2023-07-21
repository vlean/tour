package main

import "fmt"

type Node struct {
	key   string
	value int
	prev  *Node
	next  *Node
}
type LRUCache struct {
	cap   int
	nodes map[string]*Node
	head  *Node
	tail  *Node
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		cap:   cap,
		nodes: make(map[string]*Node, cap),
		head:  nil,
		tail:  nil,
	}
}

func (l *LRUCache) removeTail() {
	if l.tail == nil {
		return
	}

	delete(l.nodes, l.tail.key)

	if l.tail == l.head {
		l.tail = nil
		l.head = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
}

func (l *LRUCache) pushHead(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}

	l.nodes[node.key] = node
	if l.cap < len(l.nodes) {
		l.removeTail()
	}
}

func (l *LRUCache) Put(key string, value int) {
	// 节点存在
	if node, ok := l.nodes[key]; ok {
		node.value = value
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
		node.next = l.head
		node.prev = nil
		l.head = node
		return
	}
	l.pushHead(&Node{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	})
}

func (l *LRUCache) Get(key string) (val int, ok bool) {
	node, ok := l.nodes[key]
	if ok {
		val = node.value
	}
	return
}

func (l *LRUCache) Print() {
	node := l.head
	fmt.Printf("cap:%2d %2d\t", l.cap, len(l.nodes))
	for {
		if node == nil {
			fmt.Print("nil\n")
			break
		}
		fmt.Printf("%s:%d\t", node.key, node.value)
		node = node.next
	}
}

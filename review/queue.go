package main

type Node2 struct {
	data interface{}
	next *Node2
}

type LinkedList struct {
	head *Node2
	tail *Node2
}

// Enqueue 入队
func (l *LinkedList) Enqueue(data interface{}) {
	node := &Node2{data, nil}
	if l.head == nil {
		l.head = node
		//l.tail = node
	} else {
		// 也会改动l.head
		l.tail.next = node
		l.tail = node
	}
}

// Dequeue 出队
func (l *LinkedList) Dequeue() interface{} {
	if l.head == nil {
		return nil
	}
	tmp := l.head
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return tmp.data
}

// 先进先出

type Node3 struct {
	value interface{}
	next  *Node3
}

type Queue struct {
	head *Node3
	tail *Node3
}

func (q *Queue) Enqueue(data interface{}) {
	node := &Node3{value: data}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

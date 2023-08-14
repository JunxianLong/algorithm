package main

import "sync"

// Stack 栈
type Stack struct {
	top    *Node
	length int
	lock   sync.RWMutex
}

// Node 节点
type Node struct {
	value interface{}
	prev  *Node
}

// NewStack Create a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Len 长度
func (s *Stack) Len() int {
	return s.length
}

// Peek 查看顶部元素
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	n := &Node{value: v, prev: s.top}
	s.top = n
	s.length++
}

// isValid 是否有效括号
func isValid(s string) bool {
	stack := NewStack()
	brackets := map[uint8]uint8{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := range s {
		if stack.Len() > 0 {
			v := stack.Peek()
			if b, ok := brackets[s[i]]; ok && b == v.(uint8) {
				stack.Pop()
				continue
			}
		}
		stack.Push(s[i])
	}
	return stack.Len() == 0
}

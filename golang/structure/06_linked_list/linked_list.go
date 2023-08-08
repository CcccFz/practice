package _6_linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	next *Node
	val  interface{}
}

func NewNode(val interface{}, next *Node) *Node {
	return &Node{val:  val, next: next,}
}

type LinkedList struct {
	head   *Node
	length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: NewNode(nil, nil),
		length: 0,
	}
}

func (l *LinkedList) Len() uint {
	return l.length
}

// O(1)
func (l *LinkedList) insertAfter(p *Node, val interface{}) bool {
	if p == nil {
		return false
	}

	tmp := p.next
	p.next = NewNode(val, tmp)
	l.length += 1

	return  true
}

// O(N)
func (l *LinkedList) insertBefore(p *Node, val interface{}) bool {
	if p == nil || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head

	for cur.next != p {
		if cur == p {
			break
		}

		pre = cur
		cur = cur.next
	}

	pre.next = NewNode(val, cur)
	l.length += 1
	return true
}

// O(1)
func (l *LinkedList) InsertToHead(val interface{}) bool {
	return l.insertAfter(l.head, val)
}

// O(N)
func (l *LinkedList) InsertToTail(val interface{}) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	return l.insertAfter(cur, val)
}

// O(N)
func (l *LinkedList) FindByIdx(idx uint) *Node {
	if idx >= l.length {
		return nil
	}

	cur := l.head
	for i := 0; uint(i) <= idx; i ++ {
		cur = cur.next
	}

	return cur
}

// O(N)
func (l *LinkedList) DeleteNode(p *Node) bool {
	if p == nil || p == l.head {
		return false
	}

	cur := l.head

	for cur.next != p {
		if cur.next == p {
			break
		}
		cur = cur.next
	}

	cur.next = p.next
	l.length -= 1
	return true
}

func (l *LinkedList) String() string {
	s := make([]string, 0)
	cur := l.head

	for cur.next != nil {
		cur = cur.next
		s = append(s, fmt.Sprint(cur.val))
	}

	return strings.Join(s, ", ")
}
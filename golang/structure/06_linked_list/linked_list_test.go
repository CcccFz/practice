package _6_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	a := assert.New(t)

	l1, l2 := NewLinkedList(), NewLinkedList()
	a.NotNil(l1)
	a.NotNil(l2)

	for i := 0; i < 5; i++ {
		l1.InsertToTail(i+1)
	}
	for i := 5; i > 0; i-- {
		l2.InsertToHead(i)
	}
	a.Equal(uint(5), l1.Len())
	a.Equal(uint(5), l2.Len())
	a.Equal("1, 2, 3, 4, 5", l1.String())
	a.Equal("1, 2, 3, 4, 5", l2.String())
}

func TestFind(t *testing.T) {
	a := assert.New(t)

	l := NewLinkedList()
	a.NotNil(l)

	for i := 0; i < 5; i++ {
		l.InsertToTail(i+1)
	}

	a.Equal(1, l.FindByIdx(0).val)
	a.Equal(5, l.FindByIdx(4).val)
	a.Nil(l.FindByIdx(5))
	a.Nil(l.FindByIdx(6))
}

func TestDelete(t *testing.T) {
	a := assert.New(t)

	l := NewLinkedList()
	a.NotNil(l)

	for i := 0; i < 5; i++ {
		l.InsertToTail(i+1)
	}

	a.Equal(false, l.DeleteNode(l.head))

	a.Equal(true, l.DeleteNode(l.head.next))
	a.Equal("2, 3, 4, 5", l.String())
	a.Equal(uint(4), l.Len())

	a.Equal(true, l.DeleteNode(l.head.next.next.next.next))
	a.Equal("2, 3, 4", l.String())
	a.Equal(uint(3), l.Len())

	a.Equal(false, l.DeleteNode(l.head.next.next.next.next))
}
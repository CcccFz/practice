package _6_linked_list

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LinkedListAlgoTestSuite struct {
	suite.Suite
	l *LinkedList
}

func TestLinkedListAlgoTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListAlgoTestSuite))
}

func (s *LinkedListAlgoTestSuite) SetupTest() {
	s.l = NewLinkedList()

	for i := 0; i < 5; i++ {
		s.l.InsertToTail(i+1)
	}
}

func (s *LinkedListAlgoTestSuite) TestReverse() {
	s.l.Reverse()
	s.Equal(uint(5), s.l.Len())
	s.Equal("5, 4, 3, 2, 1", s.l.String())
}

func (s *LinkedListAlgoTestSuite) TestMerge() {
	l1, l2 := NewLinkedList(), NewLinkedList()

	for i := 1; i <= 5; i++ {
		l1.InsertToTail(2 * i - 1)
	}

	for i := 1; i <= 5; i++ {
		l2.InsertToTail(2 * i)
	}

	s.Equal("1, 2, 3, 4, 5, 6, 7, 8, 9, 10", Merge(l1, l2).String())
}

func (s *LinkedListAlgoTestSuite) TestDeleteBottomN() {
	s.l.DeleteBottomN(0)
	s.Equal("1, 2, 3, 4, 5", s.l.String())

	s.l.DeleteBottomN(1)
	s.Equal("1, 2, 3, 4", s.l.String())

	s.l.DeleteBottomN(2)
	s.Equal("1, 2, 4", s.l.String())

	s.l.DeleteBottomN(3)
	s.Equal("2, 4", s.l.String())

	s.l.DeleteBottomN(2)
	s.Equal("4", s.l.String())

	s.l.DeleteBottomN(3)
	s.Equal("4", s.l.String())
}

func (s *LinkedListAlgoTestSuite) TestHasCircle() {
	s.Equal(false, s.l.HasCircle())

	cur := s.l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = s.l.head.next.next

	s.Equal(true, s.l.HasCircle())
}

func (s *LinkedListAlgoTestSuite) TestFindMiddleNode() {
	s.Equal(3, s.l.FindMiddleNode().val)
}